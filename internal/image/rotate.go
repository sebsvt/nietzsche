package image

import (
	"os"

	"github.com/disintegration/imaging"
)

type RotateParameters struct {
	InPath  string
	OutPath string
	Angle   float64
}

func Rotate(params *RotateParameters) error {
	// validate angle
	if params.Angle < -360 || params.Angle > 360 {
		return ErrInvalidAngle
	}

	file, err := os.Open(params.InPath)
	if err != nil {
		return ErrCouldNotReadFile
	}
	defer file.Close()

	img, err := imaging.Decode(file)
	if err != nil {
		return ErrCouldNotDecodeFile
	}

	rotated := imaging.Rotate(img, params.Angle, nil)

	err = imaging.Save(rotated, params.OutPath)
	if err != nil {
		return ErrCouldNotSaveFile
	}

	return nil
}
