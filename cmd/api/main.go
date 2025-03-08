package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sebsvt/nietzsche/cmd/api/handler"
	"github.com/sebsvt/nietzsche/internal/nietzsche"
	"github.com/sebsvt/nietzsche/pkg/storage"
)

func main() {
	storage, err := storage.NewLocalStorage("storage", "https://nietzsche.sebsvt.com")
	if err != nil {
		panic(err)
	}
	nz := nietzsche.NewNietzsche(storage)
	nz_router := handler.NewNietzscheHandler(nz)
	app := fiber.New()
	// enable logger and cors
	app.Use(cors.New())
	app.Use(logger.New())
	app.Get("/", nz_router.Start)
	app.Post("/upload", nz_router.Upload)
	app.Post("/process", nz_router.Process)
	app.Get("/download", nz_router.Download)
	app.Listen(":8080")
}
