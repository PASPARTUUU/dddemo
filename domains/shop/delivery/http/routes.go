package http

import (
	echo "github.com/labstack/echo/v4"
)

func RegisterHTTPEndpoints(router *echo.Group, h *Handler) {

	shop := router.Group("/shop")
	{
		shop.GET("/hello", h.hello)
	}

	dishes := router.Group("/dish")
	{
		dishes.POST("/create/:number", h.Dish.CreateDish)
		dishes.GET("/show", h.Dish.ShowDishes)
		dishes.POST("/delete", h.Dish.DeleteDish)

		dishes.GET("/tmp_create", h.Dish.TmpCreateDish)
	}

}
