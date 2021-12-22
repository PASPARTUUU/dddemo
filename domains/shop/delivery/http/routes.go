package http

import (
	echo "github.com/labstack/echo/v4"
)

func RegisterHTTPEndpoints(router *echo.Group, h *Handler) {

	dishes := router.Group("/dish")
	{
		dishes.POST("/create", h.Dish.CreateDish)
		dishes.GET("/show", h.Dish.ShowDishes)
		dishes.POST("/delete", h.Dish.DeleteDish)
	}

}
