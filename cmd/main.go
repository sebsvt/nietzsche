package main

import (
	"log"

	"github.com/sebsvt/nietzsche/internal/pdf"
)

func main() {
	if err := pdf.SplitRemove(&pdf.SplitRemoveParams{
		InputPath:   "test_assets/long_page.pdf",
		OutputPath:  "test_assets/long_removeds.pdf",
		RemovePages: "1-3,6",
		MergeAfter:  false,
	}); err != nil {
		log.Fatal(err)
	}
}
