package image

import (
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/disintegration/imaging"
)

type ConvertParameters struct {
	InPath  string
	OutPath string
	Format  string
}

func Convert(params *ConvertParameters) (string, error) {
	file, err := os.Open(params.InPath)
	if err != nil {
		return "", ErrCouldNotReadFile
	}
	defer file.Close()

	img, err := imaging.Decode(file)
	if err != nil {
		return "", ErrCouldNotDecodeFile
	}

	outFile, err := os.Create(params.OutPath)
	if err != nil {
		return "", ErrCouldNotCreateFile
	}
	defer outFile.Close()

	switch params.Format {
	case "jpg", "jpeg":
		if err := jpeg.Encode(outFile, img, &jpeg.Options{Quality: 90}); err != nil {
			return "", ErrCouldNotEncodeFile
		}
	case "png":
		if err := png.Encode(outFile, img); err != nil {
			return "", ErrCouldNotEncodeFile
		}
	case "gif", "gif_animation":
		if err := gif.Encode(outFile, img, nil); err != nil {
			return "", ErrCouldNotEncodeFile
		}
	default:
		return "", ErrInvalidFormat
	}

	return params.OutPath, nil
}
