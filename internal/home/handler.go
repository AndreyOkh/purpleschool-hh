package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"hh/internal/vacancy"
	templ_adapter "hh/pkg/templ-adapter"
	"hh/views"
	"math"
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
	pageItems := 2
	amountVacancies := h.vacancyRepository.CountVacancies()
	amountPages := int(math.Ceil(float64(amountVacancies / pageItems)))
	page := c.QueryInt("page", 1)
	if page > amountPages || page < 1 {
		return c.SendStatus(http.StatusNotFound)
	}
	vacancies, err := h.vacancyRepository.GetAll(pageItems, (page-1)*pageItems)
	if err != nil {
		h.log.Error().Err(err).Msg("Error getting vacancies")
		return c.SendStatus(http.StatusInternalServerError)
	}
	component := views.Main(vacancies, amountPages, page)
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
