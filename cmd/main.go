package main

import (
	"fmt"
	"log"
	"time"

	"github.com/sebsvt/nietzsche/internal/pdf"
)

func main() {
	start := time.Now()
	// pdf.StartServer()

	err := pdf.FromOffice(&pdf.ConverterFromOfficeParams{
		InputPath: "test_assets/sample.docx",
		OutputDir: "test_assets",
	})
	if err != nil {
		elapsed := time.Since(start)
		fmt.Printf("Time taken: %s\n", elapsed)
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s\n", elapsed)
}
