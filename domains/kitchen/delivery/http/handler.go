package http

import (
	"dddemo/domains/kitchen"
	"dddemo/pkg/meintemplate"

	"github.com/labstack/echo/v4"
)

// uc "dddemo/domains/kitchen/usecase"

type Handler struct {
	// UseCases    uc.UseCase
	Ingredient *HandlerIngredient
	Chef       *HandlerChef
	templates  meintemplate.Templates
}

func NewHandler(templates meintemplate.Templates, ingrRepo kitchen.IngredientRepo) *Handler {

	return &Handler{
		Ingredient: NewHandlerIngredient(ingrRepo),
		Chef:       NewHandlerChef(ingrRepo),
		templates:  templates,
	}
}

func (h *Handler) hello(ectx echo.Context) error {
	data := struct {
		Title   string
		Message string
	}{
		Title:   "Последние заметки",
		Message: "Здесь пока ничего нет!",
	}

	return h.templates.Render(ectx, 200, "kitchen", data)
}
