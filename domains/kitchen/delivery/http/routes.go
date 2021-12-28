package http

import (
	echo "github.com/labstack/echo/v4"
)

func RegisterHTTPEndpoints(router *echo.Group, h *Handler) {

	kitchen := router.Group("/kitchen")
	kitchen.Static("/domains/kitchen/web/static", "./domains/kitchen/web/static")
	{
		kitchen.GET("/hello", h.hello)
	}

	ingredients := router.Group("/ingredient")
	{
		ingredients.POST("/add", h.Ingredient.Add)
		ingredients.GET("/show", h.Ingredient.ShowIngredients)
		ingredients.POST("/take", h.Ingredient.TakeIngredients)
	}

	chef := router.Group("/chef")
	{
		chef.POST("/cook", h.Chef.Cook)
	}
}
