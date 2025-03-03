package vacancy

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	templ_adapter "hh/pkg/templ-adapter"
	"hh/views/components"
)

type VacancyHandler struct {
	router fiber.Router
	log    zerolog.Logger
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &VacancyHandler{
		router: router,
		log:    *customLogger,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)

}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	email := c.FormValue("email")
	h.log.Info().Str("email", email).Msg("VacancyHandler")
	h.log.Info().Msg(string(c.Body()))

	component := components.Notification("Success")
	return templ_adapter.Render(c, component)
}
