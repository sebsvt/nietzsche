package pdf

import (
	"github.com/jung-kurt/gofpdf"
)

type ConverterFromImageParams struct {
	InputPath   string
	OutputPath  string
	Orientation string
	Margin      int
	PageSize    string
}

func FromImage(params *ConverterFromImageParams) error {
	orientation, err := validateOrientation(params.Orientation)
	if err != nil {
		return err
	}
	pageSize, err := validatePageSize(params.PageSize)
	if err != nil {
		return err
	}
	if err := validateMargin(params.Margin); err != nil {
		return err
	}

	// from image to pdf
	pdf := gofpdf.New(orientation, "mm", pageSize, "")
	pdf.AddPage()
	// adapt margin to page size
	floatMargin := float64(params.Margin)
	pdf.Image(params.InputPath, floatMargin, floatMargin, 210-floatMargin*2, 297-floatMargin*2, false, "", 0, "")
	if err := pdf.OutputFileAndClose(params.OutputPath); err != nil {
		return ErrFailedToReadOrWrite
	}
	return nil
}

func validateOrientation(orientation string) (string, error) {
	if orientation != "Portrait" && orientation != "Landscape" {
		return "", ErrInvalidOrientation
	}
	if orientation == "Portrait" {
		return "P", nil
	}
	return "L", nil
}

func validatePageSize(pageSize string) (string, error) {
	if pageSize != "A4" && pageSize != "fit" && pageSize != "letter" {
		return "", ErrInvalidPageSize
	}
	// since gofpdf doesn't support fit, we need to calculate the size of the page
	if pageSize == "fit" {
		return "A4", nil
	}
	return pageSize, nil
}

func validateMargin(margin int) error {
	if margin < 0 {
		return ErrInvalidMargin
	}
	return nil
}
