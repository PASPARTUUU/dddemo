package http

import (
	"net/http"

	"dddemo/domains/shop"
	uc "dddemo/domains/shop/usecase"

	"github.com/labstack/echo/v4"
)

type HandlerDish struct {
	useCase shop.DishUseCase
}

func NewHandlerDish(dishRepo shop.DishRepo) *HandlerDish {
	return &HandlerDish{
		useCase: uc.NewDishUseCase(dishRepo),
	}
}

func (h HandlerDish) CreateDish(ectx echo.Context) error {
	return h.useCase.UCreateDish("Borshch", "potato;tomato;capusto", 123)
}

func (h HandlerDish) ShowDishes(ectx echo.Context) error {
	d, _ := h.useCase.UGetDishes()
	return ectx.JSON(http.StatusOK, d)
}

func (h HandlerDish) DeleteDish(ectx echo.Context) error {
	return h.useCase.UDeleteDish("Borshch")
}
