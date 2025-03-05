package main

import (
	"fmt"
	"log"

	"github.com/sebsvt/nietzsche/internal/remover"
)

// func main() {
// 	fmt.Println("Starting...")
// }

func main() {
	fmt.Println("Starting...")
	if err := remover.Remove("test_assets/meow.jpg", "test_assets/meow_no_bg.jpg", "model.onnx"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")
}
