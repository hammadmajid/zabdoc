package templates

import (
	"html/template"
)

var Tpl = template.Must(template.New("document").Funcs(template.FuncMap{
	"add": func(a, b int) int {
		return a + b
	},
	"div": func(a, b int) int {
		if b == 0 {
			return 0
		}
		return a / b
	},
	"lt": func(a, b int) bool {
		return a < b
	},
	"ge": func(a, b int) bool {
		return a >= b
	},
	"gt": func(a, b int) bool {
		return a > b
	},
}).Parse(`
<!DOCTYPE html>
<html lang="en">

<head>
    <title>{{ .RegNo }}</title>
    <style>` + StylesTemplate + `</style>
</head>

<body>
<!-- Cover Page -->
<div class="cover-page">
    ` + HeaderTemplate + `
    <div class="content">
        {{if .IsMultiMode}}
        <div class="spacer-adaptive"></div>
        {{else}}
        <div class="spacer-small"></div>
        {{end}}
        ` + MarksTemplate + `
        {{if .IsMultiMode}}
        <div class="spacer-adaptive-large"></div>
        {{else}}
        <div class="spacer-large"></div>
        <div class="spacer-large"></div>
        {{end}}
        <div class="course-title-section">
            <div class="course-title">{{.Course}}</div>
            {{if .ProjectTitle}}
                <div class="assignment-number">{{.DocType}}</div>
            {{else}}
                <div class="assignment-number">{{.DocType}} #{{.Number}}</div>
            {{end}}
            <div class="submission-date">Submission date: {{.Date}}</div>
        </div>
        {{if .ProjectTitle}}
        {{if .IsMultiMode}}
        <div class="spacer-adaptive"></div>
        {{else}}
        <div class="spacer-small"></div>
        {{end}}
        <div class="project-title-section">
            <div class="project-title">{{.ProjectTitle}}</div>
        </div>
        {{end}}
        {{if .IsMultiMode}}
        <div class="spacer-adaptive-large"></div>
        {{else}}
        <div class="spacer-large"></div>
        <div class="spacer-large"></div>
        {{end}}
        {{if .IsMultiMode}}
        ` + SharedInfoTableTemplate + `
        ` + StudentInfoTableTemplate + `
        {{else}}
        ` + InfoTableTemplate + `
        {{end}}
    </div>
    ` + FooterTemplate + `
</div>

<!-- Content Pages -->
` + ContentPagesTemplate + `
</body>
</html>`))
