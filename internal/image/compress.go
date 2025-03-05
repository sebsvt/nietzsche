package image

import (
	"bytes"
	"image/png"
	"os"

	"github.com/disintegration/imaging"
	"github.com/sebsvt/nietzsche/pkg/file"
)

type CompressionLevel string

const (
	CompressionLevelLow         = "low"
	CompressionLevelRecommended = "recommended"
	CompressionLevelExtreme     = "extreme"
)

type CompressParameters struct {
	InPath           string
	OutPath          string
	CompressionLevel CompressionLevel
}

func Compress(params *CompressParameters) (string, error) {
	var level png.CompressionLevel

	if err := validateCompressionLevel(params.CompressionLevel); err != nil {
		return "", err
	}

	f, err := file.ReadFile(params.InPath)
	if err != nil {
		return "", ErrCouldNotReadFile
	}

	img, err := imaging.Decode(bytes.NewReader(f.Content))
	if err != nil {
		return "", ErrCouldNotDecodeFile
	}

	// match the compression level
	if params.CompressionLevel == CompressionLevelLow {
		level = png.DefaultCompression
	} else if params.CompressionLevel == CompressionLevelRecommended {
		level = png.BestSpeed
	} else if params.CompressionLevel == CompressionLevelExtreme {
		level = png.BestCompression
	}

	buff := bytes.NewBuffer(nil)
	err = imaging.Encode(buff, img, imaging.PNG, imaging.PNGCompressionLevel(level))
	if err != nil {
		return "", ErrCouldNotEncodeFile
	}

	err = os.WriteFile(params.OutPath, buff.Bytes(), 0644)
	if err != nil {
		return "", ErrCouldNotSaveFile
	}

	return params.OutPath, nil
}

func validateCompressionLevel(level CompressionLevel) error {
	switch level {
	case CompressionLevelLow, CompressionLevelRecommended, CompressionLevelExtreme:
		return nil
	}
	return ErrInvalidCompressionLevel
}
