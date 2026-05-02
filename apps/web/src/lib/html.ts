// @apps/web/src/lib/html.ts
// Go template -> JS. No images. URLs for fonts/logo.

interface Student {
  Name: string;
  RegNo: string;
}

interface DocumentData {
  Students: Student[];
  Class: string;
  Course: string;
  CourseCode: string;
  Instructor: string;
  DocType: string;
  Number: string;
  Date: string;
  Marks: string;
}

// Helper fns matching Go template funcs
const add = (a: number, b: number) => a + b;
const div = (a: number, b: number) => (b === 0 ? 0 : Math.floor(a / b));
const lt = (a: number, b: number) => a < b;
const ge = (a: number, b: number) => a >= b;

const styles = `
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

    /* Force backgrounds to always appear in print */
    * {
        -webkit-print-color-adjust: exact;
        print-color-adjust: exact;
    }
</style>`;

function renderStudentTable(students: Student[]): string {
  if (ge(students.length, 4)) {
    // 4+ students: 4-column grid
    const mid = div(add(students.length, 1), 2);
    let grid = '';
    for (let i = 0; i < mid; i++) {
      const left = students[i];
      const right = students[i + mid];
      grid += `
        <div class="student-name">${left.Name}</div>
        <div class="student-regno">${left.RegNo}</div>
        ${right ? `
        <div class="student-name">${right.Name}</div>
        <div class="student-regno">${right.RegNo}</div>
        ` : `
        <div class="student-name"></div>
        <div class="student-regno"></div>
        `}
      `;
    }
    return `
    <div class="student-info-table compact">
      <div class="student-table-header">
        <div>Student Name</div>
        <div>Reg. Number</div>
        <div>Student Name</div>
        <div>Reg. Number</div>
      </div>
      <div class="student-table-body-grid">
        ${grid}
      </div>
    </div>`;
  } else {
    // 1-3 students: 2-column layout
    const rows = students.map(s => `
      <div class="student-entry-simple">
        <div>${s.Name}</div>
        <div>${s.RegNo}</div>
      </div>
    `).join('');
    return `
    <div class="student-info-table">
      <div class="student-table-header-simple">
        <div>Student Name</div>
        <div>Reg. Number</div>
      </div>
      <div class="student-table-body-simple">
        ${rows}
      </div>
    </div>`;
  }
}

export function buildHTML(data: DocumentData): string {
  const isMultiMode = ge(data.Students.length, 2);
  const firstStudent = data.Students[0];

  const infoTable = isMultiMode
    ? `
    <div class="info-table">
      <div class="info-row">
        <div class="info-label">Submitted to:</div>
        <div class="info-value">${data.Instructor}</div>
      </div>
      <div class="info-row">
        <div class="info-label">Class/Section:</div>
        <div class="info-value">${data.Class}</div>
      </div>
    </div>
    `
    : `
    <div class="info-table">
      <div class="info-row">
        <div class="info-label">Submitted to:</div>
        <div class="info-value">${data.Instructor}</div>
      </div>
      <div class="info-row">
        <div class="info-label">Student Name:</div>
        <div class="info-value">${firstStudent.Name}</div>
      </div>
      <div class="info-row">
        <div class="info-label">Reg. Number:</div>
        <div class="info-value">${firstStudent.RegNo}</div>
      </div>
      <div class="info-row">
        <div class="info-label">Class/Section:</div>
        <div class="info-value">${data.Class}</div>
      </div>
    </div>
    `;

  const spacerSmall = '<div class="spacer-small"></div>';
  const spacerLarge = '<div class="spacer-large"></div><div class="spacer-large"></div>';
  const spacerAdaptive = '<div class="spacer-adaptive"></div>';
  const spacerAdaptiveLarge = '<div class="spacer-adaptive-large"></div>';

  return `<!DOCTYPE html>
<html lang="en">

<head>
    <title>${firstStudent.RegNo}</title>
    ${styles}
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
        ${isMultiMode ? spacerAdaptive : spacerSmall}
        <div class="marks">
            <div class="marks-line">Total Marks: <span class="underline">${data.Marks}</span></div>
            <div class="marks-line">Obtained Marks: <span class="underline"></span></div>
        </div>
        ${isMultiMode ? spacerAdaptiveLarge : spacerLarge}
        <div class="course-title-section">
            <div class="course-title">${data.Course}</div>
            <div class="assignment-number">${data.DocType} #${data.Number}</div>
            <div class="submission-date">Submission date: ${data.Date}</div>
        </div>
        ${isMultiMode ? spacerAdaptiveLarge : spacerLarge}
        ${infoTable}
        ${isMultiMode ? renderStudentTable(data.Students) : ''}
    </div>

    <div class="footer">
        <div>${data.CourseCode}</div>
        <div>${data.Class}</div>
        <div>SZABIST-ISB</div>
    </div>
</div>

</body>
</html>`;
}
