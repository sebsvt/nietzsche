package pdf

import "fmt"

var (
	ErrParamsNil           = fmt.Errorf("params are nil")
	ErrInputFilePathEmpty  = fmt.Errorf("input file path is empty")
	ErrOutputFilePathEmpty = fmt.Errorf("output file path is empty")
	ErrPasswordRequired    = fmt.Errorf("password is required")
	ErrFailedToProtectFile = fmt.Errorf("failed to protect file")
	ErrFailedToUnlockFile  = fmt.Errorf("failed to unlock file")
)
