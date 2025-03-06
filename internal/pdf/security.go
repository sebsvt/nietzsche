package pdf

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

type SecurityParams struct {
	InputFilePath  string
	OutputFilePath string
	Password       string
}

func Unlock(params *SecurityParams) error {
	if params == nil {
		return ErrParamsNil
	}
	if params.InputFilePath == "" {
		return ErrInputFilePathEmpty
	}
	if params.OutputFilePath == "" {
		return ErrOutputFilePathEmpty
	}
	if params.Password == "" {
		return ErrPasswordRequired
	}

	conf := model.NewDefaultConfiguration()
	conf.UserPW = params.Password
	conf.OwnerPW = params.Password

	if err := api.DecryptFile(params.InputFilePath, params.OutputFilePath, conf); err != nil {
		return ErrFailedToUnlockFile
	}

	return nil
}

func Protect(params *SecurityParams) error {

	if params == nil {
		return ErrParamsNil
	}
	if params.InputFilePath == "" {
		return ErrInputFilePathEmpty
	}
	if params.OutputFilePath == "" {
		return ErrOutputFilePathEmpty
	}
	if params.Password == "" {
		return ErrPasswordRequired
	}

	conf := model.NewDefaultConfiguration()
	conf.UserPW = params.Password
	conf.OwnerPW = params.Password

	if err := api.EncryptFile(params.InputFilePath, params.OutputFilePath, conf); err != nil {
		return ErrFailedToProtectFile
	}

	return nil
}
