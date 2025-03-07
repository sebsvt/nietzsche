package pdf

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/sebsvt/nietzsche/pkg/file"
)

var (
	ErrInputFilePathEmpty   = fmt.Errorf("input file path is empty")
	ErrOutputFilePathEmpty  = fmt.Errorf("output file path is empty")
	ErrFileNotFound         = fmt.Errorf("file not found")
	ErrInvalidFileExtension = fmt.Errorf("invalid file extension")
	ErrInvalidConformance   = fmt.Errorf("invalid PDF/A conformance level")
	ErrValidationFailed     = fmt.Errorf("PDF/A validation failed")
	ErrConversionFailed     = fmt.Errorf("PDF/A conversion failed")
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
	// Validate input parameters
	if params.InputPath == "" {
		return ErrInputFilePathEmpty
	}
	if params.OutputPath == "" {
		return ErrOutputFilePathEmpty
	}
	if !file.FileExists(params.InputPath) {
		return ErrFileNotFound
	}
	if !isValidConformanceLevel(params.Conformance) {
		return fmt.Errorf("%w: %s", ErrInvalidConformance, params.Conformance)
	}

	// Create a temporary directory for processing
	tempDir, err := os.MkdirTemp("", "pdf-pdfa-validate-")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a configuration for PDF/A validation and conversion
	config := model.NewDefaultConfiguration()
	config.ValidationMode = model.ValidationRelaxed

	// Step 1: Validate if the PDF is already in PDF/A format with the desired conformance
	isValid, err := validatePDFA(params.InputPath, params.Conformance, config)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	// If already valid, just copy to output
	if isValid {
		if err := copyFile(params.InputPath, params.OutputPath); err != nil {
			return fmt.Errorf("failed to copy valid PDF/A: %w", err)
		}
		return nil
	}

	// Step 2: If not valid, attempt conversion
	conformanceLevels := []string{params.Conformance}

	// If downgrade is allowed, add lower conformance levels
	if params.AllowDowngrade {
		conformanceLevels = getDowngradePath(params.Conformance)
	}

	var conversionErr error
	tempOutput := filepath.Join(tempDir, "converted.pdf")

	// Try each conformance level in order until one succeeds
	for _, level := range conformanceLevels {
		conversionErr = convertToPDFA(params.InputPath, tempOutput, level, config)
		if conversionErr == nil {
			// Conversion succeeded, validate the result
			isValid, err := validatePDFA(tempOutput, level, config)
			if err != nil {
				continue // Validation had an error, try next level
			}
			if isValid {
				// Copy the validated PDF/A to the output path
				if err := copyFile(tempOutput, params.OutputPath); err != nil {
					return fmt.Errorf("failed to copy converted PDF/A: %w", err)
				}
				return nil
			}
		}
	}

	// If we reach here, all conversion attempts failed
	return fmt.Errorf("%w: tried all conformance levels", ErrConversionFailed)
}

// validatePDFA checks if a PDF file conforms to a specific PDF/A standard
func validatePDFA(filePath, conformance string, config *model.Configuration) (bool, error) {
	// Parse the conformance level to pdfcpu format
	pdfaLevel, err := parsePDFALevel(conformance)
	if err != nil {
		return false, err
	}

	// Validate the PDF
	valid, err := api.ValidateFileWithConfig(filePath, config)
	if err != nil {
		// If error contains specific validation errors, it's not valid
		return false, nil
	}

	if !valid {
		return false, nil
	}

	// For PDF/A validation, we need to check additional properties
	// This requires inspecting the PDF structure
	ctx, err := api.ReadContextFile(filePath)
	if err != nil {
		return false, err
	}

	// Attempt to get PDF/A conformance from the document
	if ctx.XRefTable.Version < pdfaLevel.PDFVersion() {
		return false, nil // PDF version too low for specified conformance
	}

	// A complete validation would require checking metadata, fonts, color spaces, etc.
	// For a comprehensive check, a specialized library might be needed

	return true, nil
}

// convertToPDFA converts a PDF to PDF/A format with the specified conformance level
func convertToPDFA(inputPath, outputPath, conformance string, config *model.Configuration) error {
	// Parse the conformance level
	pdfaLevel, err := parsePDFALevel(conformance)
	if err != nil {
		return err
	}

	// Set conversion options
	cmd := &pdfcpu.Command{
		Mode:    pdfcpu.PDFA,
		InFile:  inputPath,
		OutFile: outputPath,
	}
	cmd.PDFALevel = pdfaLevel

	// Perform the conversion
	_, err = api.Process(cmd)
	if err != nil {
		return fmt.Errorf("PDF/A conversion failed: %w", err)
	}

	return nil
}

// parsePDFALevel converts a string conformance level to pdfcpu format
func parsePDFALevel(conformance string) (pdfcpu.PDFALevel, error) {
	switch strings.ToLower(conformance) {
	case "pdfa-1b":
		return pdfcpu.PDFA1B, nil
	case "pdfa-1a":
		return pdfcpu.PDFA1A, nil
	case "pdfa-2b":
		return pdfcpu.PDFA2B, nil
	case "pdfa-2u":
		return pdfcpu.PDFA2U, nil
	case "pdfa-2a":
		return pdfcpu.PDFA2A, nil
	case "pdfa-3b":
		return pdfcpu.PDFA3B, nil
	case "pdfa-3u":
		return pdfcpu.PDFA3U, nil
	case "pdfa-3a":
		return pdfcpu.PDFA3A, nil
	default:
		return pdfcpu.PDFA1B, fmt.Errorf("unsupported PDF/A conformance level: %s", conformance)
	}
}

// isValidConformanceLevel checks if the provided conformance level is supported
func isValidConformanceLevel(level string) bool {
	level = strings.ToLower(level)
	for _, validLevel := range ValidPDFAConformanceLevels {
		if level == validLevel {
			return true
		}
	}
	return false
}

// getDowngradePath returns all conformance levels at or below the specified level
// in descending order of strictness
func getDowngradePath(targetLevel string) []string {
	level := strings.ToLower(targetLevel)
	var levels []string

	// PDF/A conformance hierarchy (from most strict to least strict)
	conformanceHierarchy := []string{
		"pdfa-3a", "pdfa-3u", "pdfa-3b",
		"pdfa-2a", "pdfa-2u", "pdfa-2b",
		"pdfa-1a", "pdfa-1b",
	}

	// Find the target level in the hierarchy
	found := false
	for _, l := range conformanceHierarchy {
		if found || l == level {
			found = true
			levels = append(levels, l)
		}
	}

	return levels
}

// Helper function to copy a file
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = os.Stat(src)
	if err != nil {
		return fmt.Errorf("source file error: %w", err)
	}

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return destFile.Sync()
}
