package main

import (
	"github.com/sebsvt/nietzsche/internal/image"
	"github.com/sebsvt/nietzsche/pkg/logging"
)

func main() {
	logging.Info("Starting...")
	image.AddWatermark(&image.AddWatermarkParameters{
		Elements: []image.WatermarkElement{
			{
				Type: "text",
			},
		},
	})
}
