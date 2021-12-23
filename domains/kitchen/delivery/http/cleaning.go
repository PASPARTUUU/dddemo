package http

import (
	"net/http"

	"dddemo/domains/kitchen"
	uc "dddemo/domains/kitchen/kitchenusecase"

	"github.com/labstack/echo/v4"
)

type HandlerCleaning struct {
	useCase kitchen.Cleaning
}

func NewHandlerCleaning(ingredientRepo kitchen.IngredientRepo) *HandlerCleaning {
	return &HandlerCleaning{
		useCase: uc.NewCleaningUseCase(ingredientRepo),
	}
}

func (h HandlerCleaning) CleanIt(ectx echo.Context) error {
	res, _ := h.useCase.CleanIt()
	return ectx.JSON(http.StatusOK, res)
}
