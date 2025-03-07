package pdf

import (
	"os"
	"testing"

	"github.com/sebsvt/nietzsche/pkg/file"
)

func TestOCRScan(t *testing.T) {
	type testCase struct {
		name           string
		inputPath      string
		outputPath     string
		expectedError  error
		expectedOutput bool
	}
	tests := []testCase{
		{
			name:           "input path is empty",
			inputPath:      "",
			outputPath:     "",
			expectedError:  ErrInputFilePathEmpty,
			expectedOutput: false,
		},
		{
			name:           "output path is empty",
			inputPath:      "../../test_assets/page_002.pdf",
			outputPath:     "",
			expectedError:  ErrOutputFilePathEmpty,
			expectedOutput: false,
		},
		{
			name:           "invalid input file extension",
			inputPath:      "../../test_assets/page_002.jpg",
			outputPath:     "../../test_assets/page_002_ocr_square.pdf",
			expectedError:  ErrInvalidFileExtension,
			expectedOutput: false,
		},
		{
			name:           "invalid output file extension",
			inputPath:      "../../test_assets/page_002.pdf",
			outputPath:     "../../test_assets/page_002_ocr_square.jpg",
			expectedError:  ErrInvalidFileExtension,
			expectedOutput: false,
		},
		{
			name:           "valid but does not exist",
			inputPath:      "../../test_assets/page_002_does_not_exist.pdf",
			outputPath:     "../../test_assets/page_002_ocr_square.pdf",
			expectedError:  ErrFileNotFound,
			expectedOutput: false,
		},
		{
			name:           "ocr scan",
			inputPath:      "../../test_assets/page_002.pdf",
			outputPath:     "../../test_assets/page_002_ocr_square.pdf",
			expectedError:  nil,
			expectedOutput: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := OCRScan(&OCRScanParams{
				InputPath:  test.inputPath,
				OutputPath: test.outputPath,
			})
			if err != test.expectedError {
				t.Errorf("expected error %v, got %v", test.expectedError, err)
			}
			// if expectedOutput is true, check that the output file exists and then delete it
			if test.expectedOutput {
				if !file.FileExists(test.outputPath) {
					t.Errorf("expected output file %v to exist", test.outputPath)
				}
				os.Remove(test.outputPath)
			}
		})
	}
}
