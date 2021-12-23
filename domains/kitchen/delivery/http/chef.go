package http

import (
	"net/http"

	"dddemo/domains/kitchen"
	uc "dddemo/domains/kitchen/kitchenusecase"
	"dddemo/models"

	"github.com/labstack/echo/v4"
)

type HandlerChef struct {
	useCase kitchen.ChefUseCase
}

func NewHandlerChef(ingredientRepo kitchen.IngredientRepo) *HandlerChef {
	return &HandlerChef{
		useCase: uc.NewChefUseCase(ingredientRepo),
	}
}

func (h HandlerChef) Cook(ectx echo.Context) error {

	var bind struct {
		Name   string `json:"name"`
		Recipe string `json:"recipe"`
	}
	ectx.Bind(&bind)

	dish := models.Dish{
		Name:   bind.Name,
		Recipe: bind.Recipe,
	}

	res, _ := h.useCase.Cook(dish)

	return ectx.String(http.StatusOK, res)
}
