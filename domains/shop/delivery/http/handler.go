package http

import (
	"dddemo/domains/shop"
	"dddemo/pkg/meintemplate"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Dish      *HandlerDish
	templates meintemplate.Templates
}

func NewHandler(templates meintemplate.Templates, dishRepo shop.DishRepo) *Handler {
	return &Handler{
		Dish:      NewHandlerDish(templates, dishRepo),
		templates: templates,
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

	return h.templates.Render(ectx, 200, "shop", data)
}
