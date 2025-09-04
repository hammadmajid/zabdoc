package utils

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"os"
)

type Template int

const (
	Index Template = iota
	Assignment
)

var TemplateFiles = map[Template]*template.Template{
	Index:      template.Must(template.ParseFiles("web/templates/index.html")),
	Assignment: template.Must(template.ParseFiles("web/templates/assignment.html")),
}

func GetLogoDataURI() (string, error) {
	logoBytes, err := os.ReadFile("web/static/szabist-logo.png")
	if err != nil {
		return "", fmt.Errorf("failed to read logo file: %w", err)
	}

	base64Logo := base64.StdEncoding.EncodeToString(logoBytes)
	return fmt.Sprintf("data:image/png;base64,%s", base64Logo), nil
}
