package templ_adapter

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Render(ctx *fiber.Ctx, component templ.Component, code int) error {
	return adaptor.HTTPHandler(templ.Handler(component, templ.WithStatus(code)))(ctx)
}
