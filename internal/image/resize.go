package image

import (
	"bytes"
	"image"

	"github.com/disintegration/imaging"
	"github.com/sebsvt/nietzsche/pkg/file"
)

type ResizeParameters struct {
	InPath             string
	OutPath            string
	ResizeMode         string
	PixelWidth         int
	PixelHeight        int
	Percentage         int
	MaintainRatio      bool
	NoEnlargeIfSmaller bool
}

// return file path and error
func Resize(params *ResizeParameters) error {
	var image *image.NRGBA

	img, err := file.ReadFile(params.InPath)
	if err != nil {
		return ErrCouldNotReadFile
	}
	_image, err := imaging.Decode(bytes.NewReader(img.Content))
	if err != nil {
		return ErrCouldNotDecodeFile
	}
	// validate parameters
	if err := validateResizeParameters(params); err != nil {
		return err
	}

	switch params.ResizeMode {
	// resize image by percentage
	case "percentage":
		if params.Percentage > 0 {
			// calculate the new width and height
			newWidth := _image.Bounds().Dx() * params.Percentage / 100
			newHeight := _image.Bounds().Dy() * params.Percentage / 100
			image = imaging.Resize(_image, newWidth, newHeight, imaging.CatmullRom)
		}

	// resize image by width and height (pixels)
	case "pixel":
		image = imaging.Resize(_image, params.PixelWidth, params.PixelHeight, imaging.CatmullRom)

	default:
		return ErrInvalidResizeMode
	}

	// save image
	err = imaging.Save(image, params.OutPath)
	if err != nil {
		return ErrCouldNotSaveFile
	}

	return nil
}

func validateResizeParameters(params *ResizeParameters) error {
	if params.ResizeMode != "percentage" && params.ResizeMode != "pixel" {
		params.ResizeMode = "pixel"
	}

	if params.ResizeMode == "percentage" && params.Percentage <= 0 {
		return ErrInvalidPercentage
	}

	if params.ResizeMode == "pixel" && (params.PixelWidth <= 0 || params.PixelHeight <= 0) {
		return ErrInvalidPixel
	}

	return nil
}
