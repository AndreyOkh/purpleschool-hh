package main

import (
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
	"hh/config"
	"hh/internal/home"
	"hh/pkg/logger"
	"runtime/debug"
)

func main() {
	config.Init()

	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)

	info, _ := debug.ReadBuildInfo()
	log.Debug().Msgf("%+v", info.Main.Version)

	app := fiber.New()

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())
	app.Static("/public", "./public")
	home.NewHomeHandler(app, customLogger)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
