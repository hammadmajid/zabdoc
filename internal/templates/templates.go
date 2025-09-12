package templates

import (
	"html/template"
)

type TemplateKey int

const (
	Assignment = iota
	LabTask
)

var templates = make(map[TemplateKey]*template.Template)

func GetTemplate(t TemplateKey) *template.Template {
	return templates[t]
}

func init() {
	templates[Assignment] = template.Must(template.New("assignment_pdf").Parse(`
<!DOCTYPE html>
<html lang="en">

<head>
    <title>Assignment Cover</title>
    <style>
        @font-face {
            font-family: "Times New Roman";
            src: url("https://cdn.zabdoc.xyz/times.ttf") format("truetype");
            font-weight: normal;
            font-style: normal;
            font-display: swap;
        }

        /* Lock page size */
        @page {
            size: A4;
            margin: 7mm;
        }

        html,
        body {
            margin: 0;
            padding: 0;
            width: 210mm;
            height: 297mm;
            font-family: "Times New Roman", serif;
            background: #fff;
            color: #222;
        }

        body {
            display: flex;
            justify-content: center;
            align-items: center;
        }

        .a4-page {
            width: 210mm;
            height: 297mm;
            box-sizing: border-box;
            padding: 5mm;
            /* use mm to match printer/PDF units */
            display: flex;
            flex-direction: column;
            justify-content: space-between;
            border: 1px solid #000;
            page-break-after: always;
        }

        .header {
            display: grid;
            grid-template-columns: 1fr 3fr;
            align-items: center;
            margin-bottom: 18pt;
        }

        .header-logo {
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100%;
        }

        .header-logo img {
            max-height: 60pt;
            max-width: 100%;
        }

        .header-content {
            text-align: center;
        }

        .header-line1 {
            font-size: 14pt;
            /* use pt for stable scaling */
            font-weight: bold;
        }

        .header-line2-box {
            border: 2px solid #222;
            padding: 4pt 10pt;
            margin-top: 6pt;
            font-size: 12pt;
            font-weight: bold;
            letter-spacing: 1pt;
            background: #f7f7f7;
        }

        .header-line2 {
            text-transform: uppercase;
        }

        .spacer-small {
            height: 16pt;
        }

        .spacer-large {
            height: 44pt;
        }

        .content {
            flex: 1;
            display: flex;
            flex-direction: column;
            justify-content: space-around;
            gap: 24pt;
        }

        .marks {
            text-align: right;
        }

        .marks-line {
            margin-bottom: 6pt;
            font-size: 15pt;
        }

        .underline {
            display: inline-block;
            width: 60pt;
            text-align: center;
            border-bottom: 2px solid #444;
            padding-left: 0.5em;
            padding-right: 0.5em;
            padding-bottom: 0.2em;
        }

        .course-title-section {
            text-align: center;
        }

        .course-title {
            font-size: 30pt;
            font-weight: bold;
            text-transform: uppercase;
            margin-bottom: 6pt;
        }

        .assignment-number {
            font-size: 25pt;
            font-weight: 600;
            margin-bottom: 4pt;
        }

        .submission-date {
            font-size: 22pt;
            margin-bottom: 0;
        }

        .info-table {
            width: 70%;
            margin: 0 auto 60pt;
            font-size: 20pt;
        }

        .info-row {
            display: flex;
            margin-bottom: 10pt;
        }

        .info-label {
            width: 60mm;
            font-weight: bold;
            text-align: right;
            padding-right: 10pt;
        }

        .info-value {
            flex: 1;
            border-bottom: 1px solid #ccc;
            padding-bottom: 2pt;
            text-align: center;
        }

        .footer {
            display: flex;
            justify-content: space-between;
            align-items: center;
            font-size: 14pt;
        }

        /* Force backgrounds to always appear in print */
        * {
            -webkit-print-color-adjust: exact;
            print-color-adjust: exact;
        }
    </style>
</head>

<body>
<div class="a4-page">
    <div class="header">
        <div class="header-logo">
            <img src="https://cdn.zabdoc.xyz/szabist-logo.png" alt="SZABIST Logo">
        </div>
        <div class="header-content">
            <div class="header-line1">Shaheed Zulfiqar Ali Bhutto Institute of Science and Technology</div>
            <div class="header-line2-box"><span class="header-line2">Computer Science Department</span></div>
        </div>
    </div>
    <div class="content">
        <div class="spacer-small"></div>
        <div class="marks">
            <div class="marks-line">Total Marks: <span class="underline">4</span></div>
            <div class="marks-line">Obtained Marks: <span class="underline"></span></div>
        </div>
        <div class="spacer-large"></div>
        <div class="course-title-section">
            <div class="course-title">{{.Course}}</div>
            <div class="assignment-number">Assignment #{{.Number}}</div>
            <div class="submission-date">Submission date: {{.Date}}</div>
        </div>
        <div class="spacer-large"></div>
        <div class="info-table">
            <div class="info-row">
                <div class="info-label">Student Name:</div>
                <div class="info-value">{{.StudentName}}</div>
            </div>
            <div class="info-row">
                <div class="info-label">Reg. Number:</div>
                <div class="info-value">{{.RegNo}}</div>
            </div>
            <div class="info-row">
                <div class="info-label">Submitted to:</div>
                <div class="info-value">{{.Instructor}}</div>
            </div>
            <div class="info-row">
                <div class="info-label">Class/Section:</div>
                <div class="info-value">{{.Class}}</div>
            </div>
        </div>
    </div>
    <div class="footer">
        <div>{{.CourseCode}}</div>
        <div>{{.Class}}</div>
        <div>SZABIST-ISB</div>
    </div>
</div>
</body>
</html>
`))

	templates[LabTask] = template.Must(template.New("lab_task").Funcs(template.FuncMap{
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
    <title>Lab Task</title>
    <style>
        @font-face {
            font-family: "Times New Roman";
            src: url("https://cdn.zabdoc.xyz/times.ttf") format("truetype");
            font-weight: normal;
            font-style: normal;
            font-display: swap;
        }

        /* Lock page size */
        @page {
            size: A4;
            margin: 7mm;
        }

        html,
        body {
            margin: 0;
            padding: 0;
            font-family: "Times New Roman", serif;
            background: #fff;
            color: #000;
            line-height: 1.6;
        }

        /* Cover Page Styles */
        .cover-page {
            width: 210mm;
            height: 297mm;
            box-sizing: border-box;
            padding: 5mm;
            display: flex;
            flex-direction: column;
            justify-content: space-between;
            border: 1px solid #000;
            page-break-after: always;
        }

        .header {
            display: grid;
            grid-template-columns: 1fr 3fr;
            align-items: center;
            margin-bottom: 18pt;
        }

        .header-logo {
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100%;
        }

        .header-logo img {
            max-height: 60pt;
            max-width: 100%;
        }

        .header-content {
            text-align: center;
        }

        .header-line1 {
            font-size: 14pt;
            font-weight: bold;
        }

        .header-line2-box {
            border: 2px solid #222;
            padding: 4pt 10pt;
            margin-top: 6pt;
            font-size: 12pt;
            font-weight: bold;
            letter-spacing: 1pt;
            background: #f7f7f7;
        }

        .header-line2 {
            text-transform: uppercase;
        }

        .spacer-small {
            height: 16pt;
        }

        .spacer-large {
            height: 44pt;
        }

        .content {
            flex: 1;
            display: flex;
            flex-direction: column;
            justify-content: space-around;
            gap: 24pt;
        }

        .marks {
            text-align: right;
        }

        .marks-line {
            margin-bottom: 6pt;
            font-size: 15pt;
        }

        .underline {
            display: inline-block;
            width: 60pt;
            text-align: center;
            border-bottom: 2px solid #444;
            padding-left: 0.5em;
            padding-right: 0.5em;
            padding-bottom: 0.2em;
        }

        .course-title-section {
            text-align: center;
        }

        .course-title {
            font-size: 30pt;
            font-weight: bold;
            text-transform: uppercase;
            margin-bottom: 6pt;
        }

        .assignment-number {
            font-size: 25pt;
            font-weight: 600;
            margin-bottom: 4pt;
        }

        .submission-date {
            font-size: 22pt;
            margin-bottom: 0;
        }

        .info-table {
            width: 70%;
            margin: 0 auto 60pt;
            font-size: 20pt;
        }

        .info-row {
            display: flex;
            margin-bottom: 10pt;
        }

        .info-label {
            width: 60mm;
            font-weight: bold;
            text-align: right;
            padding-right: 10pt;
        }

        .info-value {
            flex: 1;
            border-bottom: 1px solid #ccc;
            padding-bottom: 2pt;
            text-align: center;
        }

        .footer {
            display: flex;
            justify-content: space-between;
            align-items: center;
            font-size: 14pt;
        }

        /* Content Pages Styles */
        .content-pages {
            font-size: 12pt;
            line-height: 1.6;
        }

        .task-section {
            margin-bottom: 2em;
            page-break-inside: avoid;
        }

        .task-header {
            font-size: 16pt;
            font-weight: bold;
            margin-bottom: 1em;
            border-bottom: 2px solid #222;
            padding-bottom: 0.5em;
        }

        .task-content {
            margin-bottom: 1.5em;
        }

        .task-content h1, .task-content h2, .task-content h3 {
            color: #333;
            margin-top: 1.5em;
            margin-bottom: 0.5em;
        }

        .task-content h1 {
            font-size: 18pt;
            border-bottom: 1px solid #ccc;
            padding-bottom: 0.3em;
        }

        .task-content h2 {
            font-size: 16pt;
        }

        .task-content h3 {
            font-size: 14pt;
        }

        .task-content p {
            margin-bottom: 1em;
            text-align: justify;
        }

        .task-content ul, .task-content ol {
            margin-bottom: 1em;
            padding-left: 2em;
        }

        .task-content li {
            margin-bottom: 0.5em;
        }

        .task-content pre {
            background: #f5f5f5;
            padding: 1em;
            border-radius: 4px;
            border: 1px solid #ddd;
            overflow: auto;
            font-family: "Courier New", monospace;
            font-size: 10pt;
            page-break-inside: avoid;
        }

        .task-content code {
            background: #f5f5f5;
            padding: 2px 4px;
            border-radius: 3px;
            font-family: "Courier New", monospace;
            font-size: 10pt;
        }

        .task-content blockquote {
            border-left: 4px solid #ccc;
            margin-left: 0;
            padding-left: 1em;
            color: #666;
            font-style: italic;
        }

        .task-images {
            margin-top: 1.5em;
        }

        .task-image {
            max-width: 100%;
            height: auto;
            margin-bottom: 1em;
            border: 1px solid #ddd;
            border-radius: 4px;
            page-break-inside: avoid;
            display: block;
        }

        /* Force backgrounds to always appear in print */
        * {
            -webkit-print-color-adjust: exact;
            print-color-adjust: exact;
        }
    </style>
</head>

<body>
<!-- Cover Page -->
<div class="cover-page">
    <div class="header">
        <div class="header-logo">
            <img src="https://cdn.zabdoc.xyz/szabist-logo.png" alt="SZABIST Logo">
        </div>
        <div class="header-content">
            <div class="header-line1">Shaheed Zulfiqar Ali Bhutto Institute of Science and Technology</div>
            <div class="header-line2-box"><span class="header-line2">Computer Science Department</span></div>
        </div>
    </div>
    <div class="content">
        <div class="spacer-small"></div>
        <div class="marks">
            <div class="marks-line">Total Marks: <span class="underline">7.5</span></div>
            <div class="marks-line">Obtained Marks: <span class="underline"></span></div>
        </div>
        <div class="spacer-large"></div>
        <div class="course-title-section">
            <div class="course-title">{{.Course}}</div>
            <div class="assignment-number">Lab Task #{{.Number}}</div>
            <div class="submission-date">Submission date: {{.Date}}</div>
        </div>
        <div class="spacer-large"></div>
        <div class="info-table">
            <div class="info-row">
                <div class="info-label">Student Name:</div>
                <div class="info-value">{{.StudentName}}</div>
            </div>
            <div class="info-row">
                <div class="info-label">Reg. Number:</div>
                <div class="info-value">{{.RegNo}}</div>
            </div>
            <div class="info-row">
                <div class="info-label">Submitted to:</div>
                <div class="info-value">{{.Instructor}}</div>
            </div>
            <div class="info-row">
                <div class="info-label">Class/Section:</div>
                <div class="info-value">{{.Class}}</div>
            </div>
        </div>
    </div>
    <div class="footer">
        <div>{{.CourseCode}}</div>
        <div>{{.Class}}</div>
        <div>SZABIST-ISB</div>
    </div>
</div>

<!-- Content Pages -->
{{if .Tasks}}
<div class="content-pages">
    {{range $index, $task := .Tasks}}
    <div class="task-section">
        <div class="task-header">Task {{add $task.Index 1}}</div>
        <div class="task-content">
            {{$task.Content}}
        </div>
        {{if $task.Images}}
        <div class="task-images">
            {{range $imgIndex, $image := $task.Images}}
            <img src="data:{{$image.MimeType}};base64,{{$image.Data}}" alt="Task {{add $task.Index 1}} Image {{add $imgIndex 1}}" class="task-image">
            {{end}}
        </div>
        {{end}}
    </div>
    {{end}}
</div>
{{end}}
</body>
</html>`))
}
