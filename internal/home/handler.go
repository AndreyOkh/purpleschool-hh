package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router fiber.Router
	log    zerolog.Logger
}

func NewHomeHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router: router,
		log:    *customLogger,
	}
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)

}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	//log.Trace("Trace")
	//log.Debug("Debug")
	h.log.Info().Bool("isAdmin", true).Msg("Info")
	//log.Warn("Warn")
	//log.Error("Error")
	//log.Panic("Panic")
	return c.SendString("Error")
}
