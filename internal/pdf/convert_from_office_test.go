package pdf

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sebsvt/nietzsche/pkg/file"
)

func TestFromOffice(t *testing.T) {
	type testCase struct {
		name        string
		inputPath   string
		outputDir   string
		expectedOut bool
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "success",
			inputPath:   "../../test_assets/sample.docx",
			outputDir:   "../../test_assets",
			expectedOut: true,
			expectedErr: nil,
		},
		{
			name:        "unsupported file",
			inputPath:   "../../test_assets/cat.jpg",
			outputDir:   "../../test_assets",
			expectedOut: false,
			expectedErr: ErrInvalidOffice,
		},
		{
			name:        "empty input path",
			inputPath:   "",
			outputDir:   "../../test_assets",
			expectedOut: false,
			expectedErr: ErrInputFilePathEmpty,
		},
		{
			name:        "empty output dir",
			inputPath:   "../../test_assets/sample.docx",
			outputDir:   "",
			expectedOut: false,
			expectedErr: ErrOutputDirEmpty,
		},
		{
			name:        "failed to run command, since invalid input path",
			inputPath:   "../../test_assets/invalid_file_not_found.docx",
			outputDir:   "../../test_assets",
			expectedOut: false,
			expectedErr: ErrFileNotFound,
		},
	}

	for _, testCase := range testCases {
		err := FromOffice(&ConverterFromOfficeParams{
			InputPath: testCase.inputPath,
			OutputDir: testCase.outputDir,
		})
		if testCase.expectedErr != err {
			t.Errorf("expected error %v, but got %v", testCase.expectedErr, err)
		}
		// check if the file exists if it has then delete it
		if testCase.expectedOut {
			// filename is the input path with changed extension to pdf
			extension := filepath.Ext(testCase.inputPath)
			filename := strings.Replace(testCase.inputPath, extension, ".pdf", 1)
			if file.FileExists(filename) {
				err := os.Remove(filename)
				if err != nil {
					t.Errorf("expected file to exist in %v, but it does not", testCase.outputDir)
				}
			}
		}
	}
}
