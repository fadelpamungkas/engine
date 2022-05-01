package main

import (
	"fmt"
	"log"
	"os"

	"github.com/engine/configs"
	"github.com/engine/delivery/handler"
	"github.com/engine/repository"
	"github.com/engine/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	env, err := configs.Environment()

	mongo, err := configs.Connect(env.Database)
	if err != nil {
		println(err)
	}
	log.Println("Connected to MongoDB")

	repo := repository.NewRepository(mongo)
	uc := usecase.NewUsecase(repo)
	handler.Routes(app, uc)

	port := os.Getenv("PORT")
	if port == "" {
		port = fmt.Sprintf("%v", env.Port)
	}
	log.Println("Server running on port: " + port)
	app.Listen(":" + port)

}
