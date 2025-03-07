package pdf

type SplitFixedRangeParams struct {
	InputPath  string
	OutputPath string
	Range      string
	MergeAfter bool
}

// the range is only a single number means for every n number of pages, split the pdf
// like we devide all pages into groups of n and then split the pdf into n files
// if pages amount is 10 and range is 2, we devide pages into groups of 2 and then split the pdf into 5 files
