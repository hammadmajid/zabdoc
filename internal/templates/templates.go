package templates

import (
	"html/template"
)

type TemplateKey int

const (
	HomePage TemplateKey = iota
	AboutPage
	AssignmentPage
	AssignmentPdf
)

var templates = make(map[TemplateKey]*template.Template)

func GetTemplate(t TemplateKey) *template.Template {
	return templates[t]
}

func init() {
	templates[HomePage] = template.Must(template.New("index").Parse(renderWithShell(`web/templates/home.html`)))
	templates[AssignmentPage] = template.Must(template.New("assignment").Parse(renderWithShell(`web/templates/assignment.html`)))
	templates[AboutPage] = template.Must(template.New("about").Parse(renderWithShell(`web/templates/about.html`)))
	templates[AssignmentPdf] = template.Must(template.ParseFiles("web/templates/pdfs/assignment.html"))
}
