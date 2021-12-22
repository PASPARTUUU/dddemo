package server

import (
	"github.com/labstack/echo/v4"

	kitchenhttp "dddemo/domains/kitchen/delivery/http"
	// kitchenhttp "dddemo/domains/kitchen/delivery/http"
	// kitchenstorage "dddemo/domains/kitchen/repository/localstorage"
	// kitchenusecase "dddemo/domains/kitchen/usecase"
	// ---
	shophttp "dddemo/domains/shop/delivery/http"
	// shopstorage "dddemo/domains/shop/repository/localstorage"
	// shopusecase "dddemo/domains/shop/usecase"
)

type Server struct {

	// ShopUC    *shopusecase.DishUseCase
	// KitchenUC *kitchenusecase.UseCase

	KitchenHandler *kitchenhttp.Handler
	ShopHandler    *shophttp.Handler
}

func NewServer() *Server {

	// shopRepo := shopstorage.NewDishLocalStorage()
	// kitchenRepo := kitchenstorage.NewIngredientLocalStorage()

	// kuc := kitchenusecase.NewUseCase(kitchenRepo)

	return &Server{

		// ShopUC:    shopusecase.NewDishUseCase(shopRepo),
		// KitchenUC: kuc,
		KitchenHandler: kitchenhttp.NewHandler(),
		ShopHandler:    shophttp.NewHandler(),
	}
}

func (s *Server) Run() error {

	e := echo.New()
	eg := e.Group("/tavern")

	kitchenhttp.RegisterHTTPEndpoints(eg, s.KitchenHandler)
	shophttp.RegisterHTTPEndpoints(eg, s.ShopHandler)

	return e.Start("localhost:8080")
}
