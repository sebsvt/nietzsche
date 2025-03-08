package pdf

import (
	"log"

	"github.com/sebsvt/nietzsche/pkg/file"
)

type AddWatermarkParams struct {
	InputPath  string
	OutputPath string
	Text       string
}

func AddWatermark(params *AddWatermarkParams) error {
	if params == nil {
		return ErrParamsNil
	}

	if params.InputPath == "" {
		return ErrInputFilePathEmpty
	}

	if params.OutputPath == "" {
		return ErrOutputFilePathEmpty
	}

	if params.Text == "" {
		return ErrTextEmpty
	}

	if !file.IsValidExtension(params.InputPath, []string{".pdf"}) {
		return ErrInvalidFileExtension
	}

	if !file.IsValidExtension(params.OutputPath, []string{".pdf"}) {
		return ErrInvalidFileExtension
	}
	log.Println("Adding watermark to PDF...")
	return nil
}
