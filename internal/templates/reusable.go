package templates

const StylesTemplate = `
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

    /* Adaptive spacers for multi-student mode */
    .spacer-adaptive {
        height: 3vh;
    }

    .spacer-adaptive-large {
        height: 6vh;
    }

    .content {
        flex: 1;
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        gap: 0;
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

    .project-title-section {
        text-align: center;
    }

    .project-title {
        font-size: 22pt;
        font-weight: 500;
        margin-bottom: 8pt;
        color: #333;
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

    .info-value2 {
        flex: 1;
        padding-bottom: 2pt;
        text-align: center;
    }

    /* Multi-student table styles with adaptive sizing */
    .student-info-table {
        width: 85%;
        margin: 1vh auto 2vh;
        font-size: 16pt;
    }

    .student-info-table.compact {
        font-size: 14pt;
        margin: 0.5vh auto 1vh;
    }

    /* Simple two-column layout for 1-3 students */
    .student-table-header-simple {
        display: grid;
        grid-template-columns: 1fr 1fr;
        font-weight: bold;
        border-bottom: 2px solid #333;
        margin-bottom: 6pt;
        padding-bottom: 4pt;
        text-align: center;
    }

    .student-table-body-simple {
        display: flex;
        flex-direction: column;
    }

    .student-entry-simple {
        display: grid;
        grid-template-columns: 1fr 1fr;
        padding: 3pt 0;
        border-bottom: 1px solid #ddd;
        text-align: center;
    }

    /* Four-column grid layout for 4+ students */
    .student-table-header {
        display: grid;
        grid-template-columns: 1fr 1fr 1fr 1fr;
        font-weight: bold;
        border-bottom: 2px solid #333;
        margin-bottom: 6pt;
        padding-bottom: 4pt;
        gap: 10pt;
        text-align: center;
    }

    .student-table-body-grid {
        display: grid;
        grid-template-columns: 1fr 1fr 1fr 1fr;
        gap: 10pt;
        row-gap: 6pt;
    }

    .student-name {
        text-align: center;
        padding: 2pt 0;
        border-bottom: 1px solid #ddd;
    }

    .student-regno {
        text-align: center;
        padding: 2pt 0;
        border-bottom: 1px solid #ddd;
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

    .content-page {
        width: 210mm;
        height: 297mm;
        box-sizing: border-box;
        padding: 5mm;
        border: 1px solid #000;
        page-break-after: always;
        display: flex;
        flex-direction: column;
    }

    .content-page:last-child {
        page-break-after: auto;
    }

    .content-page-header {
        flex-shrink: 0;
        margin-bottom: 8pt;
    }

    .content-page-body {
        flex: 1;
        overflow: hidden;
        padding: 4pt 0;
    }

    .content-page-footer {
        flex-shrink: 0;
        margin-top: 8pt;
    }

    .section-images {
        margin-top: 1.5em;
    }

    .section-image {
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
</style>`

const HeaderTemplate = `
<div class="header">
    <div class="header-logo">
        <img src="https://cdn.zabdoc.xyz/szabist-logo.png" alt="SZABIST Logo">
    </div>
    <div class="header-content">
        <div class="header-line1">Shaheed Zulfiqar Ali Bhutto Institute of Science and Technology</div>
        <div class="header-line2-box"><span class="header-line2">Computer Science Department</span></div>
    </div>
</div>`

const MarksTemplate = `
<div class="marks">
    <div class="marks-line">Total Marks: <span class="underline">{{.Marks}}</span></div>
    <div class="marks-line">Obtained Marks: <span class="underline"></span></div>
</div>`

const FooterTemplate = `
<div class="footer">
    <div>{{.CourseCode}}</div>
    <div>{{.Class}}</div>
    <div>SZABIST-ISB</div>
</div>`

const InfoTableTemplate = `
<div class="info-table">
    <div class="info-row">
        <div class="info-label">Submitted to:</div>
        <div class="info-value">{{.Instructor}}</div>
    </div>
    <div class="info-row">
        <div class="info-label">Student Name:</div>
        <div class="info-value">{{.FirstStudent.Name}}</div>
    </div>
    <div class="info-row">
        <div class="info-label">Reg. Number:</div>
        <div class="info-value">{{.FirstStudent.RegNo}}</div>
    </div>
    <div class="info-row">
        <div class="info-label">Class/Section:</div>
        <div class="info-value">{{.Class}}</div>
    </div>
</div>`

const SharedInfoTableTemplate = `
<div class="info-table">
    <div class="info-row">
        <div class="info-label">Submitted to:</div>
        <div class="info-value">{{.Instructor}}</div>
    </div>
    <div class="info-row">
        <div class="info-label">Class/Section:</div>
        <div class="info-value">{{.Class}}</div>
    </div>
</div>`

const StudentInfoTableTemplate = `
<div class="student-info-table{{if ge (len .Students) 4}} compact{{end}}">
    {{if ge (len .Students) 4}}
    {{/* 4+ students: split into two side-by-side tables */}}
    <div class="student-table-header">
        <div>Student Name</div>
        <div>Reg. Number</div>
        <div>Student Name</div>
        <div>Reg. Number</div>
    </div>
    <div class="student-table-body-grid">
        {{$mid := div (add (len .Students) 1) 2}}
        {{range $index, $student := .Students}}
            {{if lt $index $mid}}
            <div class="student-name">{{$student.Name}}</div>
            <div class="student-regno">{{$student.RegNo}}</div>
            {{with index $.Students (add $index $mid)}}
            <div class="student-name">{{.Name}}</div>
            <div class="student-regno">{{.RegNo}}</div>
            {{else}}
            <div class="student-name"></div>
            <div class="student-regno"></div>
            {{end}}
            {{end}}
        {{end}}
    </div>
    {{else}}
    {{/* 1-3 students: simple two-column layout */}}
    <div class="student-table-header-simple">
        <div>Student Name</div>
        <div>Reg. Number</div>
    </div>
    <div class="student-table-body-simple">
        {{range $student := .Students}}
        <div class="student-entry-simple">
            <div>{{$student.Name}}</div>
            <div>{{$student.RegNo}}</div>
        </div>
        {{end}}
    </div>
    {{end}}
</div>`

const ContentPagesTemplate = `
{{if .Images}}
<div class="content-pages">
    <div class="content-page">
        <div class="content-page-header">
            <div class="header">
                <div class="header-logo">
                    <img src="https://cdn.zabdoc.xyz/szabist-logo.png" alt="SZABIST Logo">
                </div>
                <div class="header-content">
                    <div class="header-line1">Shaheed Zulfiqar Ali Bhutto Institute of Science and Technology</div>
                    <div class="header-line2-box"><span class="header-line2">Computer Science Department</span></div>
                </div>
            </div>
        </div>
        <div class="content-page-body">
            <div class="section-images">
                {{range $imgIndex, $image := .Images}}
                <img src="data:{{$image.MimeType}};base64,{{$image.Data}}" alt="Image {{add $imgIndex 1}}" class="section-image">
                {{end}}
            </div>
        </div>
        <div class="content-page-footer">
            <div class="footer">
                <div>{{$.CourseCode}}</div>
                <div>{{$.Class}}</div>
                <div>SZABIST-ISB</div>
            </div>
        </div>
    </div>
</div>
{{end}}`
