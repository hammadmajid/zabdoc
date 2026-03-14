package services

import (
	"fmt"
	htmlstd "html"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"

	"zabdoc/internal/api/dto"

	"golang.org/x/sync/errgroup"
)

// Scraper handles attendance and marks data retrieval from ZabDesk
type Scraper struct {
	client *http.Client
}

const maxConcurrentRequests = 5

var (
	reCourseLinks    = regexp.MustCompile(`chkSubmit\('([^']+)','([^']+)','([^']+)','([^']+)'\)`)
	reAttendanceRows = regexp.MustCompile(`(?s)<tr>\s*<td[^>]*>(\d+)</td>\s*<td[^>]*>([\d/]+)</td>\s*<td[^>]*>\s*([a-zA-Z]+)\s*</td>\s*</tr>`)
	reMarksTable     = regexp.MustCompile(`(?is)<table[^>]*class="textColor"[^>]*>.*?<th[^>]*>\s*Marks\s*Head\s*</th>.*?<th[^>]*>\s*Marks\s*Obtained\s*</th>.*?</table>`)
	reRows           = regexp.MustCompile(`(?is)<tr[^>]*>.*?</tr>`)
	reTD             = regexp.MustCompile(`(?is)<td[^>]*>(.*?)</td>`)
	reTags           = regexp.MustCompile(`<[^>]*>`)
)

// NewScraper creates a new Scraper instance
func NewScraper() *Scraper {
	jar, _ := cookiejar.New(nil)
	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 20,
		IdleConnTimeout:     90 * time.Second,
	}

	return &Scraper{
		client: &http.Client{
			Jar:       jar,
			Transport: transport,
			Timeout:   20 * time.Second,
		},
	}
}

