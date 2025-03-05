package image

import (
	"testing"
)

func TestCompress(t *testing.T) {
	type testCase struct {
		name             string
		inPath           string
		outPath          string
		compressionLevel CompressionLevel
		expectedError    error
	}
	testcases := []testCase{
		{
			name:             "compress cat.jpg with low compression level",
			inPath:           "../../test_assets/cat.jpg",
			outPath:          "../../test_assets/cat_compressed.jpg",
			compressionLevel: CompressionLevelLow,
			expectedError:    nil,
		},
		{
			name:             "compress cat.jpg with recommended compression level",
			inPath:           "../../test_assets/cat.jpg",
			outPath:          "../../test_assets/cat_compressed.jpg",
			compressionLevel: CompressionLevelRecommended,
			expectedError:    nil,
		},
		{
			name:             "compress cat.jpg with extreme compression level",
			inPath:           "../../test_assets/cat.jpg",
			outPath:          "../../test_assets/cat_compressed.jpg",
			compressionLevel: CompressionLevelExtreme,
			expectedError:    nil,
		},
		{
			name:             "compress cat.jpg with invalid compression level",
			inPath:           "../../test_assets/cat.jpg",
			outPath:          "../../test_assets/cat_compressed.jpg",
			compressionLevel: CompressionLevel("invalid"),
			expectedError:    ErrInvalidCompressionLevel,
		},
		{
			name:             "compress with invalid path",
			inPath:           "../../test_assets/invalid.jpg",
			outPath:          "../../test_assets/cat_compressed.jpg",
			compressionLevel: CompressionLevelLow,
			expectedError:    ErrCouldNotReadFile,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			_, err := Compress(&CompressParameters{
				InPath:           testcase.inPath,
				OutPath:          testcase.outPath,
				CompressionLevel: testcase.compressionLevel,
			})
			if err != testcase.expectedError {
				t.Errorf("expected error %v, got %v", testcase.expectedError, err)
			}
		})

	}
}
