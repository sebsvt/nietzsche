package pdf

import (
	"context"
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
		return ErrOutputFilePathEmpty
	}

	if params.URL == "" {
		return ErrURLRequired
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
		return ErrFailedToConvertFromURL
	}

	err = os.WriteFile(params.OutputPath, pdfBuffer, 0644)
	if err != nil {
		return ErrFailedToWriteFile
	}
	return nil
}
