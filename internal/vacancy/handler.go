package vacancy

import (
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	templ_adapter "hh/pkg/templ-adapter"
	"hh/pkg/validator"
	"hh/views/components"
	"net/http"
)

type VacancyHandler struct {
	router     fiber.Router
	log        zerolog.Logger
	repository *Repository
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, repository *Repository) {
	h := &VacancyHandler{
		router:     router,
		log:        *customLogger,
		repository: repository,
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

	if len(errors.Errors) > 0 {
		message := validator.FormatErrors(errors)
		component := components.Notification(message, components.NotificationFailure)
		return templ_adapter.Render(c, component, http.StatusBadRequest)
	}

	if err := h.repository.AddVacancy(form); err != nil {
		component := components.Notification("Произошла внутренняя ошибка", components.NotificationFailure)
		return templ_adapter.Render(c, component, http.StatusInternalServerError)
	}

	message := "Форма отправлена успешно"
	component := components.Notification(message, components.NotificationSuccess)
	return templ_adapter.Render(c, component, http.StatusOK)
}
