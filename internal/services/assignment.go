package services

import (
	"bytes"
	"context"
	"strings"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/hammadmajid/zabcover/internal/utils"
)

func Generate(ctx context.Context) ([]byte, error) {
	var buf bytes.Buffer
	err := utils.TemplateFiles[utils.Assignment].Execute(&buf, nil)
	if err != nil {
		return nil, err
	}

	html := buf.String()
	dataURL := "data:text/html," + strings.ReplaceAll(html, "\n", "")

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var pdfBuf []byte

	err = chromedp.Run(ctx,
		chromedp.Navigate(dataURL),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().Do(ctx)
			if err != nil {
				return err
			}
			pdfBuf = buf
			return nil
		}),
	)
	if err != nil {
		return nil, err
	}

	return pdfBuf, nil
}
