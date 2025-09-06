package services

import (
	"bytes"
	"context"
	"log"
	"net/url"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/hammadmajid/zabcover/internal/api/dto"
	"github.com/hammadmajid/zabcover/internal/utils"
)

type assignmentService struct {
	logger *log.Logger
}

func newAssignmentService(l *log.Logger) *assignmentService {
	return &assignmentService{
		logger: l,
	}
}

func (as *assignmentService) Generate(assignment dto.AssignmentRequest) ([]byte, error) {
	var buf bytes.Buffer
	if err := utils.GetTemplate(utils.Assignment).Execute(&buf, assignment); err != nil {
		return nil, err
	}

	html := buf.String()
	dataURL := "data:text/html," + url.PathEscape(html)

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var pdfBuf []byte

	err := chromedp.Run(ctx,
		// Load HTML
		chromedp.Navigate(dataURL),

		// Force viewport size (avoid default 800x600)
		emulation.SetDeviceMetricsOverride(1280, 1024, 1.0, false),

		// Use print media emulation
		emulation.SetEmulatedMedia(),

		// Allow layout to settle
		chromedp.Sleep(500*time.Millisecond),

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
		as.logger.Printf("Error generating pdf: %v", err)
		return nil, err
	}

	return pdfBuf, nil
}
