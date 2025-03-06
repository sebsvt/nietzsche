package pdf

import (
	"os"
	"testing"

	"github.com/sebsvt/nietzsche/pkg/file"
)

func TestConvertFromURL(t *testing.T) {
	type testCase struct {
		name        string
		url         string
		outputPath  string
		expectedErr error
		exptectOut  bool
	}

	testCases := []testCase{
		{
			name:        "success",
			url:         "https://www.google.com",
			outputPath:  "../../test_assets/test.pdf",
			expectedErr: nil,
			exptectOut:  true,
		},
		{
			name:        "url is empty string should return an error",
			url:         "",
			outputPath:  "testdata/test.pdf",
			expectedErr: ErrURLRequired,
			exptectOut:  false,
		},
		{
			name:        "output path is empty string should return an error",
			url:         "https://www.google.com",
			outputPath:  "",
			expectedErr: ErrOutputFilePathEmpty,
			exptectOut:  false,
		},
		{
			name:        "not found site should return an error",
			url:         "https://www.notfound-asoajfodsijfaoi.com",
			outputPath:  "../../test_assets/test.pdf",
			expectedErr: ErrFailedToConvertFromURL,
			exptectOut:  false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := ConvertFromURL(&ConverterFromURLParams{
				URL:        testCase.url,
				OutputPath: testCase.outputPath,
			})
			if err != testCase.expectedErr {
				t.Errorf("expected error %v, got %v", testCase.expectedErr, err)
			}
			if testCase.exptectOut && !file.FileExists(testCase.outputPath) {
				t.Errorf("expected output file %v to exist", testCase.outputPath)
			}
			// delete the output file if it exists
			if file.FileExists(testCase.outputPath) {
				err := os.Remove(testCase.outputPath)
				if err != nil {
					t.Errorf("expected output file %v to exist", testCase.outputPath)
				}
			}
		})
	}
}
