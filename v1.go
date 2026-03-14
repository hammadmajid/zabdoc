package main

import (
	"fmt"
	htmlstd "html"
	"html/template"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

type MarkEntry struct {
	Head     string
	Max      string
	Obtained string
}

type CourseResult struct {
	CourseName string
	Marks      []MarkEntry
}

const layoutHTML = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="https://unpkg.com/terminal.css@0.7.4/dist/terminal.min.css" />
    <title>Semester Results Extraction</title>
    <style>
        :root { --page-width: 900px; }
        .course-card { margin-bottom: 40px; border: 1px solid #444; padding: 20px; border-radius: 4px; }
        .total-row { background-color: #333; font-weight: bold; }
        .not-entered { color: #888; font-style: italic; }
    </style>
</head>
<body class="terminal">
    <div class="container">
        <header class="terminal-nav">
            <div class="terminal-logo">
                <div class="logo terminal-prompt">DATA_EXTRACT_RECAP_SHEET</div>
            </div>
        </header>
        {{content}}
    </div>
</body>
</html>`

const resultHTML = `
<section>
    {{range .}}
    <div class="course-card">
        <header>
            <h2 style="margin-bottom: 5px;">{{.CourseName}}</h2>
        </header>
        <table class="terminal-table">
            <thead>
                <tr>
                    <th>Marks Head</th>
                    <th>Max Marks</th>
                    <th>Obtained</th>
                </tr>
            </thead>
            <tbody>
                {{range .Marks}}
                <tr>
                    <td>{{.Head}}</td>
                    <td>{{.Max}}</td>
                    <td class="{{if eq .Obtained "Not Entered"}}not-entered{{end}}">{{.Obtained}}</td>
                </tr>
                {{end}}
            </tbody>
        </table>
    </div>
    {{else}}
    <div class="terminal-alert terminal-alert-error">No results found or extraction failed.</div>
    {{end}}
    <a href="/" class="btn btn-primary">New Extraction</a>
</section>`

const indexHTML = `
<section>
	<h1>Semester Results Extraction</h1>
	<p>Enter your ZabDesk credentials to fetch recap sheets.</p>
	<form action="/fetch" method="post">
		<div class="form-group">
			<label for="username">Username</label>
			<input id="username" name="username" type="text" required />
		</div>
		<div class="form-group">
			<label for="password">Password</label>
			<input id="password" name="password" type="password" required />
		</div>
		<button class="btn btn-primary" type="submit">Fetch Results</button>
	</form>
</section>`

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/fetch", handleFetchResults)
	fmt.Println("Result Scraper running on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server error:", err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	render(w, indexHTML, nil)
}

func handleFetchResults(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	user, pass := r.FormValue("username"), r.FormValue("password")
	data, err := scrapeResults(user, pass)
	if err != nil {
		http.Error(w, "failed to fetch results: "+err.Error(), http.StatusBadGateway)
		return
	}
	render(w, resultHTML, data)
}

func scrapeResults(username, password string) ([]CourseResult, error) {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	loginURL := "https://springzabdesk.szabist-isb.edu.pk/VerifyLogin.asp?sid=519991263"
	if _, err := client.PostForm(loginURL, url.Values{
		"txtLoginName": {username},
		"txtPassword":  {password},
		"txtCampus_Id": {"1"},
	}); err != nil {
		return nil, err
	}

	listURL := "https://springzabdesk.szabist-isb.edu.pk/Student/QryCourseRecapSheet.asp?OptionName=Current%20Semester%20Results&sid=519991263"
	resp, err := client.Get(listURL)
	if err != nil {
		return nil, err
	}
	b, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	reLinks := regexp.MustCompile(`chkSubmit\('([^']+)','([^']+)','([^']+)','([^']+)'\)`)
	matches := reLinks.FindAllStringSubmatch(string(b), -1)

	var results []CourseResult
	for _, m := range matches {
		dResp, err := client.PostForm(listURL, url.Values{
			"txtFac": {m[1]}, "txtSem": {m[2]}, "txtSec": {m[3]}, "txtCou": {m[4]},
		})
		if err != nil {
			return nil, err
		}
		db, _ := io.ReadAll(dResp.Body)
		dResp.Body.Close()
		dHTML := string(db)

		res := CourseResult{
			CourseName: parseTag(dHTML, "Course:"),
		}

		marksTable := extractMarksTable(dHTML)
		rows := parseMarksRows(marksTable)
		res.Marks = append(res.Marks, rows...)

		results = append(results, res)
	}
	return results, nil
}

func parseTag(html, label string) string {
	re := regexp.MustCompile(fmt.Sprintf(`(?i)<th[^>]*>%s</th>\s*<td[^>]*>(.*?)</td>`, regexp.QuoteMeta(label)))
	m := re.FindStringSubmatch(html)
	if len(m) > 1 {
		return strings.TrimSpace(regexp.MustCompile("<[^>]*>").ReplaceAllString(m[1], ""))
	}
	return "N/A"
}

func extractMarksTable(pageHTML string) string {
	reTable := regexp.MustCompile(`(?is)<table[^>]*class="textColor"[^>]*>.*?<th[^>]*>\s*Marks\s*Head\s*</th>.*?<th[^>]*>\s*Marks\s*Obtained\s*</th>.*?</table>`)
	tableMatch := reTable.FindString(pageHTML)
	if tableMatch != "" {
		return tableMatch
	}
	return pageHTML
}

func parseMarksRows(tableHTML string) []MarkEntry {
	if tableHTML == "" {
		return nil
	}

	reRow := regexp.MustCompile(`(?is)<tr[^>]*>.*?</tr>`)
	reTD := regexp.MustCompile(`(?is)<td[^>]*>(.*?)</td>`)

	rawRows := reRow.FindAllString(tableHTML, -1)
	marks := make([]MarkEntry, 0, len(rawRows))

	for _, rStr := range rawRows {
		if strings.Contains(strings.ToLower(rStr), "colspan") {
			continue
		}

		tds := reTD.FindAllStringSubmatch(rStr, -1)
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

		marks = append(marks, MarkEntry{
			Head:     head,
			Max:      cleanText(tds[1][1]),
			Obtained: cleanText(tds[2][1]),
		})
	}

	return marks
}

func cleanText(s string) string {
	withoutTags := regexp.MustCompile(`<[^>]*>`).ReplaceAllString(s, " ")
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

func render(w http.ResponseWriter, contentTmpl string, data interface{}) {
	tString := strings.Replace(layoutHTML, "{{content}}", contentTmpl, 1)
	tmpl, err := template.New("v").Parse(tString)
	if err != nil {
		http.Error(w, "template parse error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "template execute error: "+err.Error(), http.StatusInternalServerError)
	}
}
