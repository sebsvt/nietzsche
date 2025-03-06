package pdf

import (
	"context"
	"errors"
	"os"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type ConverterFromURLParams struct {
	OutputPath string
	URL        string
}

func ConvertFromURL(params *ConverterFromURLParams) error {
	if params.OutputPath == "" {
		return errors.New("output path is required")
	}

	if params.URL == "" {
		return errors.New("url is required")
	}
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var pdfBuffer []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(params.URL),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfBuffer, _, err = page.PrintToPDF().Do(ctx)
			return err
		}),
	)
	if err != nil {
		return err
	}

	err = os.WriteFile(params.OutputPath, pdfBuffer, 0644)
	if err != nil {
		return err
	}
	return nil
}

func FromImage(image string) error {

	return nil
}

func FromMarkdown(markdown string) error {

	return nil
}

func FromBase64(base64 string) error {

	return nil
}

func FromPDF(pdf string) error {

	return nil
}
