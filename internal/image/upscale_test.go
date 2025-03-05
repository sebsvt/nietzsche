package image

import (
	"os"
	"testing"

	"github.com/sebsvt/nietzsche/pkg/file"
)

func TestUpscale(t *testing.T) {
	type testCase struct {
		name        string
		params      *UpscaleParameters
		expectedErr error
		expectedOut bool
	}
	testcases := []testCase{
		{
			name: "upscale image by 2x",
			params: &UpscaleParameters{
				InPath:     "../../test_assets/meow.jpg",
				OutPath:    "../../test_assets/meow_upscaled.jpg",
				Multiplier: 2,
			},
			expectedErr: nil,
			expectedOut: true,
		},
		{
			name: "upscale image by 4x",
			params: &UpscaleParameters{
				InPath:     "../../test_assets/meow.jpg",
				OutPath:    "../../test_assets/meow_upscaled.jpg",
				Multiplier: 4,
			},
			expectedErr: nil,
			expectedOut: true,
		},
		{
			name: "invalid multiplier should return error",
			params: &UpscaleParameters{
				InPath:     "../../test_assets/meow.jpg",
				OutPath:    "../../test_assets/meow_upscaled.jpg",
				Multiplier: 3,
			},
			expectedErr: ErrInvalidMultiplier,
			expectedOut: false,
		},
		{
			name: "invalid path should return error",
			params: &UpscaleParameters{
				InPath:     "invalid/path/to/file.jpg",
				OutPath:    "invalid/path/to/file_upscaled.jpg",
				Multiplier: 2,
			},
			expectedOut: false,
			expectedErr: ErrCouldNotReadFile,
		},
	}

	for _, testcase := range testcases {
		err := Upscale(testcase.params)
		if err != testcase.expectedErr {
			t.Errorf("expected error: %v, got: %v", testcase.expectedErr, err)
		}
		defer os.Remove(testcase.params.OutPath)
		if testcase.expectedOut {
			has := file.FileExists(testcase.params.OutPath)
			if !has {
				t.Errorf("expected file to be created at: %v, but it was not", testcase.params.OutPath)
			}
		}
	}
}
