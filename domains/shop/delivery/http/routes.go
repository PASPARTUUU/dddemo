package http

import (
	echo "github.com/labstack/echo/v4"
)

func RegisterHTTPEndpoints(router *echo.Group, h *Handler) {

	// css стили (.Static(...)) не переходят из родительской группы
	// router.Static("/domains/shop/web/static", "./domains/shop/web/static")
	// router.Static("/static", "./static")

	shop := router.Group("/shop")
	shop.Static("/domains/shop/web/static", "./domains/shop/web/static")
	{
		shop.GET("/hello", h.hello)
	}

	dishes := router.Group("/dish")
	{
		dishes.POST("/create", h.Dish.CreateDish)
		dishes.GET("/show", h.Dish.ShowDishes)
		dishes.POST("/delete", h.Dish.DeleteDish)
	}

}
