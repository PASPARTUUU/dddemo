package server

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"dddemo/pkg/meintemplate"
	servicehttp "dddemo/service/delivery/http"

	//---
	"dddemo/domains/kitchen"
	kitchenhttp "dddemo/domains/kitchen/delivery/http"
	kitchenstorage "dddemo/domains/kitchen/repository/localstorage"
	"dddemo/domains/shop"

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

	// ---

	domains []Domain
	// templates map[string]*template.Template
	templates meintemplate.Templates
}

type TempPath struct{}

func NewServer() *Server {
	var err error

	shopDishRepo := shopstorage.NewDishLocalStorage()
	kitchenIngrRepo := kitchenstorage.NewIngredientLocalStorage()

	// kuc := kitchenusecase.NewUseCase(kitchenRepo)

	server := &Server{
		Service: servicehttp.NewHTTPHandlerServices(kitchenIngrRepo, shopDishRepo),

		// ShopUC:    shopusecase.NewDishUseCase(shopRepo),
		// KitchenUC: kuc,

		// ---
		domains: []Domain{shop.NewShop(), kitchen.NewKitchen()},
	}

	server.templates, err = server.parseTemplates("./web/templates")
	if err != nil {
		fmt.Printf("err server.parseTemplates: %v\n", err)
		panic(err)
	}

	kitchenHandler := kitchenhttp.NewHandler(server.templates, kitchenIngrRepo)
	shopHandler := shophttp.NewHandler(server.templates, shopDishRepo)

	server.KitchenHandler = kitchenHandler
	server.ShopHandler = shopHandler

	return server
}

func (s *Server) Run() error {

	e := echo.New()
	// e.Use(middleware.Static("./static"))
	e.Static("/web/static", "./web/static")

	eg := e.Group("/tavern")

	servicehttp.RegisterHTTPEndpoints(eg, s.Service)
	kitchenhttp.RegisterHTTPEndpoints(eg, s.KitchenHandler)
	shophttp.RegisterHTTPEndpoints(eg, s.ShopHandler)

	e.GET("/hello", s.hello)
	eg.GET("/hello", s.hello)

	return e.Start("localhost:8080")
}

func (s *Server) hello(ectx echo.Context) error {
	var err error
	data := struct {
		Title   string
		Message string
	}{
		Title:   "Последние заметки",
		Message: "Здесь пока ничего нет!",
	}

	err = s.templates.Render(ectx, 200, "aaa", data)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	return nil
}
