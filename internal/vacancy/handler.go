package vacancy

import (
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	templ_adapter "hh/pkg/templ-adapter"
	"hh/pkg/validator"
	"hh/views/components"
	"time"
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
	form := VacancyCreateForm{
		Role:     c.FormValue("role"),
		Company:  c.FormValue("company"),
		Type:     c.FormValue("type"),
		Salary:   c.FormValue("salary"),
		Location: c.FormValue("location"),
		Email:    c.FormValue("email"),
	}
	errors := validate.Validate(&validators.EmailIsPresent{
		Name:    "Email",
		Field:   form.Email,
		Message: "e-mail пуст или указан не верно",
	}, &validators.StringIsPresent{
		Name:    "Role",
		Field:   form.Role,
		Message: "Не указана должнось",
	}, &validators.StringIsPresent{
		Name:    "Company",
		Field:   form.Company,
		Message: "Не указано название компании",
	}, &validators.StringIsPresent{
		Name:    "Type",
		Field:   form.Type,
		Message: "Не указана сфера деятельности компании",
	}, &validators.StringIsPresent{
		Name:    "Salary",
		Field:   form.Salary,
		Message: "Не указана заработная плата",
	}, &validators.StringIsPresent{
		Name:    "Location",
		Field:   form.Location,
		Message: "Не указано местоположение компании",
	})
	message := validator.FormatErrors(errors)
	var status components.NotificationStatus
	if len(errors.Errors) > 0 {
		status = components.NotificationFailure
	} else {
		status = components.NotificationSuccess
		message = "Форма отправлена успешно"
	}
	component := components.Notification(message, status)
	time.Sleep(time.Second * 2)
	return templ_adapter.Render(c, component)
}
