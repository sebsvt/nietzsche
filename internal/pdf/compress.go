package pdf

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/sebsvt/nietzsche/pkg/file"
)

type CompressionLevel string

const (
	CompressionLevelLow         CompressionLevel = "low"
	CompressionLevelRecommended CompressionLevel = "recommended"
	CompressionLevelExtreme     CompressionLevel = "extreme"
)

type CompressParams struct {
	InputPath        string
	OutputPath       string
	CompressionLevel CompressionLevel
}

func Compress(params *CompressParams) error {
	if params.InputPath == "" {
		return ErrInputFilePathEmpty
	}
	if params.OutputPath == "" {
		return ErrOutputFilePathEmpty
	}
	if !file.IsValidExtension(params.InputPath, []string{".pdf"}) {
		return ErrInvalidFileExtension
	}
	if !file.IsValidExtension(params.OutputPath, []string{".pdf"}) {
		return ErrInvalidFileExtension
	}
	if !file.FileExists(params.InputPath) {
		return ErrFileNotFound
	}

	// compress the pdf by 3 levels: low, recommended, extreme
	conf := model.NewDefaultConfiguration()

	switch params.CompressionLevel {
	case CompressionLevelLow:
		conf.Optimize = true
		conf.OptimizeDuplicateContentStreams = false
		conf.OptimizeResourceDicts = false

	case CompressionLevelExtreme:
		conf.Optimize = true
		conf.OptimizeDuplicateContentStreams = true
		conf.OptimizeResourceDicts = true

	default:
		// Default to reccomended
		conf.Optimize = true
		conf.OptimizeDuplicateContentStreams = true
		conf.OptimizeResourceDicts = false
	}

	if err := api.OptimizeFile(params.InputPath, params.OutputPath, conf); err != nil {
		return ErrFailedToCompress
	}

	return nil
}
