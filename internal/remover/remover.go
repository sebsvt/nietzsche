package remover

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/disintegration/imaging"
	ort "github.com/yalue/onnxruntime_go"
)

const (
	inputWidth  = 1024
	inputHeight = 1024
)

type BackgroundRemover struct {
	modelPath   string
	inputName   string
	outputName  string
	initialized bool
	mutex       sync.Mutex
	modelBytes  []byte
}

func NewBackgroundRemover(modelPath string) (*BackgroundRemover, error) {
	// Set the path to the ONNX Runtime shared library based on OS
	var ortPath string
	switch runtime.GOOS {
	case "windows":
		ortPath = "onnxruntime.dll"
	case "darwin":
		ortPath = "libonnxruntime.dylib"
	default: // Linux and others
		ortPath = "libonnxruntime.so"
	}
	ort.SetSharedLibraryPath(ortPath)

	// Initialize ONNX Runtime environment
	err := ort.InitializeEnvironment()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize ONNX Runtime: %v", err)
	}

	// Load model into memory
	modelBytes, err := os.ReadFile(modelPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read model file: %v", err)
	}

	remover := &BackgroundRemover{
		modelPath:  modelPath,
		modelBytes: modelBytes,
	}

	// Initialize the remover
	err = remover.initialize()
	if err != nil {
		return nil, err
	}

	return remover, nil
}

func (r *BackgroundRemover) initialize() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.initialized {
		return nil
	}

	// Get model input and output info
	inputs, outputs, err := ort.GetInputOutputInfo(r.modelPath)
	if err != nil {
		return fmt.Errorf("failed to get model info: %v", err)
	}

	if len(inputs) == 0 || len(outputs) == 0 {
		return fmt.Errorf("model has no inputs or outputs")
	}

	// Store input and output names
	r.inputName = inputs[0].Name
	r.outputName = outputs[0].Name

	r.initialized = true
	return nil
}

func (r *BackgroundRemover) removeBackground(img image.Image) (image.Image, error) {
	// Resize the image preserving aspect ratio
	resizedImg := resizePreservingAspectRatio(img, inputWidth, inputHeight)

	// Prepare input tensor
	inputData := preprocessImageToTensor(resizedImg)
	inputShape := ort.NewShape(1, 3, inputHeight, inputWidth) // NCHW format
	inputTensor, err := ort.NewTensor(inputShape, inputData)
	if err != nil {
		return nil, fmt.Errorf("failed to create input tensor: %v", err)
	}
	defer inputTensor.Destroy()

	// Prepare output tensor
	outputShape := ort.NewShape(1, 1, inputHeight, inputWidth) // NCHW format, 1 channel for mask
	outputTensor, err := ort.NewEmptyTensor[float32](outputShape)
	if err != nil {
		return nil, fmt.Errorf("failed to create output tensor: %v", err)
	}
	defer outputTensor.Destroy()

	// Create a new session with the tensors
	session, err := ort.NewAdvancedSession(r.modelPath,
		[]string{r.inputName}, []string{r.outputName},
		[]ort.Value{inputTensor}, []ort.Value{outputTensor}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create ONNX session: %v", err)
	}
	defer session.Destroy()

	// Run inference
	err = session.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run inference: %v", err)
	}

	// Get output data
	outputData := outputTensor.GetData()

	// Process the output mask and create the final image
	return applyMaskToImage(img, outputData), nil
}

func (r *BackgroundRemover) Close() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if !r.initialized {
		return
	}

	// Clean up ONNX Runtime
	ort.DestroyEnvironment()
	r.initialized = false
}

// Remove is the main function to remove background from an image
func Remove(inputPath, outputPath, modelPath string) error {
	// Initialize the background remover
	remover, err := NewBackgroundRemover(modelPath)
	if err != nil {
		return fmt.Errorf("failed to create background remover: %v", err)
	}
	defer remover.Close()

	// Load the input image
	img, err := loadAndPreprocessImage(inputPath)
	if err != nil {
		return fmt.Errorf("failed to load image: %v", err)
	}

	// Remove background
	resultImg, err := remover.removeBackground(img)
	if err != nil {
		return fmt.Errorf("failed to remove background: %v", err)
	}

	// Save the result
	return saveImage(outputPath, resultImg)
}

