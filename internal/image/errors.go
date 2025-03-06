package image

import "fmt"

var (
	ErrCouldNotReadFile        = fmt.Errorf("could not read file")
	ErrCouldNotDecodeFile      = fmt.Errorf("could not decode file")
	ErrCouldNotEncodeFile      = fmt.Errorf("could not encode file")
	ErrCouldNotSaveFile        = fmt.Errorf("could not save file")
	ErrInvalidResizeMode       = fmt.Errorf("invalid resize mode")
	ErrInvalidPercentage       = fmt.Errorf("percentage must be greater than 0")
	ErrInvalidPixelWidth       = fmt.Errorf("pixel width must be greater than 0")
	ErrInvalidPixelHeight      = fmt.Errorf("pixel height must be greater than 0")
	ErrInvalidPixel            = fmt.Errorf("pixel width or height must be valid")
	ErrWidthOrHeightZero       = fmt.Errorf("width or height must be greater than 0")
	ErrCoordinatesInvalid      = fmt.Errorf("coordinates must be greater than 0")
	ErrInvalidCompressionLevel = fmt.Errorf("invalid compression level")
	ErrInvalidFormat           = fmt.Errorf("invalid format")
	ErrInvalidAngle            = fmt.Errorf("invalid angle")
	ErrInvalidMultiplier       = fmt.Errorf("invalid multiplier")
	ErrCouldNotCreateFile      = fmt.Errorf("could not create file")
)
