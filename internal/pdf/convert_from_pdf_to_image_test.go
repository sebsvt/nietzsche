package pdf

import (
	"os"
	"testing"

	"github.com/sebsvt/nietzsche/pkg/file"
)

func TestConvertFromPDFToImage(t *testing.T) {
	type testCase struct {
		name           string
		inputPath      string
		outputPath     string
		expectedError  error
		expectedOutput bool
	}

	testCases := []testCase{
		{
			name:           "invalid input file type",
			inputPath:      "test_assets/test.txt",
			outputPath:     "test_assets/test.zip",
			expectedError:  ErrInvalidFileExtension,
			expectedOutput: false,
		},
		{
			name:           "invalid output file type",
			inputPath:      "test_assets/test.pdf",
			outputPath:     "test_assets/test.txt",
			expectedError:  ErrInvalidFileExtension,
			expectedOutput: false,
		},
		{
			name:           "empty input path",
			inputPath:      "",
			outputPath:     "test_assets/test.zip",
			expectedError:  ErrInputFilePathEmpty,
			expectedOutput: false,
		},
		{
			name:           "empty output path",
			inputPath:      "test_assets/test.pdf",
			outputPath:     "",
			expectedError:  ErrOutputFilePathEmpty,
			expectedOutput: false,
		},
		{
			name:           "input file does not exist",
			inputPath:      "test_assets/does_not_exist.pdf",
			outputPath:     "test_assets/test.zip",
			expectedError:  ErrFileNotFound,
			expectedOutput: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := ConvertFromPDFToImage(ConvertFromPDFToImageParams{
				InputPath:  testCase.inputPath,
				OutputPath: testCase.outputPath,
			})
			if err != testCase.expectedError {
				t.Errorf("expected error %v, got %v", testCase.expectedError, err)
			}
			// if expectedOutput is true, check if the output file exists and then delete it
			if testCase.expectedOutput {
				if !file.FileExists(testCase.outputPath) {
					t.Errorf("expected output file %v to exist", testCase.outputPath)
				}
				os.Remove(testCase.outputPath)
			}
		})
	}
}
