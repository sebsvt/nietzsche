package image

import (
	"testing"

	"github.com/sebsvt/nietzsche/pkg/file"
)

func TestRotate(t *testing.T) {
	type testCase struct {
		name        string
		params      *RotateParameters
		expectedErr error
		expectedOut bool
	}
	testcases := []testCase{
		{
			name: "rotate image by 90 degrees",
			params: &RotateParameters{
				InPath:  "../../test_assets/meow.jpg",
				OutPath: "../../test_assets/meow_rotated.jpg",
				Angle:   90,
			},
			expectedErr: nil,
			expectedOut: true,
		},
		{
			name: "rotate image by 180 degrees",
			params: &RotateParameters{
				InPath:  "../../test_assets/meow.jpg",
				OutPath: "../../test_assets/meow_rotated.jpg",
				Angle:   180,
			},
			expectedErr: nil,
			expectedOut: true,
		},
		{
			name: "not image file should return error",
			params: &RotateParameters{
				InPath:  "../../test_assets/not_image.txt",
				OutPath: "../../test_assets/meow_rotated.jpg",
				Angle:   90,
			},
			expectedErr: ErrCouldNotDecodeFile,
			expectedOut: false,
		},
		{
			name: "invalid angle should return error",
			params: &RotateParameters{
				InPath:  "../../test_assets/meow.jpg",
				OutPath: "../../test_assets/meow_rotated.jpg",
				Angle:   3600,
			},
			expectedErr: ErrInvalidAngle,
			expectedOut: false,
		},
		{
			name: "invalid path should return error",
			params: &RotateParameters{
				InPath:  "invalid/path/to/file.jpg",
				OutPath: "invalid/path/to/file_rotated.jpg",
				Angle:   90,
			},
			expectedErr: ErrCouldNotReadFile,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			err := Rotate(testcase.params)
			defer file.DeleteFile(testcase.params.OutPath)
			if err != testcase.expectedErr {
				t.Errorf("expected error %v, got %v", testcase.expectedErr, err)
			}
			if testcase.expectedOut && !file.FileExists(testcase.params.OutPath) {
				t.Errorf("expected file %v to exist", testcase.params.OutPath)
			}
		})
	}
}
