package services

import (
	"context"
	"strings"

	"github.com/chromedp/cdproto/page"
)
import "github.com/chromedp/chromedp"

func Generate(ctx context.Context) ([]byte, error) {
	html := `
			<!DOCTYPE html>
			<html>
			<head><title>PDF Test</title></head>
			<body><h1>Hello PDF</h1><p>Generated from template.</p></body>
			</html>
		`
	dataURL := "data:text/html," + strings.ReplaceAll(html, "\n", "")

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var pdfBuf []byte
	err := chromedp.Run(ctx,
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
