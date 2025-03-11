package main

import (
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
	"hh/config"
	"hh/internal/home"
	"hh/internal/vacancy"
	"hh/pkg/database"
	"hh/pkg/logger"
	"runtime/debug"
)

func main() {
	config.Init()

	// Load configs
	dbConf := config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)

	// Creat DB POOL
	dbPool := database.CreateDBPool(dbConf, customLogger)
	defer dbPool.Close()

	// Return info
	info, _ := debug.ReadBuildInfo()
	log.Debug().Msgf("%+v", info.Main.Version)

	app := fiber.New()

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())
	app.Static("/public", "./public")

	// Repositories
	vacancyRepository := vacancy.NewRepository(dbPool, customLogger)

	// Handlers
	home.NewHomeHandler(app, customLogger)
	vacancy.NewHandler(app, customLogger, vacancyRepository)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
