package pdf

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

type SplitRemoveParams struct {
	InputPath   string
	OutputPath  string
	RemovePages string
	MergeAfter  bool
}

// SplitRemove removes pages from a PDF, then either merges or zips the output.
func SplitRemove(params *SplitRemoveParams) error {
	if params.InputPath == "" {
		return fmt.Errorf("input file path is empty")
	}
	if params.OutputPath == "" {
		return fmt.Errorf("output file path is empty")
	}
	if params.RemovePages == "" {
		return fmt.Errorf("remove pages field is empty")
	}

	// Parse the pages to remove
	pagesToRemove, err := ParsePageRange(params.RemovePages)
	if err != nil {
		return err
	}

	// Remove specified pages
	if err := api.RemovePagesFile(params.InputPath, params.OutputPath, pagesToRemove, nil); err != nil {
		return fmt.Errorf("failed to remove pages: %w", err)
	}

	// If MergeAfter is enabled, we're done here
	if params.MergeAfter {
		return nil
	}

	// Generate a unique temp directory for split PDFs
	id := uuid.New().String()
	tempDir := filepath.Join(os.TempDir(), "split_pdf_"+id)
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}

	// Split the modified PDF into multiple files
	if err := api.SplitFile(params.OutputPath, tempDir, 1, nil); err != nil {
		return fmt.Errorf("failed to split PDF: %w", err)
	}

	// ✅ FIX: Ensure the ZIP file is saved in a directory, not inside a file
	outputDir := filepath.Dir(params.OutputPath) // Get the directory of OutputPath
	zipFilePath := filepath.Join(outputDir, "split_pdf_"+id+".zip")

	if err := ZipFolder(tempDir, zipFilePath); err != nil {
		return fmt.Errorf("failed to zip split PDFs: %w", err)
	}

	return nil
}

// ZipFolder compresses a folder and its contents into a ZIP file.
func ZipFolder(sourceFolder, zipFileName string) error {
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(sourceFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get the relative path
		relPath, err := filepath.Rel(sourceFolder, path)
		if err != nil {
			return err
		}

		// Create a header for the ZIP entry
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = relPath

		// Set compression for files, store for directories
		if info.IsDir() {
			header.Method = zip.Store
		} else {
			header.Method = zip.Deflate
		}

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// If it's a file, copy its content
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
