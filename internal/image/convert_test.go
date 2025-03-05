package image

import (
	"testing"

	"github.com/sebsvt/nietzsche/pkg/file"
)

func TestConvert(t *testing.T) {
	type testCase struct {
		name          string
		inPath        string
		outPath       string
		format        string
		expectedError error
		expectedOut   bool
	}

	testcases := []testCase{
		{
			name:          "convert jpg to png",
			inPath:        "../../test_assets/meow.jpg",
			outPath:       "../../test_assets/meow.png",
			format:        "png",
			expectedError: nil,
			expectedOut:   true,
		},
		{
			name:          "convert jpg to png with invalid path",
			inPath:        "../../test_assets/invalid.jpg",
			outPath:       "../../test_assets/meow.png",
			format:        "png",
			expectedError: ErrCouldNotReadFile,
			expectedOut:   false,
		},
		{
			name:          "convert jpg to png with invalid format",
			inPath:        "../../test_assets/meow.jpg",
			outPath:       "../../test_assets/meow.apfoie",
			format:        "apfoie",
			expectedError: ErrInvalidFormat,
			expectedOut:   false,
		},
		{
			name:          "convert jpg to gif",
			inPath:        "../../test_assets/meow.jpg",
			outPath:       "../../test_assets/meow.gif",
			format:        "gif",
			expectedError: nil,
			expectedOut:   true,
		},
		{
			name:          "invaide file input",
			inPath:        "../../test_assets/invalid.apfoie",
			outPath:       "../../test_assets/meow.gif",
			format:        "gif",
			expectedError: ErrCouldNotReadFile,
			expectedOut:   false,
		},
		{
			name:          "not image file should return error",
			inPath:        "../../test_assets/not_image.txt",
			outPath:       "../../test_assets/meow.gif",
			format:        "gif",
			expectedError: ErrCouldNotDecodeFile,
			expectedOut:   false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			err := Convert(&ConvertParameters{InPath: testcase.inPath, OutPath: testcase.outPath, Format: testcase.format})
			defer file.DeleteFile(testcase.outPath)
			if err != testcase.expectedError {
				t.Errorf("expected error %v, got %v", testcase.expectedError, err)
			}
			if testcase.expectedOut && !file.FileExists(testcase.outPath) {
				t.Errorf("expected file %v to exist", testcase.outPath)
			}
		})
	}
}
