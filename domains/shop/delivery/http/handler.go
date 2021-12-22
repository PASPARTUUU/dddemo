package http

import (
	"dddemo/domains/shop/repository/localstorage"
)

type Handler struct {
	Dish *HandlerDish
}

func NewHandler() *Handler {

	dishRepo := localstorage.NewDishLocalStorage()

	return &Handler{
		Dish: NewHandlerDish(dishRepo),
	}
}
