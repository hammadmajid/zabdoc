package templates

import (
	"bytes"
	"html/template"
	"path/filepath"
	"zabdoc/internal/api/dto"
)

type TemplateType int

const (
	Assignment TemplateType = iota
	LabTask
	LabProject
)

type TemplateConfig struct {
	TemplateType       TemplateType
	IncludeContentPage bool
}

type TemplateBuilder struct {
	// Pre-built templates for each configuration - NO runtime building
	assignmentTemplate    *template.Template
	assignmentWithContent *template.Template
	labTaskTemplate       *template.Template
	labTaskWithContent    *template.Template
	labProjectTemplate    *template.Template
	labProjectWithContent *template.Template

	// Separate header/footer templates for PDF service
	headerTemplate *template.Template
	footerTemplate *template.Template
}

func NewTemplateBuilder() (*TemplateBuilder, error) {
	// helper functions for templates
	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"html": func(s string) template.HTML {
			return template.HTML(s)
		},
	}

	// Parse all component templates once
	components := template.New("").Funcs(funcMap)
	pattern := filepath.Join("internal/templates/components", "*.tmpl")
	components, err := components.ParseGlob(pattern)
	if err != nil {
		return nil, err
	}

	// Pre-build ALL template variations at startup - ZERO runtime compilation
	builder := &TemplateBuilder{}

	// Assignment templates
	builder.assignmentTemplate, err = buildTemplate(components, `{{template "cover-page" .}}`)
	if err != nil {
		return nil, err
	}

	builder.assignmentWithContent, err = buildTemplate(components, `{{template "cover-page" .}}{{template "content-page" .}}`)
	if err != nil {
		return nil, err
	}

	// Lab task templates (same structure for now)
	builder.labTaskTemplate = builder.assignmentTemplate
	builder.labTaskWithContent = builder.assignmentWithContent

	// Lab project templates (same structure for now)
	builder.labProjectTemplate = builder.assignmentTemplate
	builder.labProjectWithContent = builder.assignmentWithContent

	// Header/footer templates for PDF service
	builder.headerTemplate, err = buildTemplate(components, `{{template "pdf-header" .}}`)
	if err != nil {
		return nil, err
	}

	builder.footerTemplate, err = buildTemplate(components, `{{template "pdf-footer" .}}`)
	if err != nil {
		return nil, err
	}

	return builder, nil
}

// Helper function to build a template with specific content
func buildTemplate(base *template.Template, contentTemplate string) (*template.Template, error) {
	tmpl, err := base.Clone()
	if err != nil {
		return nil, err
	}

	_, err = tmpl.New("content").Parse(contentTemplate)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func (tb *TemplateBuilder) BuildTemplate(config TemplateConfig) (*template.Template, error) {
	// NO RUNTIME BUILDING - just return pre-built template
	switch config.TemplateType {
	case Assignment:
		if config.IncludeContentPage {
			return tb.assignmentWithContent, nil
		}
		return tb.assignmentTemplate, nil
	case LabTask:
		if config.IncludeContentPage {
			return tb.labTaskWithContent, nil
		}
		return tb.labTaskTemplate, nil
	case LabProject:
		if config.IncludeContentPage {
			return tb.labProjectWithContent, nil
		}
		return tb.labProjectTemplate, nil
	default:
		if config.IncludeContentPage {
			return tb.assignmentWithContent, nil
		}
		return tb.assignmentTemplate, nil
	}
}

func (tb *TemplateBuilder) Execute(config TemplateConfig, data dto.GenerateRequest) ([]byte, error) {
	tmpl, err := tb.BuildTemplate(config)
	if err != nil {
		return nil, err
	}

	// Add config to data for template access
	templateData := struct {
		dto.GenerateRequest
		ShowHeaderFooter bool
	}{
		GenerateRequest: data,
		// When using Chrome header/footer, we still want page margins set in CSS but not CSS @page header/footer attachments.
		ShowHeaderFooter: false,
	}

	var buf bytes.Buffer
	err = tmpl.ExecuteTemplate(&buf, "base", templateData)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (tb *TemplateBuilder) ExecuteHeaderFooter(headerBuf, footerBuf *bytes.Buffer, data dto.GenerateRequest) error {
	// Execute header template directly - NO building/cloning
	if err := tb.headerTemplate.ExecuteTemplate(headerBuf, "content", data); err != nil {
		headerBuf.Reset()
	}

	// Execute footer template directly - NO building/cloning
	if err := tb.footerTemplate.ExecuteTemplate(footerBuf, "content", data); err != nil {
		footerBuf.Reset()
	}

	return nil
}
