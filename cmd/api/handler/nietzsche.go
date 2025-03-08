package handler

import (
	"io"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sebsvt/nietzsche/internal/nietzsche"
)

type nietzscheHandler struct {
	nz nietzsche.Nietzsche
}

func NewNietzscheHandler(nz nietzsche.Nietzsche) *nietzscheHandler {
	return &nietzscheHandler{
		nz: nz,
	}
}

func (n *nietzscheHandler) Start(c *fiber.Ctx) error {
	uuid := uuid.New().String()
	task := "task_" + uuid
	return c.JSON(fiber.Map{
		"server":           "api.nietzsche.com",
		"task":             task,
		"remaning_credits": rand.Intn(100000),
	})
}

// get task and file from form request
func (n *nietzscheHandler) Upload(c *fiber.Ctx) error {
	task := c.FormValue("task")
	_ = task
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to get file",
		})
	}
	content, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to get file content",
		})
	}
	contentBytes, err := io.ReadAll(content)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to read file content",
		})
	}
	uploadResponse, err := n.nz.Upload(file.Filename, contentBytes)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to upload file",
		})
	}
	return c.JSON(fiber.Map{
		"server_filename": uploadResponse.ServerFileName,
	})
}

func (n *nietzscheHandler) Process(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"server": "api.nietzsche.com",
		"task":   "task_123",
		"status": "success",
	})
}

func (n *nietzscheHandler) Download(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"server": "api.nietzsche.com",
		"task":   "task_123",
		"status": "success",
	})
}
