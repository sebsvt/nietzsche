package pdf

import "fmt"

var (
	ErrParamsNil              = fmt.Errorf("params are nil")
	ErrInputFilePathEmpty     = fmt.Errorf("input file path is empty")
	ErrOutputFilePathEmpty    = fmt.Errorf("output file path is empty")
	ErrPasswordRequired       = fmt.Errorf("password is required")
	ErrFailedToProtectFile    = fmt.Errorf("failed to protect file")
	ErrFailedToUnlockFile     = fmt.Errorf("failed to unlock file")
	ErrFailedToConvert        = fmt.Errorf("failed to convert")
	ErrURLRequired            = fmt.Errorf("url is required")
	ErrFailedToConvertFromURL = fmt.Errorf("failed to convert from url")
	ErrFailedToWriteFile      = fmt.Errorf("failed to write file")
	ErrFailedToReadOrWrite    = fmt.Errorf("failed to read or write")
	ErrInvalidOrientation     = fmt.Errorf("invalid orientation")
	ErrInvalidPageSize        = fmt.Errorf("invalid page size")
	ErrInvalidMargin          = fmt.Errorf("invalid margin")
)
