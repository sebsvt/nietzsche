package pdf

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/sebsvt/nietzsche/pkg/file"
)

var supportedPDFAFormats = map[string][]string{
	"pdfa-1b": {"-dPDFA=1", "-dPDFACompatibilityPolicy=1"},
	"pdfa-1a": {"-dPDFA=1", "-dPDFACompatibilityPolicy=1", "-dPDFALevel=1", "-dPDFAType=1"},
	"pdfa-2b": {"-dPDFA=2", "-dPDFACompatibilityPolicy=1"},
	"pdfa-2u": {"-dPDFA=2", "-dPDFACompatibilityPolicy=1", "-dPDFALevel=2", "-dPDFAType=2"},
	"pdfa-2a": {"-dPDFA=2", "-dPDFACompatibilityPolicy=1", "-dPDFALevel=2", "-dPDFAType=1"},
	"pdfa-3b": {"-dPDFA=3", "-dPDFACompatibilityPolicy=1"},
	"pdfa-3u": {"-dPDFA=3", "-dPDFACompatibilityPolicy=1", "-dPDFALevel=3", "-dPDFAType=2"},
	"pdfa-3a": {"-dPDFA=3", "-dPDFACompatibilityPolicy=1", "-dPDFALevel=3", "-dPDFAType=1"},
}

type FromPDFToPDFAParams struct {
	InputPath  string // Path to input PDF file
	OutputPath string // Path to output PDF/A file
	Format     string // PDF/A format (e.g., "pdfa-1b", "pdfa-2b", etc.)
}

func FromPDFToPDFA(params *FromPDFToPDFAParams) error {
	// Validate parameters
	if params == nil {
		return ErrPDFAParamsNil
	}
	if params.InputPath == "" {
		return ErrPDFAInputPathEmpty
	}
	if params.OutputPath == "" {
		return ErrPDFAOutputPathEmpty
	}
	if !file.IsValidExtension(params.InputPath, []string{".pdf"}) {
		return ErrInputFileIsNotPDF
	}
	if !file.FileExists(params.InputPath) {
		return ErrFileNotFound
	}
	if _, exists := supportedPDFAFormats[params.Format]; !exists {
		return ErrInvalidPDFAFormat
	}

	// Base arguments for all formats
	args := []string{
		"-dBATCH",
		"-dNOPAUSE",
		"-sDEVICE=pdfwrite",
		"-sColorConversionStrategy=UseDeviceIndependentColor",
		"-sOutputFile=" + params.OutputPath,
	}

	// Append format-specific arguments
	args = append(args, supportedPDFAFormats[params.Format]...)
	args = append(args, params.InputPath)

	// Execute Ghostscript command
	var stderr bytes.Buffer
	cmd := exec.Command("gs", args...)
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%w: %s", ErrFailedToPDFAConvert, stderr.String())
	}

	return nil
}
