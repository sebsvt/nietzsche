package pdf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/raykov/gofpdf"
	"github.com/raykov/mdtopdf"
)

type ConverterFromMarkdownParams struct {
	InputPath  string
	OutputPath string
}

func FromMarkdown(params *ConverterFromMarkdownParams) error {
	if params == nil {
		return ErrParamsNil
	}

	if params.InputPath == "" {
		return ErrInputFilePathEmpty
	}

	if params.OutputPath == "" {
		return ErrOutputFilePathEmpty
	}
	if !isMarkdownFile(params.InputPath) {
		return ErrInvalidMarkdown
	}

	md, err := os.Open(params.InputPath)
	if err != nil {
		return ErrFailedToReadFile
	}
	defer md.Close()

	pdf, err := os.Create(params.OutputPath)
	if err != nil {
		return ErrFailedToCreateFile
	}
	defer pdf.Close()

	pageNumExtension := func(pdf *gofpdf.Fpdf) {
		pdf.SetFooterFunc(func() {
			left, _, right, bottom := pdf.GetMargins()
			width, height := pdf.GetPageSize()
			fontSize := 12.0

			pNum := fmt.Sprint(pdf.PageNo())
			pdf.SetXY(width-left/2-pdf.GetStringWidth(pNum), height-bottom/2)
			pdf.SetFontSize(fontSize)
			pdf.SetTextColor(200, 200, 200)
			pdf.SetFontStyle("B")
			pdf.SetRightMargin(0)
			pdf.Write(fontSize, pNum)
			pdf.SetRightMargin(right)
		})
	}

	if err = mdtopdf.Convert(md, pdf, pageNumExtension); err != nil {
		return ErrFailedToConvert
	}
	return nil
}

func isMarkdownFile(filePath string) bool {
	ext := filepath.Ext(filePath)
	return ext == ".md" || ext == ".markdown"
}
