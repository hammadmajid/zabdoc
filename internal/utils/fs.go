package utils

import "html/template"

type Template int

const (
	Index Template = iota
	Assignment
)

var TemplateFiles = map[Template]*template.Template{
	Index:      template.Must(template.ParseFiles("web/templates/index.html")),
	Assignment: template.Must(template.ParseFiles("web/templates/assignment.html")),
}
