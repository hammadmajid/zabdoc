package http

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"zabdoc/internal/api/dto"
	"zabdoc/internal/templates"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/yuin/goldmark"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger) Handler {
	return Handler{
		logger: logger,
	}
}

func (h Handler) Generate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		h.logger.Printf("parse multipart form: %v", err)
		return
	}

	data := dto.GenerateRequest{
		StudentName: r.FormValue("studentName"),
		RegNo:       r.FormValue("regNo"),
		Class:       r.FormValue("class"),
		Course:      r.FormValue("course"),
		CourseCode:  r.FormValue("courseCode"),
		Instructor:  r.FormValue("instructor"),
		DocType:     r.FormValue("type"),
		Number:      r.FormValue("number"),
		Date:        r.FormValue("date"),
	}

	// Process sections
	var sections []dto.Section
	for i := 0; i < 15; i++ { // Max 15 sections as per frontend
		contentKey := fmt.Sprintf("section-%d-content", i)
		filesKey := fmt.Sprintf("section-%d-files", i)

		content := r.FormValue(contentKey)
		if content == "" {
			continue // Skip empty sections
		}

		// Convert markdown to HTML
		var buf bytes.Buffer
		if err := goldmark.Convert([]byte(content), &buf); err != nil {
			panic(err)
		}

		// TODO: input sanitization

		task := dto.Section{
			Index:   i,
			Content: buf.String(),
		}

		// Process uploaded images
		if files := r.MultipartForm.File[filesKey]; files != nil {
			for _, fileHeader := range files {
				file, err := fileHeader.Open()
				if err != nil {
					h.logger.Printf("Error opening file: %v", err)
					continue
				}
				defer file.Close()

				// Detect MIME type
				ext := filepath.Ext(fileHeader.Filename)
				mimeType := mime.TypeByExtension(ext)
				if mimeType == "" {
					mimeType = "image/jpeg" // fallback
				}

				// Read file content
				fileData, err := io.ReadAll(file)
				if err != nil {
					h.logger.Printf("Error reading file: %v", err)
					continue
				}

				if len(fileData) == 0 {
					h.logger.Printf("Empty image file: %s", fileHeader.Filename)
					continue
				}

				// Encode as base64
				base64Data := base64.StdEncoding.EncodeToString(fileData)

				// Store as TaskImage with base64 data and MIME type
				taskImage := dto.SectionImage{
					Data:     base64Data,
					MimeType: mimeType,
				}
				task.Images = append(task.Images, taskImage)

				h.logger.Printf("Image processed: %s (%d bytes → %d base64 chars)",
					mimeType, len(fileData), len(base64Data))
			}
		}

		sections = append(sections, task)
	}

	data.Sections = sections

	pdf, err := h.generate(data)
	if err != nil {
		http.Error(w, "failed to generate PDF", http.StatusInternalServerError)
		h.logger.Printf("generate PDF: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=lab-task.pdf")
	if _, writeErr := w.Write(pdf); writeErr != nil {
		h.logger.Printf("write PDF: %v", writeErr)
	}
}

func (h *Handler) generate(data dto.GenerateRequest) ([]byte, error) {

	var buf bytes.Buffer
	if err := templates.Tpl.Execute(&buf, data); err != nil {
		return nil, err
	}

	// Write to temporary file
	tmpFile, err := os.CreateTemp("", "labtask_*.html")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	if _, err := tmpFile.Write(buf.Bytes()); err != nil {
		return nil, err
	}
	tmpFile.Close()

	fileURL := "file://" + tmpFile.Name()

	// Force Chrome to run without sandbox (needed on Heroku)
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-setuid-sandbox", true),
	)

	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer allocCancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var pdfBuf []byte

	err = chromedp.Run(ctx,
		// Load HTML file
		chromedp.Navigate(fileURL),

		// Force viewport size (avoid default 800x600)
		emulation.SetDeviceMetricsOverride(1280, 1024, 1.0, false),

		// Use print media emulation
		emulation.SetEmulatedMedia(),

		// Allow time for images to load
		chromedp.Sleep(2*time.Second),

		// Render PDF with strict params
		chromedp.ActionFunc(func(ctx context.Context) error {
			pdf, _, err := page.PrintToPDF().
				WithPrintBackground(true).
				WithScale(1).
				WithPreferCSSPageSize(true).
				WithMarginTop(0).
				WithMarginBottom(0).
				WithMarginLeft(0).
				WithMarginRight(0).
				Do(ctx)
			if err != nil {
				return err
			}
			pdfBuf = pdf
			return nil
		}),
	)
	if err != nil {
		h.logger.Printf("Error generating pdf: %v", err)
		return nil, err
	}

	return pdfBuf, nil
}

//goland:noinspection ALL
func (h Handler) Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available")
}
