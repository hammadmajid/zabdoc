package controllers

import (
	"bytes"
	"context"
	"log"
	"os"
	"time"

	"zabdoc/internal/api/dto"
	"zabdoc/internal/templates"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type labTaskCtrl struct {
	logger *log.Logger
}

func newLabTaskCtrl(l *log.Logger) *labTaskCtrl {
	return &labTaskCtrl{
		logger: l,
	}
}

func (as *labTaskCtrl) Generate(labTask dto.LabTaskRequest) ([]byte, error) {
	var buf bytes.Buffer
	if err := templates.GetTemplate(templates.LabTask).Execute(&buf, labTask); err != nil {
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
		as.logger.Printf("Error generating pdf: %v", err)
		return nil, err
	}

	return pdfBuf, nil
}
