package pdf

import (
	"os"
	"testing"
)

func TestFromPDFToPDFA(t *testing.T) {
	type testCase struct {
		name           string
		inputPath      string
		outputPath     string
		format         string
		expectedErr    error
		expectedOutput bool
	}

	testCases := []testCase{
		{
			name:           "Valid PDF/A-1B conversion",
			inputPath:      "../../test_assets/aero.pdf",
			outputPath:     "../../test_assets/test_pdfa1b.pdf",
			format:         "pdfa-1b",
			expectedErr:    nil,
			expectedOutput: true,
		},
		{
			name:           "Invalid input file path",
			inputPath:      "../../test_assets/non_existent.pdf",
			outputPath:     "../../test_assets/test_pdfa1b.pdf",
			format:         "pdfa-1b",
			expectedErr:    ErrFileNotFound,
			expectedOutput: false,
		},
		{
			name:           "Unsupported PDF/A format",
			inputPath:      "../../test_assets/aero.pdf",
			outputPath:     "../../test_assets/test_invalid.pdf",
			format:         "pdfa-99z",
			expectedErr:    ErrInvalidPDFAFormat,
			expectedOutput: false,
		},
		{
			name:           "empty input path",
			inputPath:      "",
			outputPath:     "../../test_assets/test_pdfa1b.pdf",
			format:         "pdfa-1b",
			expectedErr:    ErrPDFAInputPathEmpty,
			expectedOutput: false,
		},
		{
			name:           "empty output path",
			inputPath:      "../../test_assets/aero.pdf",
			outputPath:     "",
			format:         "pdfa-1b",
			expectedErr:    ErrPDFAOutputPathEmpty,
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := FromPDFToPDFA(&FromPDFToPDFAParams{
				InputPath:  tc.inputPath,
				OutputPath: tc.outputPath,
				Format:     tc.format,
			})
			if tc.expectedErr != err {
				t.Errorf("expected error: %v, got: %v", tc.expectedErr, err)
			}

			// if expectedOutput is true, check if the output file exists and delete it
			if tc.expectedOutput {
				if _, err := os.Stat(tc.outputPath); os.IsNotExist(err) {
					t.Errorf("expected output file to exist: %s", tc.outputPath)
				}
				os.Remove(tc.outputPath)
			}

		})
	}
}
