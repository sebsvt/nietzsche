package pdf

type SplitParams struct {
	InputPath  string
	OutputPath string
}

func Split(params *SplitParams) error {
	if params.InputPath == "" {
		return ErrInputFilePathEmpty
	}
	if params.OutputPath == "" {
		return ErrOutputFilePathEmpty
	}
	return nil
}
