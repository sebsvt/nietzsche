package pdf

import "testing"

func TestCompress(t *testing.T) {
	type testCase struct {
		name             string
		inputPath        string
		outputPath       string
		compressionLevel CompressionLevel
		expectedError    error
		expectedOutput   bool
	}

	testCases := []testCase{
		{
			name:             "input path is empty",
			inputPath:        "",
			outputPath:       "",
			compressionLevel: CompressionLevelLow,
			expectedError:    ErrInputFilePathEmpty,
		},
		{
			name:             "output path is empty",
			inputPath:        "../../test_assets/aero.pdf",
			outputPath:       "",
			compressionLevel: CompressionLevelLow,
			expectedError:    ErrOutputFilePathEmpty,
		},
		{
			name:             "input file does not exist",
			inputPath:        "../../test_assets/does_not_exist.pdf",
			outputPath:       "../../test_assets/aero_low.pdf",
			compressionLevel: CompressionLevelLow,
			expectedError:    ErrFileNotFound,
			expectedOutput:   false,
		},
		{
			name:             "low compression level",
			inputPath:        "../../test_assets/aero.pdf",
			outputPath:       "../../test_assets/aero_low.pdf",
			compressionLevel: CompressionLevelLow,
			expectedError:    nil,
			expectedOutput:   true,
		},
		{
			name:             "recommended compression level",
			inputPath:        "../../test_assets/aero.pdf",
			outputPath:       "../../test_assets/aero_recommended.pdf",
			compressionLevel: CompressionLevelRecommended,
			expectedError:    nil,
			expectedOutput:   true,
		},
		{
			name:             "extreme compression level",
			inputPath:        "../../test_assets/aero.pdf",
			outputPath:       "../../test_assets/aero_extreme.pdf",
			compressionLevel: CompressionLevelExtreme,
			expectedError:    nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := Compress(&CompressParams{
				InputPath:        testCase.inputPath,
				OutputPath:       testCase.outputPath,
				CompressionLevel: testCase.compressionLevel,
			})
			if err != testCase.expectedError {
				t.Errorf("expected error %v, got %v", testCase.expectedError, err)
			}
		})
	}

}
