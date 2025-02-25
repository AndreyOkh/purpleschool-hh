package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	templ_adapter "hh/pkg/templ-adapter"
	"hh/views"
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
	h.router.Get("/", h.home)
	h.router.Get("/error", h.error)

}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Main()
	return templ_adapter.Render(c, component)
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
