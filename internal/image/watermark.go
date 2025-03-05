package image

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"golang.org/x/image/font"

	"github.com/disintegration/imaging"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
)

// AddWatermarkParameters defines the structure for watermarking
type AddWatermarkParameters struct {
	Elements   []WatermarkElement
	InputPath  string
	OutputPath string
}

// WatermarkElement represents individual watermark components
type WatermarkElement struct {
	Type                           string
	Text                           string
	Image                          string
	Gravity                        string
	VerticalAdjustmentPercentage   int
	HorizontalAdjustmentPercentage int
	Rotation                       int
	FontFamily                     string
	FontStyle                      string
	FontSize                       int
	FontColor                      string
	Transparency                   int
	Mosaic                         bool
}

// AddWatermark processes the image with specified watermarks
func AddWatermark(params *AddWatermarkParameters) error {
	// Validate input
	if err := validateInput(params); err != nil {
		return err
	}

	// Open base image
	baseFile, err := os.Open(params.InputPath)
	if err != nil {
		return err
	}
	defer baseFile.Close()

	// Decode base image
	baseImg, format, err := image.Decode(baseFile)
	if err != nil {
		return err
	}

	// Create mutable RGBA image
	workingImg := image.NewRGBA(baseImg.Bounds())
	draw.Draw(workingImg, workingImg.Bounds(), baseImg, image.Point{}, draw.Src)

	// Process each watermark element
	for _, element := range params.Elements {
		switch strings.ToLower(element.Type) {
		case "text":
			if err := addTextWatermark(workingImg, &element); err != nil {
				return err
			}
		case "image":
			if err := addImageWatermark(workingImg, &element); err != nil {
				return err
			}
		}
	}

	// Generate output path if not provided
	outputPath := params.OutputPath

	// Create output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Encode and save
	switch format {
	case "png":
		err = png.Encode(outputFile, workingImg)
	case "jpeg":
		err = jpeg.Encode(outputFile, workingImg, nil)
	default:
		err = png.Encode(outputFile, workingImg)
	}

	if err != nil {
		return err
	}

	return nil
}

func validateInput(params *AddWatermarkParameters) error {
	if len(params.Elements) == 0 {
		return fmt.Errorf("no watermark elements provided")
	}

	return nil
}

// parseColor converts hex color to RGBA
func parseColor(hexColor string) color.RGBA {
	hexColor = strings.TrimPrefix(hexColor, "#")

	if len(hexColor) != 6 {
		return color.RGBA{0, 0, 0, 255}
	}

	r := hexToDecimal(hexColor[0:2])
	g := hexToDecimal(hexColor[2:4])
	b := hexToDecimal(hexColor[4:6])

	return color.RGBA{r, g, b, 255}
}

// hexToDecimal converts hex to decimal
func hexToDecimal(hex string) uint8 {
	var result uint8
	for _, c := range hex {
		result = result * 16
		switch {
		case c >= '0' && c <= '9':
			result += uint8(c - '0')
		case c >= 'a' && c <= 'f':
			result += uint8(c - 'a' + 10)
		case c >= 'A' && c <= 'F':
			result += uint8(c - 'A' + 10)
		}
	}
	return result
}

// calculatePosition determines watermark placement
func calculatePosition(imgWidth, imgHeight, elementWidth, elementHeight int, gravity string, vertAdj, horizAdj int) (x, y int) {
	switch strings.ToLower(gravity) {
	case "northwest":
		x, y = 0, 0
	case "north":
		x = (imgWidth - elementWidth) / 2
		y = 0
	case "northeast":
		x = imgWidth - elementWidth
		y = 0
	case "west":
		x = 0
		y = (imgHeight - elementHeight) / 2
	case "center":
		x = (imgWidth - elementWidth) / 2
		y = (imgHeight - elementHeight) / 2
	case "east":
		x = imgWidth - elementWidth
		y = (imgHeight - elementHeight) / 2
	case "southwest":
		x = 0
		y = imgHeight - elementHeight
	case "south":
		x = (imgWidth - elementWidth) / 2
		y = imgHeight - elementHeight
	case "southeast":
		x = imgWidth - elementWidth
		y = imgHeight - elementHeight
	default:
		x = (imgWidth - elementWidth) / 2
		y = (imgHeight - elementHeight) / 2
	}

	x += int(float64(imgWidth) * (float64(horizAdj) / 100.0))
	y += int(float64(imgHeight) * (float64(vertAdj) / 100.0))

	return x, y
}

