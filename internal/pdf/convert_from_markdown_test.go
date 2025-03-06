package pdf

import (
	"os"
	"testing"
)

func TestFromMarkdown(t *testing.T) {
	type testCase struct {
		name           string
		inputPath      string
		outputPath     string
		expectedErr    error
		expectedOutput bool
	}
	testCases := []testCase{
		{
			name:           "valid markdown file",
			inputPath:      "../../test_assets/markdown_test.md",
			outputPath:     "../../test_assets/markdown_test.pdf",
			expectedErr:    nil,
			expectedOutput: true,
		},
		{
			name:           "not existing file",
			inputPath:      "../../test_assets/markdown_test_invalid.md",
			outputPath:     "../../test_assets/markdown_test_invalid.pdf",
			expectedErr:    ErrFailedToReadFile,
			expectedOutput: false,
		},
		{
			name:           "not markdown file",
			inputPath:      "../../test_assets/cat.jpg",
			outputPath:     "../../test_assets/cat.pdf",
			expectedErr:    ErrInvalidMarkdown,
			expectedOutput: false,
		},
		{
			name:           "empty output path",
			inputPath:      "../../test_assets/markdown_test.md",
			outputPath:     "",
			expectedErr:    ErrOutputFilePathEmpty,
			expectedOutput: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := FromMarkdown(&ConverterFromMarkdownParams{
				InputPath:  tc.inputPath,
				OutputPath: tc.outputPath,
			})
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
			if tc.expectedOutput {
				if _, err := os.Stat(tc.outputPath); os.IsNotExist(err) {
					t.Errorf("expected output file %s to exist", tc.outputPath)
				}
			}
			if !tc.expectedOutput {
				if _, err := os.Stat(tc.outputPath); !os.IsNotExist(err) {
					t.Errorf("expected output file %s to not exist", tc.outputPath)
				}
			}
		})
	}
}
