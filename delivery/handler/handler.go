package handler

import (
	"context"

	"github.com/engine/models"
	"github.com/gofiber/fiber/v2"
)

type UsecaseI interface {
	Get(ctx context.Context) (models.Response, error)
}

type Handler struct {
	app *fiber.App
	u   UsecaseI
}

func Routes(app *fiber.App, u UsecaseI) {
	handler := &Handler{
		app: app,
		u:   u,
	}

	app.Get("/", handler.Get)
}

func (h *Handler) Get(c *fiber.Ctx) error {
	data, err := h.u.Get(context.Background())
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(data)
}
