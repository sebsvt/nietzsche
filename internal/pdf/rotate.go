package pdf

import (
	"fmt"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/sebsvt/nietzsche/pkg/file"
)

type RotateParams struct {
	InputFilePath  string
	OutputFilePath string
	Angle          int
}

func Rotate(params *RotateParams) error {
	if params == nil {
		return ErrParamsNil
	}

	if params.InputFilePath == "" {
		return ErrInputFilePathEmpty
	}

	if params.OutputFilePath == "" {
		return ErrOutputFilePathEmpty
	}

	// between 0 and 360
	if params.Angle < 0 || params.Angle > 360 {
		return ErrInvalidAngle
	}

	if !file.IsValidExtension(params.InputFilePath, []string{".pdf"}) {
		return ErrInputFileIsNotPDF
	}

	if !file.FileExists(params.InputFilePath) {
		return ErrFileNotFound
	}

	err := api.RotateFile(params.InputFilePath, params.OutputFilePath, params.Angle, nil, nil)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrFailedToRotate, err)
	}

	return nil
}
