package http

import (
	"net/http"

	"dddemo/domains/kitchen"
	"dddemo/domains/shop"
	"dddemo/service"
	"dddemo/service/usecase/cleaning"

	"github.com/labstack/echo/v4"
)

type HandlerCleaning struct {
	useCase service.Cleaning
}

func NewHandlerCleaning(ingredientRepo kitchen.IngredientRepo, dishRepo shop.DishRepo) *HandlerCleaning {
	return &HandlerCleaning{
		useCase: cleaning.NewServiceCleaning(ingredientRepo, dishRepo),
	}
}

func (h HandlerCleaning) CleanShopAndKitchen(ectx echo.Context) error {
	res, _ := h.useCase.CleanShopAndKitchen()
	return ectx.JSON(http.StatusOK, res)
}
