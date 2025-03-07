package pdf

import (
	"fmt"
)

var (
	ErrInvalidConformance = fmt.Errorf("invalid PDF/A conformance level")
	ErrValidationFailed   = fmt.Errorf("PDF/A validation failed")
	ErrConversionFailed   = fmt.Errorf("PDF/A conversion failed")
)

// ValidPDFAConformanceLevels contains all valid PDF/A conformance levels
var ValidPDFAConformanceLevels = []string{
	"pdfa-1b", "pdfa-1a",
	"pdfa-2b", "pdfa-2u", "pdfa-2a",
	"pdfa-3b", "pdfa-3u", "pdfa-3a",
}

// ValidatePDFAParams contains parameters for PDF/A validation and conversion
type ValidatePDFAParams struct {
	InputPath      string
	OutputPath     string
	Conformance    string
	AllowDowngrade bool
}

// ValidatePDFA validates and/or converts a PDF file to PDF/A format
func ValidatePDFA(params *ValidatePDFAParams) error {
	return nil
}
