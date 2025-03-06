package pdf

import (
	"os"
	"testing"

	"github.com/sebsvt/nietzsche/pkg/file"
)

func TestRotate(t *testing.T) {
	type testCase struct {
		name           string
		angle          int
		inputFilePath  string
		outputFilePath string
		expectedOutput bool
		expectedError  error
	}

	testCases := []testCase{
		{
			name:           "rotate 90 degrees",
			angle:          90,
			inputFilePath:  "../../test_assets/aero.pdf",
			outputFilePath: "../../test_assets/output.pdf",
			expectedOutput: true,
			expectedError:  nil,
		},
		{
			name:           "rotate invalid angle",
			angle:          450,
			inputFilePath:  "../../test_assets/aero.pdf",
			outputFilePath: "../../test_assets/output.pdf",
			expectedOutput: false,
			expectedError:  ErrInvalidAngle,
		},
		{
			name:           "input file is not a pdf",
			angle:          90,
			inputFilePath:  "../../test_assets/not_pdf.txt",
			outputFilePath: "../../test_assets/output.pdf",
			expectedOutput: false,
			expectedError:  ErrInputFileIsNotPDF,
		},
		{
			name:           "file not found",
			angle:          90,
			inputFilePath:  "../../test_assets/not_found.pdf",
			outputFilePath: "../../test_assets/output.pdf",
			expectedOutput: false,
			expectedError:  ErrFileNotFound,
		},
		{
			name:           "empty input file path",
			angle:          90,
			inputFilePath:  "",
			outputFilePath: "../../test_assets/output.pdf",
			expectedOutput: false,
			expectedError:  ErrInputFilePathEmpty,
		},
		{
			name:           "empty output file path",
			angle:          90,
			inputFilePath:  "../../test_assets/aero.pdf",
			outputFilePath: "",
			expectedOutput: false,
			expectedError:  ErrOutputFilePathEmpty,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := Rotate(&RotateParams{
				InputFilePath:  testCase.inputFilePath,
				OutputFilePath: testCase.outputFilePath,
				Angle:          testCase.angle,
			})
			if err != testCase.expectedError {
				t.Errorf("expected error %v, got %v", testCase.expectedError, err)
			}

			// check if the output file is the same as the expected output file
			// if it has then delete it
			if testCase.expectedOutput {
				if !file.FileExists(testCase.outputFilePath) {
					t.Errorf("output file %s does not exist", testCase.outputFilePath)
				}
				os.Remove(testCase.outputFilePath)

			}
		})
	}
}
