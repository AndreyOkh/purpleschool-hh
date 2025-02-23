package home

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"text/template"
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
	tmpl := template.Must(template.ParseFiles("./html/page.html"))
	data := struct {
		Count int
	}{
		Count: 1,
	}

	var tmp bytes.Buffer
	if err := tmpl.Execute(&tmp, data); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error(), "Template compile error")
	}
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	return c.Send(tmp.Bytes())
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
