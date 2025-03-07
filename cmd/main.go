package main

import (
	"log"

	"github.com/sebsvt/nietzsche/internal/pdf"
)

func main() {
	// pdf.StartServer()

	if err := pdf.ConvertFromPDFToImage(pdf.ConvertFromPDFToImageParams{
		InputPath:  "test_assets/aero.pdf",
		OutputPath: "test_assets/aero.zip",
	}); err != nil {
		log.Fatal(err)
	}

}
