package http

import "dddemo/domains/shop"

type Handler struct {
	Dish *HandlerDish
}

func NewHandler(dishRepo shop.DishRepo) *Handler {

	return &Handler{
		Dish: NewHandlerDish(dishRepo),
	}
}
