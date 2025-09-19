package templates

import (
	"html/template"
)

var Tpl = template.Must(template.New("document").Funcs(template.FuncMap{
	"add": func(a, b int) int {
		return a + b
	},
	"html": func(s string) template.HTML {
		return template.HTML(s)
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
        <div class="spacer-small"></div>
        ` + MarksTemplate + `
        <div class="spacer-large"></div>
        <div class="course-title-section">
            <div class="course-title">{{.Course}}</div>
            {{if .ProjectTitle}}
                <div class="project-title">{{.ProjectTitle}}</div>
                <div class="assignment-number">{{.DocType}}</div>
            {{else}}
                <div class="assignment-number">{{.DocType}} #{{.Number}}</div>
            {{end}}
            <div class="submission-date">Submission date: {{.Date}}</div>
        </div>
        <div class="spacer-large"></div>
        ` + InfoTableTemplate + `
    </div>
    ` + FooterTemplate + `
</div>

<!-- Content Pages -->
` + ContentPagesTemplate + `
</body>
</html>`))
