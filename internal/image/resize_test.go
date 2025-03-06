package image

import (
	"testing"

	"github.com/sebsvt/nietzsche/pkg/file"
)

func TestResize(t *testing.T) {
	type testCase struct {
		name        string
		params      *ResizeParameters
		expectedErr error
		expectedOut bool
	}

	testCases := []testCase{
		{
			name: "not exist file path should return an error as can't read file",
			params: &ResizeParameters{
				InPath:  "invalid/path/to/file.jpg",
				OutPath: "invalid/path/to/file_resized.jpg",
			},
			expectedErr: ErrCouldNotReadFile,
			expectedOut: false,
		},
		{
			name: "invalide file type should return an error",
			params: &ResizeParameters{
				InPath:  "../../test_assets/not_image.txt",
				OutPath: "../../test_assets/not_image_resized.txt",
			},
			expectedErr: ErrCouldNotDecodeFile,
			expectedOut: false,
		},
		{
			name: "resize image by percentage",
			params: &ResizeParameters{
				ResizeMode: "percentage",
				Percentage: 100,
				InPath:     "../../test_assets/cat.jpg",
				OutPath:    "../../test_assets/cat_resized.jpg",
			},
			expectedErr: nil,
			expectedOut: true,
		},
		{
			name: "resize image by percentage but invalide or not provided percentage",
			params: &ResizeParameters{
				ResizeMode:  "percentage",
				InPath:      "../../test_assets/cat.jpg",
				OutPath:     "../../test_assets/cat_resized.jpg",
				PixelWidth:  100,
				PixelHeight: 100,
			},
			expectedOut: false,
			expectedErr: ErrInvalidPercentage,
		},
		{
			name: "resize image by width and height",
			params: &ResizeParameters{
				ResizeMode:  "pixel",
				PixelWidth:  100,
				PixelHeight: 100,
				InPath:      "../../test_assets/cat.jpg",
				OutPath:     "../../test_assets/cat_resized.jpg",
			},
			expectedErr: nil,
			expectedOut: true,
		},
		{
			name: "resize image by width and height but invalide or not provided width and height",
			params: &ResizeParameters{
				ResizeMode: "pixel",
				InPath:     "../../test_assets/cat.jpg",
				OutPath:    "../../test_assets/cat_resized.jpg",
			},
			expectedOut: false,
			expectedErr: ErrInvalidPixel,
		},
	}

	for _, testCase := range testCases {
		err := Resize(testCase.params)
		if err != testCase.expectedErr {
			t.Errorf("expected error: %v, got: %v", testCase.expectedErr, err)
		}
		// check if the file has been created at the out path
		has := file.FileExists(testCase.params.OutPath)
		if !has && testCase.expectedOut {
			t.Errorf("expected file to be created at: %v, but it was not", testCase.params.OutPath)
		}
		// delete the file after the test
		if testCase.expectedOut {
			err = file.DeleteFile(testCase.params.OutPath)
			if err != nil {
				t.Fatal("failed to delete file in test case")
			}
		}
	}
}

func TestValidateResizeParameters(t *testing.T) {
	type testCase struct {
		name         string
		params       *ResizeParameters
		expectedMode string
		expectedErr  error
	}

	testCases := []testCase{
		{
			name: "percentage mode with valid percentage",
			params: &ResizeParameters{
				ResizeMode: "percentage",
				Percentage: 100,
			},
			expectedMode: "percentage",
			expectedErr:  nil,
		},
		{
			name: "percentage mode with invalid percentage",
			params: &ResizeParameters{
				ResizeMode: "percentage",
				Percentage: 0,
			},
			expectedMode: "percentage",
			expectedErr:  ErrInvalidPercentage,
		},
		{
			name: "pixel mode with invalid pixel width and height",
			params: &ResizeParameters{
				ResizeMode:  "pixel",
				PixelWidth:  0,
				PixelHeight: 0,
			},
			expectedMode: "pixel",
			expectedErr:  ErrInvalidPixel,
		},
		{
			name: "pixel mode with invalid pixel width",
			params: &ResizeParameters{
				ResizeMode: "pixel",
				PixelWidth: 0,
			},
			expectedMode: "pixel",
			expectedErr:  ErrInvalidPixel,
		},
		{
			name: "pixel mode with invalid pixel height",
			params: &ResizeParameters{
				ResizeMode:  "pixel",
				PixelHeight: 0,
			},
			expectedMode: "pixel",
			expectedErr:  ErrInvalidPixel,
		},
		{
			name: "invalid resize mode should be set to pixel as a default",
			params: &ResizeParameters{
				ResizeMode:  "invalid",
				PixelWidth:  100,
				PixelHeight: 100,
			},
			expectedMode: "pixel",
			expectedErr:  nil,
		},
		{
			name: "invaide mode is set to be pixel but if width and height are not set it should return an error",
			params: &ResizeParameters{
				ResizeMode: "pixel",
			},
			expectedMode: "pixel",
			expectedErr:  ErrInvalidPixel,
		},
	}

	for _, testCase := range testCases {
		err := validateResizeParameters(testCase.params)
		if err != testCase.expectedErr {
			t.Errorf("expected error: %v, got: %v", testCase.expectedErr, err)
		}
		if testCase.params.ResizeMode != testCase.expectedMode {
			t.Errorf("expected mode: %v, got: %v", testCase.expectedMode, testCase.params.ResizeMode)
		}
	}
}
