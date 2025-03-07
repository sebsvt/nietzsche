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
	if !file.IsValidExtension(params.InputPath, []string{".pdf"}) {
		return ErrInvalidFileExtension
	}
	if !file.IsValidExtension(params.OutputPath, []string{".zip"}) {
		return ErrInvalidFileExtension
	}
	if !file.FileExists(params.InputPath) {
		return ErrFileNotFound
	}

	doc, err := fitz.New(params.InputPath)
	if err != nil {
		return ErrFailedToOpenPDF
	}
	defer doc.Close()

	zipFile, err := os.Create(params.OutputPath)
	if err != nil {
		return ErrFailedToCreateZipFile
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for i := 0; i < doc.NumPage(); i++ {
		img, err := doc.Image(i)
		if err != nil {
			return ErrFailedToExtractImage
		}

		filename := fmt.Sprintf("page_%03d.jpg", i+1) // Use consistent naming with padding
		fileWriter, err := zipWriter.Create(filename)
		if err != nil {
			return ErrFailedToCreateFileInZip
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
			return ErrFailedToWriteImageData
		}

		// Check if encoding had an error
		if err := <-encodeErrCh; err != nil {
			return ErrFailedToEncodeImage
		}
	}

	return nil
}
