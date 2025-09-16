package services

import (
	"bytes"
	"context"
	"os"
	"time"
	"zabdoc/internal/api/dto"
	"zabdoc/internal/templates"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/sony/gobreaker"
)

// PDFService handles PDF generation from HTML.
type PDFService struct {
	ctx context.Context
	cb  *gobreaker.CircuitBreaker
}

func NewPDFService() *PDFService {
	// options required to run chrome in heroku container stack
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-setuid-sandbox", true),
	)

	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, _ := chromedp.NewContext(allocCtx)

	// TODO: send notification to maintainer(s)
	// if chromedp crashes multiple times then prevent further requests
	// from going through until the issue is fixed
	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "chromedp-pdf",
		MaxRequests: 3,                // max requests allowed in half-open state before tripping again.
		Interval:    60 * time.Second, // interval to reset internal counts (failures/successes).
		Timeout:     2 * time.Minute,  // how long to wait before transitioning from open to half-open state.
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// Trip the breaker after 5 consecutive failures.
			return counts.ConsecutiveFailures >= 5
		},
	})

	return &PDFService{
		ctx: ctx,
		cb:  cb,
	}
}

func (ps *PDFService) GeneratePDF(data dto.GenerateRequest) ([]byte, error) {
	result, err := ps.cb.Execute(func() (interface{}, error) {
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
	})
	if err != nil {
		return nil, err
	}
	return result.([]byte), nil
}
