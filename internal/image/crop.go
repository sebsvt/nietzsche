package image

import (
	"bytes"
	"image"

	"github.com/disintegration/imaging"
	"github.com/sebsvt/nietzsche/pkg/file"
)

type CropParameters struct {
	InPath  string
	OutPath string
	Width   int
	Height  int
	X       int
	Y       int
}

func Crop(params *CropParameters) error {
	if params.Width <= 0 || params.Height <= 0 {
		return ErrWidthOrHeightZero
	}

	// optional but should not less than zero
	if params.X < 0 || params.Y < 0 {
		return ErrCoordinatesInvalid
	}

	f, err := file.ReadFile(params.InPath)
	if err != nil {
		return ErrCouldNotReadFile
	}
	_img, err := imaging.Decode(bytes.NewReader(f.Content))
	if err != nil {
		return ErrCouldNotDecodeFile
	}
	img := imaging.Crop(_img, image.Rectangle{
		Min: image.Point{X: params.X, Y: params.Y},
		Max: image.Point{X: params.X + params.Width, Y: params.Y + params.Height},
	})

	err = imaging.Save(img, params.OutPath)
	if err != nil {
		return ErrCouldNotSaveFile
	}

	return nil
}
