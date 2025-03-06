package main

import (
	"fmt"
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	// Find a suitable font file
	fontPath := "aria.ttf"

	// Add font
	err := pdf.AddTTFFont("main", fontPath)
	if err != nil {
		log.Fatalf("Failed to add font: %v", err)
	}

	// Set font
	err = pdf.SetFont("main", "", 14)
	if err != nil {
		log.Fatalf("Failed to set font: %v", err)
	}

	for i := 0; i < 5; i++ {
		pdf.AddPage()
		pdf.Br(20)

		// Create placeholder with font set
		err := pdf.PlaceHolderText("totalnumber", 30)
		if err != nil {
			log.Printf("Error creating placeholder: %v", err)
			return
		}
	}

	// Fill in placeholder
	err = pdf.FillInPlaceHoldText("totalnumber", fmt.Sprintf("%d", 5), gopdf.Left)
	if err != nil {
		log.Printf("Error filling placeholder: %v", err)
		return
	}

	// Write PDF
	err = pdf.WritePdf("placeholder_text.pdf")
	if err != nil {
		log.Printf("Error writing PDF: %v", err)
		return
	}

	fmt.Println("PDF created successfully")
}
