package pdf

import (
	"fmt"
	"os/exec"

	"github.com/sebsvt/nietzsche/pkg/file"
)

type OCRScanParams struct {
	InputPath  string
	OutputPath string
}

func OCRScan(params *OCRScanParams) error {
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
	// check that ocrmypdf is installed
	if _, err := exec.LookPath("ocrmypdf"); err != nil {
		return ErrOCRMyPDFIsNotInstalled
	}

	cmd := exec.Command("ocrmypdf", params.InputPath, params.OutputPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run ocrmypdf: %w", err)
	}

	return nil
}
