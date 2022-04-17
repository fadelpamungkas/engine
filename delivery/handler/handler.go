package handler

import (
	"context"

	"github.com/engine/usecase"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	app *fiber.App
	u   usecase.UsecaseI
}

func NewHandler(app *fiber.App, u usecase.UsecaseI) *Handler {
	return &Handler{
		app: app,
		u:   u,
	}
}

func (h *Handler) Get(c *fiber.Ctx) error {
	data, err := h.u.Get(context.Background())
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(data)
}
