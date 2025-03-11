package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"hh/internal/vacancy"
	templ_adapter "hh/pkg/templ-adapter"
	"hh/views"
	"net/http"
)

type HomeHandler struct {
	router            fiber.Router
	log               zerolog.Logger
	vacancyRepository *vacancy.Repository
}

func NewHomeHandler(router fiber.Router, customLogger *zerolog.Logger, vacancyRepository *vacancy.Repository) {
	h := &HomeHandler{
		router:            router,
		log:               *customLogger,
		vacancyRepository: vacancyRepository,
	}
	h.router.Get("/", h.home)
	h.router.Get("/error", h.error)
	h.router.Get("/favicon.ico", h.favicon)

}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	vacancies, err := h.vacancyRepository.GetAll()
	if err != nil {
		h.log.Error().Err(err).Msg("Error getting vacancies")
		return c.SendStatus(http.StatusInternalServerError)
	}
	component := views.Main(vacancies)
	return templ_adapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) favicon(c *fiber.Ctx) error {
	return c.SendFile("./public/favicon.ico")
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
