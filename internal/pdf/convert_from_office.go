package pdf

import (
	"os/exec"

	"github.com/sebsvt/nietzsche/pkg/file"
)

type ConverterFromOfficeParams struct {
	InputPath string
	OutputDir string
}

// supported files: word, powerpoint, excel
var officeFileExtensions = []string{".docx", ".doc", ".pptx", ".ppt", ".xlsx", ".xls"}

func FromOffice(params *ConverterFromOfficeParams) error {
	if params == nil {
		return ErrParamsNil
	}
	if params.InputPath == "" {
		return ErrInputFilePathEmpty
	}
	if params.OutputDir == "" {
		return ErrOutputDirEmpty
	}
	if !file.IsValidExtension(params.InputPath, officeFileExtensions) {
		return ErrInvalidOffice
	}
	// file must exist
	if !file.FileExists(params.InputPath) {
		return ErrFileNotFound
	}

	cmd := exec.Command("soffice", "--convert-to", "pdf", "--outdir", params.OutputDir, params.InputPath)
	if err := cmd.Run(); err != nil {
		return ErrFailedToRunCommand
	}

	return nil
}
