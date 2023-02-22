package handler

import (
	"context"

	"github.com/engine/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UsecaseI interface {
	Get(ctx context.Context) (models.Response, error)
	GetForEachRowGoroutines(ctx context.Context) (models.Response, error)
	GetForEachRow(ctx context.Context) (models.Response, error)
	GetRow(ctx context.Context) (models.Response, error)
	Insert(ctx context.Context, req models.AccountRequest) (models.Response, error)
	Update(ctx context.Context) (models.Response, error)
	Delete(ctx context.Context) (models.Response, error)
}

type handler struct {
	u   UsecaseI
	log *zap.Logger
}

func Routes(app *fiber.App, u UsecaseI, log *zap.Logger) {
	h := &handler{
		u:   u,
		log: log,
	}

	app.Get("/", h.GetData)
	app.Get("/ferg", h.GetForEachRowGoroutines)
	app.Get("/fer", h.GetForEachRow)
	app.Get("/row", h.GetRow)
	app.Post("/insert", h.InsertData)
	app.Get("/update", h.UpdateData)
	app.Get("/delete", h.DeleteData)
}

func (h *handler) GetData(c *fiber.Ctx) error {
	data, err := h.u.Get(context.Background())
	if err != nil {
		h.log.Error("Error fetch data",
			zap.Error(err),
		)
		return err
	}
	return c.Status(fiber.StatusOK).JSON(data)
}

func (h *handler) GetForEachRowGoroutines(c *fiber.Ctx) error {
	data, err := h.u.GetForEachRowGoroutines(context.Background())
	if err != nil {
		h.log.Error("Error fetch data",
			zap.Error(err),
		)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func (h *handler) GetForEachRow(c *fiber.Ctx) error {
	data, err := h.u.GetForEachRow(context.Background())
	if err != nil {
		h.log.Error("Error fetch data",
			zap.Error(err),
		)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func (h *handler) GetRow(c *fiber.Ctx) error {
	data, err := h.u.GetRow(context.Background())
	if err != nil {
		h.log.Error("Error fetch data",
			zap.Error(err),
		)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func (h *handler) InsertData(c *fiber.Ctx) error {
	var validate = validator.New()
	var req = models.AccountRequest{}

	//validate the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data:    err.Error(),
		})
	}
	//use validator library to validate required fields
	if validationErr := validate.Struct(&req); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  fiber.StatusBadRequest,
			Message: "Validation Error",
			Data:    validationErr.Error(),
		})
	}

	data, err := h.u.Insert(context.Background(), req)
	if err != nil {
		h.log.Error("Error input data",
			zap.Error(err),
		)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func (h *handler) UpdateData(c *fiber.Ctx) error {
	data, err := h.u.Update(context.Background())
	if err != nil {
		h.log.Error("Error update data",
			zap.Error(err),
		)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func (h *handler) DeleteData(c *fiber.Ctx) error {
	data, err := h.u.Delete(context.Background())
	if err != nil {
		h.log.Error("Error delete data",
			zap.Error(err),
		)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(data)
}
