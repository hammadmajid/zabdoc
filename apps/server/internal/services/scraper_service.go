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
	"zabdoc/internal/types/models"
	"zabdoc/internal/types/requests"

	"golang.org/x/sync/errgroup"
)

// ScraperService handles attendance and marks data retrieval from ZabDesk
type ScraperService struct {
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

// NewScraperService creates a new Scraper instance
func NewScraperService() *ScraperService {
	jar, _ := cookiejar.New(nil)
	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 20,
		IdleConnTimeout:     90 * time.Second,
	}

	return &ScraperService{
		client: &http.Client{
			Jar:       jar,
			Transport: transport,
			Timeout:   20 * time.Second,
		},
	}
}

// ScrapeCourseData fetches and merges attendance and marks by course name.
func (s *ScraperService) ScrapeCourseData(req *requests.Scrape) (map[string]models.CourseScrapeData, error) {
	loginURL := "https://springzabdesk.szabist-isb.edu.pk/VerifyLogin.asp"
	resp, err := s.client.PostForm(loginURL, url.Values{
		"txtLoginName": {req.Username},
		"txtPassword":  {req.Password},
		"txtCampus_Id": {"1"},
	})
	if err != nil {
		return nil, err
	}
	if err := resp.Body.Close(); err != nil {
		return nil, err
	}

	var (
		attendanceByCourse map[string]models.CourseAttendance
		marksByCourse      map[string]models.CourseMarks
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

	results := make(map[string]models.CourseScrapeData)

	for courseName, attendance := range attendanceByCourse {
		entry := results[courseName]
		entry.Attendance = attendance
		entry.Marks = marksByCourse[courseName]
		results[courseName] = entry
	}

	for courseName, marks := range marksByCourse {
		if _, exists := results[courseName]; exists {
			continue
		}
		results[courseName] = models.CourseScrapeData{
			Attendance: models.CourseAttendance{CourseName: courseName},
			Marks:      marks,
		}
	}

	return results, nil
}

func (s *ScraperService) scrapeAttendanceByCourse() (map[string]models.CourseAttendance, error) {
	listURL := "https://springzabdesk.szabist-isb.edu.pk/Student/QryCourseAttendance.asp?OptionName=View%20Attendance"
	listPage, err := s.getPage(listURL)
	if err != nil {
		return nil, err
	}

	matches := reCourseLinks.FindAllStringSubmatch(listPage, -1)
	results := make(map[string]models.CourseAttendance, len(matches))

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
			course := models.CourseAttendance{
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

func (s *ScraperService) scrapeMarksByCourse() (map[string]models.CourseMarks, error) {
	listURL := "https://springzabdesk.szabist-isb.edu.pk/Student/QryCourseRecapSheet.asp?OptionName=Current%20Semester%20Results"
	listPage, err := s.getPage(listURL)
	if err != nil {
		return nil, err
	}

	matches := reCourseLinks.FindAllStringSubmatch(listPage, -1)
	results := make(map[string]models.CourseMarks, len(matches))

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
			courseMarks := models.CourseMarks{
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

func (s *ScraperService) getPage(pageURL string) (string, error) {
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

func (s *ScraperService) getCourseDetailPage(listURL string, match []string) (string, error) {
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

func (s *ScraperService) parseAttendanceRows(html string) []models.AttendanceRecord {
	rows := reAttendanceRows.FindAllStringSubmatch(html, -1)
	records := make([]models.AttendanceRecord, 0, len(rows))

	for _, r := range rows {
		records = append(records, models.AttendanceRecord{
			Lecture: r[1],
			Date:    r[2],
			Status:  strings.TrimSpace(r[3]),
		})
	}

	return records
}

func (s *ScraperService) extractMarksTable(pageHTML string) string {
	tableMatch := reMarksTable.FindString(pageHTML)
	if tableMatch != "" {
		return tableMatch
	}

	return pageHTML
}

func (s *ScraperService) parseMarksRows(tableHTML string) []models.MarkEntry {
	if tableHTML == "" {
		return nil
	}

	rawRows := reRows.FindAllString(tableHTML, -1)
	marks := make([]models.MarkEntry, 0, len(rawRows))

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

		marks = append(marks, models.MarkEntry{
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
func (s *ScraperService) parseTag(html, label string) string {
	re := regexp.MustCompile(fmt.Sprintf(`(?i)<th[^>]*>%s</th>\s*<td[^>]*>(.*?)</td>`, regexp.QuoteMeta(label)))
	m := re.FindStringSubmatch(html)
	if len(m) > 1 {
		return strings.TrimSpace(reTags.ReplaceAllString(m[1], ""))
	}
	return "N/A"
}
