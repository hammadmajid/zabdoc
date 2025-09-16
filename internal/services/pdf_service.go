package services

import (
	"bytes"
	"context"
	"os"
	"zabdoc/internal/api/dto"
	"zabdoc/internal/templates"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// PDFService handles PDF generation from HTML.
type PDFService struct {
	ctx context.Context
}

func NewPDFService() *PDFService {
	// options required to run chrome in heroku container stack
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-setuid-sandbox", true),
	)

	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	ctx, _ := chromedp.NewContext(allocCtx)

	return &PDFService{
		ctx: ctx,
	}
}

func (ps *PDFService) GeneratePDF(data dto.GenerateRequest) ([]byte, error) {
	var buf bytes.Buffer
	if err := templates.Tpl.Execute(&buf, data); err != nil {
		return nil, err
	}

	tmpFile, err := os.CreateTemp("", "section_*.html")
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

	var pdfBuf []byte

	err = chromedp.Run(ps.ctx,
		chromedp.Navigate(fileURL),
		emulation.SetDeviceMetricsOverride(1280, 1024, 1.0, false),
		emulation.SetEmulatedMedia(),
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
		return nil, err
	}

	return pdfBuf, nil
}
