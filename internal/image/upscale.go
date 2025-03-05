package image

import (
	"bytes"

	"github.com/disintegration/imaging"
	"github.com/sebsvt/nietzsche/pkg/file"
)

type UpscaleParameters struct {
	InPath     string
	OutPath    string
	Multiplier int
}

func Upscale(params *UpscaleParameters) (string, error) {
	// img, err := imaging.Decode(bytes.NewReader(f.Content))
	f, err := file.ReadFile(params.InPath)
	if err != nil {
		return "", ErrCouldNotReadFile
	}
	img, err := imaging.Decode(bytes.NewReader(f.Content))
	if err != nil {
		return "", ErrCouldNotDecodeFile
	}
	// scale image up
	width := img.Bounds().Dx() * params.Multiplier
	height := img.Bounds().Dy() * params.Multiplier
	img = imaging.Resize(img, width, height, imaging.CatmullRom)

	err = imaging.Save(img, params.OutPath)
	if err != nil {
		return "", ErrCouldNotSaveFile
	}

	return params.OutPath, nil
}
