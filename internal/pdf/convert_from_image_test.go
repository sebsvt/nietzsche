package pdf

import (
	"os"
	"testing"

	"github.com/sebsvt/nietzsche/pkg/file"
)

func TestFromImage(t *testing.T) {
	type testCase struct {
		name        string
		inputPath   string
		outputPath  string
		orientation string
		pageSize    string
		margin      int
		expectedErr error
		exptectOut  bool
	}
	testCases := []testCase{
		{
			name:        "portrait, A4, 0 margin, fit page size",
			inputPath:   "../../test_assets/meow.jpg",
			outputPath:  "../../test_assets/meow.pdf",
			orientation: "Portrait",
			pageSize:    "fit",
			margin:      0,
			expectedErr: nil,
			exptectOut:  true,
		},
		{
			name:        "landscape, A4, 0 margin, fit page size",
			inputPath:   "../../test_assets/meow.jpg",
			outputPath:  "../../test_assets/meow.pdf",
			orientation: "Landscape",
			pageSize:    "fit",
			margin:      0,
			expectedErr: nil,
		},
		{
			name:        "invalid orientation",
			inputPath:   "../../test_assets/meow.jpg",
			outputPath:  "../../test_assets/meow.pdf",
			orientation: "Invalid",
			pageSize:    "fit",
			margin:      0,
			expectedErr: ErrInvalidOrientation,
		},
		{
			name:        "invalid page size",
			inputPath:   "../../test_assets/meow.jpg",
			outputPath:  "../../test_assets/meow.pdf",
			orientation: "Portrait",
			pageSize:    "Invalid",
			margin:      0,
			expectedErr: ErrInvalidPageSize,
		},
		{
			name:        "invalid margin",
			inputPath:   "../../test_assets/meow.jpg",
			outputPath:  "../../test_assets/meow.pdf",
			orientation: "Portrait",
			pageSize:    "A4",
			margin:      -1,
			expectedErr: ErrInvalidMargin,
		},
		{
			name:        "invalid path",
			inputPath:   "../../test_assets/invalid.jpg",
			outputPath:  "../../test_assets/meow.pdf",
			orientation: "Portrait",
			pageSize:    "A4",
			margin:      0,
			expectedErr: ErrFailedToReadOrWrite,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := FromImage(&ConverterFromImageParams{
				InputPath:   testCase.inputPath,
				OutputPath:  testCase.outputPath,
				Orientation: testCase.orientation,
				PageSize:    testCase.pageSize,
				Margin:      testCase.margin,
			})
			if err != testCase.expectedErr {
				t.Errorf("expected error %v, got %v", testCase.expectedErr, err)
			}
			if testCase.exptectOut && !file.FileExists(testCase.outputPath) {
				t.Errorf("expected output file %v to exist", testCase.outputPath)
			}
			// delete the output file if it exists
			if file.FileExists(testCase.outputPath) {
				os.Remove(testCase.outputPath)
			}
		})
	}
}
