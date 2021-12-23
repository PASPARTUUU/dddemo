package http

import "dddemo/domains/kitchen"

// uc "dddemo/domains/kitchen/usecase"

type Handler struct {
	// UseCases    uc.UseCase
	Ingredient *HandlerIngredient
	Chef       *HandlerChef
}

func NewHandler(ingrRepo kitchen.IngredientRepo) *Handler {

	return &Handler{
		Ingredient: NewHandlerIngredient(ingrRepo),
		Chef:       NewHandlerChef(ingrRepo),
	}
}