// reduceNoiseAndCompress applies noise reduction techniques
func reduceNoiseAndCompress(img image.Image) *image.RGBA {
	// Apply gentle noise reduction
	processedImg := imaging.Blur(img, 0.3)

	// Adjust contrast slightly
	processedImg = imaging.AdjustContrast(processedImg, -5)

	// Convert to RGBA
	// return imaging.Clone(processedImg)
	rgba := image.NewRGBA(processedImg.Bounds())
	draw.Draw(rgba, rgba.Bounds(), processedImg, processedImg.Bounds().Min, draw.Src)
	return rgba
}

// addImageWatermark adds an image watermark
func addImageWatermark(baseImg *image.RGBA, element *WatermarkElement) error {
	// Open overlay image
	overlayFile, err := os.Open(element.Image)
	if err != nil {
		return err
	}
	defer overlayFile.Close()

	// Decode and reduce noise
	overlay, _, err := image.Decode(overlayFile)
	if err != nil {
		return err
	}

	// Noise reduction
	cleanOverlay := reduceNoiseAndCompress(overlay)

	// Base image dimensions
	baseWidth := baseImg.Bounds().Dx()
	baseHeight := baseImg.Bounds().Dy()

	// Watermark sizing (20% of base image)
	watermarkWidth := int(float64(baseWidth) * 0.2)
	watermarkHeight := int(float64(baseHeight) * 0.2)

	// Resize watermark
	resizedWatermark := resize.Resize(
		uint(watermarkWidth),
		uint(watermarkHeight),
		cleanOverlay,
		resize.Lanczos3,
	)

	// Rotation
	if element.Rotation != 0 {
		resizedWatermark = imaging.Rotate(resizedWatermark, float64(element.Rotation), color.Transparent)
	}

	// Position calculation
	x, y := calculatePosition(
		baseWidth,
		baseHeight,
		resizedWatermark.Bounds().Dx(),
		resizedWatermark.Bounds().Dy(),
		element.Gravity,
		element.VerticalAdjustmentPercentage,
		element.HorizontalAdjustmentPercentage,
	)

	// Transparency handling
	transparentWatermark := image.NewRGBA(resizedWatermark.Bounds())
	for dy := resizedWatermark.Bounds().Min.Y; dy < resizedWatermark.Bounds().Max.Y; dy++ {
		for dx := resizedWatermark.Bounds().Min.X; dx < resizedWatermark.Bounds().Max.X; dx++ {
			originalColor := resizedWatermark.At(dx, dy)
			r, g, b, a := originalColor.RGBA()

			newAlpha := uint8(float64(a>>8) * (float64(100-element.Transparency) / 100.0))

			transparentWatermark.Set(dx, dy, color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: newAlpha,
			})
		}
	}

	// Blend watermark
	draw.DrawMask(
		baseImg,
		image.Rect(x, y, x+transparentWatermark.Bounds().Dx(), y+transparentWatermark.Bounds().Dy()),
		transparentWatermark,
		image.Point{},
		transparentWatermark,
		image.Point{},
		draw.Over,
	)

	return nil
}

func addTextWatermark(img *image.RGBA, element *WatermarkElement) error {
	// Load font
	fontPath := "aria.ttf"
	fontBytes, err := os.ReadFile(fontPath)
	if err != nil {
		return err
	}

	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return err
	}

	// Parse color
	textColor := parseColor(element.FontColor)

	// Create font context
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(float64(element.FontSize))
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.NewUniform(textColor))

	// Calculate text size
	drawer := &font.Drawer{
		Face: truetype.NewFace(f, &truetype.Options{Size: float64(element.FontSize)}),
	}
	textWidth := drawer.MeasureString(element.Text)
	textHeight := element.FontSize

	// Calculate position
	x, y := calculatePosition(
		img.Bounds().Dx(),
		img.Bounds().Dy(),
		// fixed.Int26_6ToFloat64(textWidth),
		int(float64(textWidth)/64),

		textHeight,
		element.Gravity,
		element.VerticalAdjustmentPercentage,
		element.HorizontalAdjustmentPercentage,
	)

	// Draw text
	pt := freetype.Pt(x, y+element.FontSize)
	_, err = c.DrawString(element.Text, pt)

	return err
}
