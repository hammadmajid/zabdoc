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
	ctx             context.Context
	cancel          context.CancelFunc
	allocCtx        context.Context
	allocCancel     context.CancelFunc
	cb              *gobreaker.CircuitBreaker
	templateBuilder *templates.TemplateBuilder
}

func NewPDFService() *PDFService {
	// options required to run chrome in heroku container stack
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-setuid-sandbox", true),
	)

	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocCtx)

	// Initialize template builder
	templateBuilder, err := templates.NewTemplateBuilder()
	if err != nil {
		panic(err) // Handle this better in production
	}

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
		ctx:             ctx,
		cancel:          cancel,
		allocCtx:        allocCtx,
		allocCancel:     allocCancel,
		cb:              cb,
		templateBuilder: templateBuilder,
	}
}

// GeneratePDFWithConfig generates a PDF with custom configuration
func (ps *PDFService) GeneratePDF(data dto.GenerateRequest, config templates.TemplateConfig) ([]byte, error) {
	result, err := ps.cb.Execute(func() (any, error) {
		htmlBytes, err := ps.templateBuilder.Execute(config, data)
		if err != nil {
			return nil, err
		}

		// Render header and footer templates directly using the cached templates
		var headerBuf, footerBuf bytes.Buffer
		if err := ps.templateBuilder.ExecuteHeaderFooter(&headerBuf, &footerBuf, data); err != nil {
			return nil, err
		}

		tmpFile, err := os.CreateTemp("", "section_*.html")
		if err != nil {
			return nil, err
		}
		defer os.Remove(tmpFile.Name())
		defer tmpFile.Close()

		if _, err := tmpFile.Write(htmlBytes); err != nil {
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
					WithDisplayHeaderFooter(true).
					WithHeaderTemplate(headerBuf.String()).
					WithFooterTemplate(footerBuf.String()).
					WithMarginTop(0.6).
					WithMarginBottom(0.6).
					WithMarginLeft(0.25).
					WithMarginRight(0.25).
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

func (ps *PDFService) Shutdown() {
	if ps.cancel != nil {
		ps.cancel()
	}
	if ps.allocCancel != nil {
		ps.allocCancel()
	}
}
