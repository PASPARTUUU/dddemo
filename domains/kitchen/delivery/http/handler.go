package http

import (
	"dddemo/domains/kitchen/repository/localstorage"
	// uc "dddemo/domains/kitchen/usecase"
)

type Handler struct {
	// UseCases    uc.UseCase
	Ingredient *HandlerIngredient
	Chef       *HandlerChef
}

func NewHandler() *Handler {

	ingrRepo := localstorage.NewIngredientLocalStorage()

	return &Handler{
		Ingredient: NewHandlerIngredient(ingrRepo),
		Chef:       NewHandlerChef(ingrRepo),
	}
}
