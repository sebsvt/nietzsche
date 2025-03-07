package pdf

import (
	"archive/zip"
	"fmt"
	"image/jpeg"
	"io"
	"os"

	"github.com/gen2brain/go-fitz"
	"github.com/sebsvt/nietzsche/pkg/file"
)

type ConvertFromPDFToImageParams struct {
	InputPath  string
	OutputPath string
}

func ConvertFromPDFToImage(params ConvertFromPDFToImageParams) error {
	// Validate parameters
	if params.InputPath == "" {
		return ErrInputFilePathEmpty
	}
	if params.OutputPath == "" {
		return ErrOutputFilePathEmpty
	}
	if !file.FileExists(params.InputPath) {
		return ErrFileNotFound
	}
	if !file.IsValidExtension(params.InputPath, []string{".pdf"}) {
		return ErrInvalidFileExtension
	}
	if !file.IsValidExtension(params.OutputPath, []string{".zip"}) {
		return ErrInvalidFileExtension
	}

	doc, err := fitz.New(params.InputPath)
	if err != nil {
		return fmt.Errorf("failed to open PDF: %w", err)
	}
	defer doc.Close()

	zipFile, err := os.Create(params.OutputPath)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %w", err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for i := 0; i < doc.NumPage(); i++ {
		img, err := doc.Image(i)
		if err != nil {
			return fmt.Errorf("failed to extract image from page %d: %w", i, err)
		}

		filename := fmt.Sprintf("page_%03d.jpg", i+1) // Use consistent naming with padding
		fileWriter, err := zipWriter.Create(filename)
		if err != nil {
			return fmt.Errorf("failed to create file in zip: %w", err)
		}

		pr, pw := io.Pipe()

		encodeErrCh := make(chan error, 1)
		go func() {
			defer pw.Close()
			err := jpeg.Encode(pw, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
			encodeErrCh <- err
		}()
		// Copy data from pipe to zip file
		_, err = io.Copy(fileWriter, pr)
		if err != nil {
			return fmt.Errorf("failed to write image data to zip: %w", err)
		}

		// Check if encoding had an error
		if err := <-encodeErrCh; err != nil {
			return fmt.Errorf("failed to encode image: %w", err)
		}
	}

	return nil
}
