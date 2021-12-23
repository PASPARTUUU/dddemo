package server

import (
	"github.com/labstack/echo/v4"

	servicehttp "dddemo/service/delivery/http"
	//---
	kitchenhttp "dddemo/domains/kitchen/delivery/http"
	kitchenstorage "dddemo/domains/kitchen/repository/localstorage"

	// kitchenhttp "dddemo/domains/kitchen/delivery/http"
	// kitchenstorage "dddemo/domains/kitchen/repository/localstorage"
	// kitchenusecase "dddemo/domains/kitchen/usecase"
	// ---
	shophttp "dddemo/domains/shop/delivery/http"
	shopstorage "dddemo/domains/shop/repository/localstorage"
	// shopstorage "dddemo/domains/shop/repository/localstorage"
	// shopusecase "dddemo/domains/shop/usecase"
)

type Server struct {
	Service *servicehttp.Handler
	// ShopUC    *shopusecase.DishUseCase
	// KitchenUC *kitchenusecase.UseCase

	KitchenHandler *kitchenhttp.Handler
	ShopHandler    *shophttp.Handler
}

func NewServer() *Server {

	shopDishRepo := shopstorage.NewDishLocalStorage()
	kitchenIngrRepo := kitchenstorage.NewIngredientLocalStorage()

	// kuc := kitchenusecase.NewUseCase(kitchenRepo)

	return &Server{
		Service: servicehttp.NewHTTPHandlerServices(kitchenIngrRepo, shopDishRepo),

		// ShopUC:    shopusecase.NewDishUseCase(shopRepo),
		// KitchenUC: kuc,
		KitchenHandler: kitchenhttp.NewHandler(kitchenIngrRepo),
		ShopHandler:    shophttp.NewHandler(shopDishRepo),
	}
}

func (s *Server) Run() error {

	e := echo.New()
	eg := e.Group("/tavern")

	servicehttp.RegisterHTTPEndpoints(eg, s.Service)
	kitchenhttp.RegisterHTTPEndpoints(eg, s.KitchenHandler)
	shophttp.RegisterHTTPEndpoints(eg, s.ShopHandler)

	return e.Start("localhost:8080")
}
