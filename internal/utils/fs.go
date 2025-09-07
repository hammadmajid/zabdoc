package utils

import (
	"html/template"
)

type Template int

const (
	Index Template = iota
	About
	Assignment
)

var templates = map[Template]*template.Template{
	Index:      template.Must(template.ParseFiles("web/templates/index.html")),
	About:      template.Must(template.ParseFiles("web/templates/about.html")),
	Assignment: template.Must(template.ParseFiles("web/templates/assignment.html")),
}

func GetTemplate(t Template) *template.Template {
	return templates[t]
}
