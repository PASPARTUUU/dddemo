package http

import (
	echo "github.com/labstack/echo/v4"
)

func RegisterHTTPEndpoints(router *echo.Group, h *Handler) {

	ingredients := router.Group("/cleaning")
	{
		ingredients.GET("/clean", h.Cleaning.CleanShopAndKitchen)
	}

}
