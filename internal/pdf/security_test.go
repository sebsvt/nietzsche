package pdf

import (
	"os"
	"testing"

	"github.com/sebsvt/nietzsche/pkg/file"
)

func TestProtect(t *testing.T) {
	type testCase struct {
		name           string
		params         *SecurityParams
		expectedErr    error
		expectedOutput bool
	}
	testCases := []testCase{
		{
			name: "protect file",
			params: &SecurityParams{
				InputFilePath:  "../../test_assets/aero.pdf",
				OutputFilePath: "../../test_assets/aero_protected.pdf",
				Password:       "password",
			},
			expectedErr:    nil,
			expectedOutput: true,
		},
		{
			name: "protect file with empty input file path should return error",
			params: &SecurityParams{
				InputFilePath:  "",
				OutputFilePath: "../../test_assets/aero_protected.pdf",
				Password:       "password",
			},
			expectedErr:    ErrInputFilePathEmpty,
			expectedOutput: false,
		},
		{
			name: "protect file with empty output file path should return error",
			params: &SecurityParams{
				InputFilePath:  "../../test_assets/aero.pdf",
				OutputFilePath: "",
				Password:       "password",
			},
			expectedErr:    ErrOutputFilePathEmpty,
			expectedOutput: false,
		},
		{
			name: "protect file with empty password should return error",
			params: &SecurityParams{
				InputFilePath:  "../../test_assets/aero.pdf",
				OutputFilePath: "../../test_assets/aero_protected.pdf",
				Password:       "",
			},
			expectedErr:    ErrPasswordRequired,
			expectedOutput: false,
		},
		{
			name: "input file does not exist should return error",
			params: &SecurityParams{
				InputFilePath:  "../../test_assets/does_not_exist.pdf",
				OutputFilePath: "../../test_assets/aero_protected.pdf",
				Password:       "password",
			},
			expectedErr:    ErrFailedToProtectFile,
			expectedOutput: false,
		},
		{
			name: "input file is not a pdf should return error",
			params: &SecurityParams{
				InputFilePath:  "../../test_assets/not_a_pdf.txt",
				OutputFilePath: "../../test_assets/not_a_pdf_protected.pdf",
				Password:       "password",
			},
			expectedErr:    ErrFailedToProtectFile,
			expectedOutput: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := Protect(tc.params)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
			defer os.Remove(tc.params.OutputFilePath)
			if tc.expectedOutput {
				if !file.FileExists(tc.params.OutputFilePath) {
					t.Errorf("expected output file to exist, got %v", tc.params.OutputFilePath)
				}
			}
		})
	}
}

func TestUnlock(t *testing.T) {
	type testCase struct {
		name           string
		params         *SecurityParams
		expectedErr    error
		expectedOutput bool
	}
	testCases := []testCase{
		{
			name: "unlock file",
			params: &SecurityParams{
				InputFilePath:  "../../test_assets/aero_protected_sample.pdf",
				OutputFilePath: "../../test_assets/aero_unlocked.pdf",
				Password:       "password",
			},
			expectedErr:    nil,
			expectedOutput: true,
		},
		{
			name: "unlock file with empty input file path should return error",
			params: &SecurityParams{
				InputFilePath:  "",
				OutputFilePath: "../../test_assets/aero_unlocked.pdf",
				Password:       "password",
			},
			expectedErr:    ErrInputFilePathEmpty,
			expectedOutput: false,
		},
		{
			name: "unlock file with empty output file path should return error",
			params: &SecurityParams{
				InputFilePath:  "../../test_assets/aero_protected_sample.pdf",
				OutputFilePath: "",
				Password:       "password",
			},
			expectedErr:    ErrOutputFilePathEmpty,
			expectedOutput: false,
		},
		{
			name: "unlock file with empty password should return error",
			params: &SecurityParams{
				InputFilePath:  "../../test_assets/aero_protected_sample.pdf",
				OutputFilePath: "../../test_assets/aero_unlocked.pdf",
				Password:       "",
			},
			expectedErr:    ErrPasswordRequired,
			expectedOutput: false,
		},
		{
			name: "input file is not a pdf should return error",
			params: &SecurityParams{
				InputFilePath:  "../../test_assets/not_image.txt",
				OutputFilePath: "../../test_assets/not_a_pdf_unlocked.pdf",
				Password:       "password",
			},
			expectedErr:    ErrFailedToUnlockFile,
			expectedOutput: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := Unlock(tc.params)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
			defer os.Remove(tc.params.OutputFilePath)
			if tc.expectedOutput {
				if !file.FileExists(tc.params.OutputFilePath) {
					t.Errorf("expected output file to exist, got %v", tc.params.OutputFilePath)
				}
			}
		})
	}
}
