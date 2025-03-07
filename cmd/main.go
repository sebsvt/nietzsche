package main

import (
	"log"

	"github.com/sebsvt/nietzsche/internal/pdf"
)

func main() {
	// pdf.StartServer()

	// if err := pdf.FromImage(&pdf.ConverterFromImageParams{
	// 	Orientation: "Portrait",
	// 	InputPath:   "test_assets/page_002.pdf",
	// 	OutputPath:  "test_assets/page_002_ocr_square.pdf",
	// 	PageSize:    "A4",
	// }); err != nil {
	// 	log.Fatal(err)
	// }
	if err := pdf.OCRScan(&pdf.OCRScanParams{
		InputPath:  "test_assets/page_002.pdf",
		OutputPath: "test_assets/page_002_ocr_square.pdf",
	}); err != nil {
		log.Fatal(err)
	}

}
