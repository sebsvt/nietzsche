package main

import (
	"log"

	"github.com/sebsvt/nietzsche/internal/pdf"
)

func main() {
	if err := pdf.FromImage(&pdf.ConverterFromImageParams{
		InputPath:   "test_assets/meow.jpg",
		OutputPath:  "test_assets/meow.pdf",
		Orientation: "Portrait",
		Margin:      0,
		PageSize:    "fit",
	}); err != nil {
		log.Fatal(err)
	}
}
