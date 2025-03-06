package image

import (
	"os"
	"testing"

	"github.com/sebsvt/nietzsche/pkg/file"
)

func TestCrop(t *testing.T) {
	type testCase struct {
		name        string
		params      *CropParameters
		expectedErr error
		expectedOut bool
	}

	testCases := []testCase{
		{
			name: "crop image by width and height",
			params: &CropParameters{
				InPath:  "../../test_assets/cat.jpg",
				OutPath: "../../test_assets/cat_cropped.jpg",
				Width:   100,
				Height:  100,
				X:       0,
				Y:       0,
			},
			expectedErr: nil,
			expectedOut: true,
		},
		{
			name: "invalide path should return an error",
			params: &CropParameters{
				InPath:  "",
				OutPath: "",
				Width:   100,
				Height:  100,
				X:       0,
				Y:       0,
			},
			expectedErr: ErrCouldNotReadFile,
			expectedOut: false,
		},
		{
			name: "width or height is zero should return an error",
			params: &CropParameters{
				InPath:  "../../test_assets/cat.jpg",
				OutPath: "../../test_assets/cat_cropped.jpg",
				Width:   0,
				Height:  0,
				X:       0,
				Y:       0,
			},
			expectedErr: ErrWidthOrHeightZero,
			expectedOut: false,
		},
		{
			name: "coordinates are negative should return an error",
			params: &CropParameters{
				InPath:  "../../test_assets/cat.jpg",
				OutPath: "../../test_assets/cat_cropped.jpg",
				Width:   100,
				Height:  100,
				X:       -1,
				Y:       -1,
			},
			expectedErr: ErrCoordinatesInvalid,
			expectedOut: false,
		},
	}

	for _, tc := range testCases {
		err := Crop(tc.params)
		if err != tc.expectedErr {
			t.Errorf("expected error: %v, got: %v", tc.expectedErr, err)
		}
		// check if the file has been created at the out path
		if tc.expectedOut {
			has := file.FileExists(tc.params.OutPath)
			if !has {
				t.Errorf("expected file to be created at: %v, but it was not", tc.params.OutPath)
			}
		}
		// delete the file after the test
		if tc.expectedOut {
			os.Remove(tc.params.OutPath)
		}
	}

}
