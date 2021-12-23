package http

import (
	"net/http"

	"dddemo/domains/kitchen"
	uc "dddemo/domains/kitchen/kitchenusecase"
	"dddemo/models"

	"github.com/labstack/echo/v4"
)

type HandlerIngredient struct {
	useCase kitchen.IngredientUseCase
}

func NewHandlerIngredient(ingredientRepo kitchen.IngredientRepo) *HandlerIngredient {
	return &HandlerIngredient{
		useCase: uc.NewIngredientUseCase(ingredientRepo),
	}
}

func (h HandlerIngredient) Add(ectx echo.Context) error {
	h.useCase.UAddIngredient(models.Ingredient{Name: "potato"})
	h.useCase.UAddIngredient(models.Ingredient{Name: "tomato"})
	h.useCase.UAddIngredient(models.Ingredient{Name: "kapusto"}) // cabbage
	return nil
}

func (h HandlerIngredient) ShowIngredients(ectx echo.Context) error {
	ingr, _ := h.useCase.UShowIngredients()
	return ectx.JSON(http.StatusOK, ingr)
}

func (h HandlerIngredient) TakeIngredients(ectx echo.Context) error {
	ingr, _ := h.useCase.UTakeIngredients()
	return ectx.JSON(http.StatusOK, ingr)
}