// ScrapeCourseData fetches and merges attendance and marks by course name.
func (s *Scraper) ScrapeCourseData(username, password string) (map[string]dto.CourseScrapeData, error) {
	loginURL := "https://springzabdesk.szabist-isb.edu.pk/VerifyLogin.asp"
	resp, err := s.client.PostForm(loginURL, url.Values{
		"txtLoginName": {username},
		"txtPassword":  {password},
		"txtCampus_Id": {"1"},
	})
	if err != nil {
		return nil, err
	}
	if err := resp.Body.Close(); err != nil {
		return nil, err
	}

	var (
		attendanceByCourse map[string]dto.CourseAttendance
		marksByCourse      map[string]dto.CourseMarks
	)

	g := new(errgroup.Group)
	g.Go(func() error {
		var scrapeErr error
		attendanceByCourse, scrapeErr = s.scrapeAttendanceByCourse()
		return scrapeErr
	})
	g.Go(func() error {
		var scrapeErr error
		marksByCourse, scrapeErr = s.scrapeMarksByCourse()
		return scrapeErr
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	results := make(map[string]dto.CourseScrapeData)

	for courseName, attendance := range attendanceByCourse {
		entry := results[courseName]
		entry.Attendence = attendance
		entry.Marks = marksByCourse[courseName]
		results[courseName] = entry
	}

	for courseName, marks := range marksByCourse {
		if _, exists := results[courseName]; exists {
			continue
		}
		results[courseName] = dto.CourseScrapeData{
			Attendence: dto.CourseAttendance{CourseName: courseName},
			Marks:      marks,
		}
	}

	return results, nil
}

func (s *Scraper) scrapeAttendanceByCourse() (map[string]dto.CourseAttendance, error) {
	listURL := "https://springzabdesk.szabist-isb.edu.pk/Student/QryCourseAttendance.asp?OptionName=View%20Attendance"
	listPage, err := s.getPage(listURL)
	if err != nil {
		return nil, err
	}

	matches := reCourseLinks.FindAllStringSubmatch(listPage, -1)
	results := make(map[string]dto.CourseAttendance, len(matches))

	var (
		mu sync.Mutex
		g  errgroup.Group
	)
	limit := make(chan struct{}, maxConcurrentRequests)

	for _, match := range matches {
		m := match
		g.Go(func() error {
			limit <- struct{}{}
			defer func() { <-limit }()

			dHTML, err := s.getCourseDetailPage(listURL, m)
			if err != nil {
				return err
			}

			courseName := s.parseTag(dHTML, "Course:")
			course := dto.CourseAttendance{
				CourseName: courseName,
				Instructor: s.parseTag(dHTML, "Instructor:"),
				Records:    s.parseAttendanceRows(dHTML),
			}

			mu.Lock()
			results[courseName] = course
			mu.Unlock()
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return results, nil
}

func (s *Scraper) scrapeMarksByCourse() (map[string]dto.CourseMarks, error) {
	listURL := "https://springzabdesk.szabist-isb.edu.pk/Student/QryCourseRecapSheet.asp?OptionName=Current%20Semester%20Results"
	listPage, err := s.getPage(listURL)
	if err != nil {
		return nil, err
	}

	matches := reCourseLinks.FindAllStringSubmatch(listPage, -1)
	results := make(map[string]dto.CourseMarks, len(matches))

	var (
		mu sync.Mutex
		g  errgroup.Group
	)
	limit := make(chan struct{}, maxConcurrentRequests)

	for _, match := range matches {
		m := match
		g.Go(func() error {
			limit <- struct{}{}
			defer func() { <-limit }()

			dHTML, err := s.getCourseDetailPage(listURL, m)
			if err != nil {
				return err
			}

			courseName := s.parseTag(dHTML, "Course:")
			marksTable := s.extractMarksTable(dHTML)
			courseMarks := dto.CourseMarks{
				Entries: s.parseMarksRows(marksTable),
			}

			mu.Lock()
			results[courseName] = courseMarks
			mu.Unlock()
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return results, nil
}

func (s *Scraper) getPage(pageURL string) (string, error) {
	resp, err := s.client.Get(pageURL)
	if err != nil {
		return "", err
	}
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(resp.Body)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (s *Scraper) getCourseDetailPage(listURL string, match []string) (string, error) {
	if len(match) < 5 {
		return "", fmt.Errorf("invalid course link data")
	}

	resp, err := s.client.PostForm(listURL, url.Values{
		"txtFac": {match[1]},
		"txtSem": {match[2]},
		"txtSec": {match[3]},
		"txtCou": {match[4]},
	})
	if err != nil {
		return "", err
	}
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(resp.Body)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (s *Scraper) parseAttendanceRows(html string) []dto.AttendanceRecord {
	rows := reAttendanceRows.FindAllStringSubmatch(html, -1)
	records := make([]dto.AttendanceRecord, 0, len(rows))

	for _, r := range rows {
		records = append(records, dto.AttendanceRecord{
			Lecture: r[1],
			Date:    r[2],
			Status:  strings.TrimSpace(r[3]),
		})
	}

	return records
}

func (s *Scraper) extractMarksTable(pageHTML string) string {
	tableMatch := reMarksTable.FindString(pageHTML)
	if tableMatch != "" {
		return tableMatch
	}

	return pageHTML
}

func (s *Scraper) parseMarksRows(tableHTML string) []dto.MarkEntry {
	if tableHTML == "" {
		return nil
	}

	rawRows := reRows.FindAllString(tableHTML, -1)
	marks := make([]dto.MarkEntry, 0, len(rawRows))

	for _, rowHTML := range rawRows {
		if strings.Contains(strings.ToLower(rowHTML), "colspan") {
			continue
		}

		tds := reTD.FindAllStringSubmatch(rowHTML, -1)
		if len(tds) < 3 {
			continue
		}

		head := cleanText(tds[0][1])
		if head == "" || strings.EqualFold(head, "Marks Head") {
			continue
		}
		if !isAllowedMarksHead(head) {
			continue
		}

		marks = append(marks, dto.MarkEntry{
			Head:     head,
			Max:      cleanText(tds[1][1]),
			Obtained: cleanText(tds[2][1]),
		})
	}

	return marks
}

func cleanText(value string) string {
	withoutTags := reTags.ReplaceAllString(value, " ")
	unescaped := htmlstd.UnescapeString(withoutTags)
	compact := strings.Join(strings.Fields(unescaped), " ")
	return strings.TrimSpace(compact)
}

func isAllowedMarksHead(head string) bool {
	value := strings.ToLower(strings.TrimSpace(head))
	allowed := []string{
		"assignment", "quiz", "mid term", "midterm", "final term", "final",
		"lab", "project", "viva", "report", "presentation", "practical", "cp", "objective",
	}

	for _, prefix := range allowed {
		if strings.HasPrefix(value, prefix) {
			return true
		}
	}

	return false
}

// parseTag extracts a value from an HTML table row by label.
func (s *Scraper) parseTag(html, label string) string {
	re := regexp.MustCompile(fmt.Sprintf(`(?i)<th[^>]*>%s</th>\s*<td[^>]*>(.*?)</td>`, regexp.QuoteMeta(label)))
	m := re.FindStringSubmatch(html)
	if len(m) > 1 {
		return strings.TrimSpace(reTags.ReplaceAllString(m[1], ""))
	}
	return "N/A"
}