// Utility functions
func resizePreservingAspectRatio(img image.Image, targetWidth, targetHeight int) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	// Calculate aspect ratios
	imgAspect := float64(width) / float64(height)
	targetAspect := float64(targetWidth) / float64(targetHeight)

	var newWidth, newHeight int

	// Determine new dimensions based on aspect ratio
	if imgAspect > targetAspect {
		// Image is wider than target, constrain by width
		newWidth = targetWidth
		newHeight = int(float64(targetWidth) / imgAspect)
	} else {
		// Image is taller than target, constrain by height
		newHeight = targetHeight
		newWidth = int(float64(targetHeight) * imgAspect)
	}

	// Resize the image using the calculated dimensions
	resizedImg := imaging.Resize(img, newWidth, newHeight, imaging.Lanczos)

	// Create a new image with the target dimensions (padded with transparency if needed)
	result := image.NewRGBA(image.Rect(0, 0, targetWidth, targetHeight))

	// Calculate position to center the resized image
	offsetX := (targetWidth - newWidth) / 2
	offsetY := (targetHeight - newHeight) / 2

	// Copy the resized image to the center of the result image
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			result.Set(x+offsetX, y+offsetY, resizedImg.At(x, y))
		}
	}

	return result
}

func loadAndPreprocessImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func preprocessImageToTensor(img image.Image) []float32 {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	// Create a flat array for the tensor data in NCHW format
	// 1 batch, 3 channels (RGB), height, width
	tensorData := make([]float32, 3*height*width)

	// Fill the tensor with normalized pixel values
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			// Convert from 0-65535 to 0-1 float32 and normalize
			rFloat := float32(r) / 65535.0
			gFloat := float32(g) / 65535.0
			bFloat := float32(b) / 65535.0

			// NCHW format: all reds, then all greens, then all blues
			tensorData[0*height*width+y*width+x] = rFloat
			tensorData[1*height*width+y*width+x] = gFloat
			tensorData[2*height*width+y*width+x] = bFloat
		}
	}

	return tensorData
}

func applyMaskToImage(img image.Image, maskData []float32) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	// Create a new RGBA image for the result
	resultImg := image.NewRGBA(bounds)

	// Calculate scaling factors to map mask coordinates to original image
	scaleX := float64(inputWidth) / float64(width)
	scaleY := float64(inputHeight) / float64(height)

	// Apply the mask
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Map coordinates to mask space
			maskX := int(float64(x) * scaleX)
			maskY := int(float64(y) * scaleY)

			// Ensure coordinates are within bounds
			if maskX >= inputWidth {
				maskX = inputWidth - 1
			}
			if maskY >= inputHeight {
				maskY = inputHeight - 1
			}

			// Get the mask value at this position (threshold at 0.5)
			maskIdx := maskY*inputWidth + maskX
			maskValue := maskData[maskIdx]

			// Get the original pixel color
			originalColor := img.At(x, y)
			r, g, b, _ := originalColor.RGBA()

			if maskValue > 0.5 {
				// Foreground: keep the original color
				resultImg.Set(x, y, color.RGBA{
					R: uint8(r >> 8),
					G: uint8(g >> 8),
					B: uint8(b >> 8),
					A: 255, // Fully opaque
				})
			} else {
				// Background: transparent
				resultImg.Set(x, y, color.RGBA{
					R: uint8(r >> 8),
					G: uint8(g >> 8),
					B: uint8(b >> 8),
					A: 0, // Fully transparent
				})
			}
		}
	}

	return resultImg
}

func saveImage(path string, img image.Image) error {
	// Create the directory if it doesn't exist
	dir := filepath.Dir(path)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// Create the output file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode and save the image
	return png.Encode(file, img)
}
