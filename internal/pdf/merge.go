package pdf

import (
	"fmt"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/sebsvt/nietzsche/pkg/file"
)

type MergeParams struct {
	InputFilePaths []string
	OutputFilePath string
}

func Merge(params *MergeParams) error {

	if len(params.InputFilePaths) == 0 {
		return ErrInputFilePathEmpty
	}

	if params.OutputFilePath == "" {
		return ErrOutputFilePathEmpty
	}

	for _, inputFilePath := range params.InputFilePaths {
		if !file.IsValidExtension(inputFilePath, []string{".pdf"}) {
			return ErrInputFileIsNotPDF
		}

		if !file.FileExists(inputFilePath) {
			return ErrFileNotFound
		}
	}

	// merge
	if err := api.MergeCreateFile(params.InputFilePaths, params.OutputFilePath, false, nil); err != nil {
		return fmt.Errorf("%w: %v", ErrFailedToMerge, err)
	}

	return nil
}
