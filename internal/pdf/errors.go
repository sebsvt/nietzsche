package pdf

import (
	"errors"
	"fmt"
)

var (
	ErrParamsNil                 = fmt.Errorf("params are nil")
	ErrInputFilePathEmpty        = fmt.Errorf("input file path is empty")
	ErrOutputFilePathEmpty       = fmt.Errorf("output file path is empty")
	ErrPasswordRequired          = fmt.Errorf("password is required")
	ErrFailedToProtectFile       = fmt.Errorf("failed to protect file")
	ErrFailedToUnlockFile        = fmt.Errorf("failed to unlock file")
	ErrFailedToConvert           = fmt.Errorf("failed to convert")
	ErrURLRequired               = fmt.Errorf("url is required")
	ErrFailedToConvertFromURL    = fmt.Errorf("failed to convert from url")
	ErrFailedToWriteFile         = fmt.Errorf("failed to write file")
	ErrFailedToReadFile          = fmt.Errorf("failed to read file")
	ErrFailedToReadOrWrite       = fmt.Errorf("failed to read or write")
	ErrInvalidOrientation        = fmt.Errorf("invalid orientation")
	ErrInvalidPageSize           = fmt.Errorf("invalid page size")
	ErrInvalidMargin             = fmt.Errorf("invalid margin")
	ErrInvalidMarkdown           = fmt.Errorf("invalid markdown file")
	ErrInvalidOffice             = fmt.Errorf("invalid office file")
	ErrFailedToCreateFile        = fmt.Errorf("failed to create file")
	ErrFailedToRunCommand        = fmt.Errorf("failed to run command")
	ErrOutputDirEmpty            = fmt.Errorf("output directory is empty")
	ErrFileNotFound              = fmt.Errorf("file not found")
	ErrInputFileIsNotPDF         = fmt.Errorf("input file is not a PDF")
	ErrOutputFileIsNotPDF        = fmt.Errorf("output file is not a PDF")
	ErrPDFAParamsNil             = errors.New("PDF/A parameters cannot be nil")
	ErrPDFAInputPathEmpty        = errors.New("PDF/A input file path cannot be empty")
	ErrPDFAOutputPathEmpty       = errors.New("PDF/A output file path cannot be empty")
	ErrInvalidPDFAFormat         = errors.New("invalid PDF/A format")
	ErrFailedToPDFAConvert       = errors.New("failed to convert to PDF/A")
	ErrAngleZero                 = fmt.Errorf("angle cannot be zero")
	ErrInvalidAngle              = fmt.Errorf("invalid angle")
	ErrFailedToRotate            = fmt.Errorf("failed to rotate")
	ErrFailedToMerge             = fmt.Errorf("failed to merge")
	ErrInvalidFileExtension      = fmt.Errorf("invalid file extension")
	ErrFailedToOpenPDF           = fmt.Errorf("failed to open PDF")
	ErrFailedToCreateZipFile     = fmt.Errorf("failed to create zip file")
	ErrFailedToWriteImageData    = fmt.Errorf("failed to write image data")
	ErrFailedToEncodeImage       = fmt.Errorf("failed to encode image")
	ErrFailedToExtractImage      = fmt.Errorf("failed to extract image")
	ErrFailedToCreateFileInZip   = fmt.Errorf("failed to create file in zip")
	ErrOCRMyPDFIsNotInstalled    = fmt.Errorf("ocrmypdf is not installed")
	ErrGhostscriptIsNotInstalled = fmt.Errorf("ghostscript is not installed")
	ErrSOfficeIsNotInstalled     = fmt.Errorf("soffice is not installed")
)
