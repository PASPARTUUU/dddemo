package http

import (
	"dddemo/domains/shop"
	uc "dddemo/domains/shop/shopusecase"
	"dddemo/pkg/meintemplate"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HandlerDish struct {
	useCase   shop.DishUseCase
	templates meintemplate.Templates
}

func NewHandlerDish(templates meintemplate.Templates, dishRepo shop.DishRepo) *HandlerDish {
	return &HandlerDish{
		useCase:   uc.NewDishUseCase(dishRepo),
		templates: templates,
	}
}

func (h HandlerDish) TmpCreateDish(ectx echo.Context) error {
	return h.templates.Render(ectx, 200, "dish_create", nil)
}

func (h HandlerDish) CreateDish(ectx echo.Context) error {
	dishNumber := ectx.Param("number")

	switch dishNumber {
	case "1":
		h.useCase.UCreateDish("Borshch", "potato;tomato;kapusto", 123)
	case "2":
		h.useCase.UCreateDish("Пельмень", "ком бизнес-логики обернутый тонким интерфейсом", 00)
	}

	return ectx.Redirect(http.StatusFound, "/tavern/dish/show")
}

func (h HandlerDish) ShowDishes(ectx echo.Context) error {
	d, _ := h.useCase.UGetDishes()

	return h.templates.Render(ectx, 200, "dish_show", map[string]interface{}{
		"Dishes": d,
	})
}

func (h HandlerDish) DeleteDish(ectx echo.Context) error {
	return h.useCase.UDeleteDish("Borshch")
}
