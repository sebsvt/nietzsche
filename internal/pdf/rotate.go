package pdf

import "github.com/pdfcpu/pdfcpu/pkg/api"

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

	err := api.RotateFile(params.InputFilePath, params.OutputFilePath, params.Angle, nil, nil)
	if err != nil {
		return ErrFailedToRotate
	}

	return nil
}
