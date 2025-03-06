package pdf

import (
	"testing"

	"github.com/sebsvt/nietzsche/pkg/file"
)

func TestMerge(t *testing.T) {
	type testCase struct {
		name           string
		inputFilePaths []string
		outputFilePath string
		expectedError  error
		expectedOutput bool
	}

	testCases := []testCase{
		{
			name: "merge two pdfs",
			inputFilePaths: []string{
				"../../test_assets/aero.pdf",
				"../../test_assets/aero.pdf",
			},
			outputFilePath: "../../test_assets/output.pdf",
			expectedError:  nil,
			expectedOutput: true,
		},
		{
			name: "Output is empty",
			inputFilePaths: []string{
				"../../test_assets/aero.pdf",
				"../../test_assets/aero.pdf",
			},
			outputFilePath: "",
			expectedError:  ErrOutputFilePathEmpty,
			expectedOutput: false,
		},
		{
			name:           "merge two pdfs with zero input file path",
			inputFilePaths: []string{},
			outputFilePath: "../../test_assets/output.pdf",
			expectedError:  ErrInputFilePathEmpty,
			expectedOutput: false,
		},
		{
			name: "merge two pdfs with invalid output file path",
			inputFilePaths: []string{
				"../../test_assets/not_image.txt",
				"../../test_assets/aero.pdf",
			},
			outputFilePath: "../../test_assets/output.pdf",
			expectedError:  ErrInputFileIsNotPDF,
			expectedOutput: false,
		},
		{
			name: "merge two pdfs with not existing input file path",
			inputFilePaths: []string{
				"../../test_assets/not_existing.pdf",
				"../../test_assets/aero.pdf",
			},
			outputFilePath: "../../test_assets/output.pdf",
			expectedError:  ErrFileNotFound,
			expectedOutput: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := Merge(&MergeParams{
				InputFilePaths: testCase.inputFilePaths,
				OutputFilePath: testCase.outputFilePath,
			})
			if err != testCase.expectedError {
				t.Errorf("expected error %v, got %v", testCase.expectedError, err)
			}

			if testCase.expectedOutput {
				if !file.FileExists(testCase.outputFilePath) {
					t.Errorf("output file %s does not exist", testCase.outputFilePath)
				}
			}
		})
	}
}
