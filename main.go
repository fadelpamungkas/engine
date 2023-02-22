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
	"go.uber.org/zap"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	// Initialize Zap logger
	logger, errLog := configs.InitializeZap()
	defer logger.Sync()
	if errLog != nil {
		log.Panicf("can't initialize zap logger: %v", errLog)
	}
	logger.Info("logger initialized")

	// Import Config
	env, errConfig := configs.Environment()
	if errConfig != nil {
		logger.Panic("Failed import configs",
			zap.Error(errConfig),
		)
	}
	logger.Info("Configs imported")

	// mongo, err := configs.Connect(env.Database)
	// if err != nil {
	// 	println(err)
	// }
	// log.Println("Connected to MongoDB")

	// Connect to PostgreSQL
	pgx, err := configs.ConnectPg(env.Postgresql)
	defer pgx.Close()
	if err != nil {
		logger.Panic("Failed to connect PostgreSQL",
			zap.Error(errConfig),
		)
	}
	logger.Info("Connected to PostgreSQL")

	repo := repository.NewRepository(pgx, logger)
	uc := usecase.NewUsecase(repo, logger)
	handler.Routes(app, uc, logger)

	port := os.Getenv("PORT")
	if port == "" {
		port = fmt.Sprintf("%v", env.Port)
	}
	logger.Info("Server running",
		zap.String("Port", port),
	)
	app.Listen(":" + port)

}
