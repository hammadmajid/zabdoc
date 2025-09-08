package controllers

import (
	"bytes"
	"context"
	"log"
	"net/url"
	"time"

	"zabdoc/internal/api/dto"
	"zabdoc/internal/utils"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type assignmentCtrl struct {
	logger *log.Logger
}

func newAssignmentCtrl(l *log.Logger) *assignmentCtrl {
	return &assignmentCtrl{
		logger: l,
	}
}

func (as *assignmentCtrl) Generate(assignment dto.AssignmentRequest) ([]byte, error) {
	var buf bytes.Buffer
	if err := utils.GetTemplate(utils.Assignment).Execute(&buf, assignment); err != nil {
		return nil, err
	}

	html := buf.String()
	dataURL := "data:text/html," + url.PathEscape(html)

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
