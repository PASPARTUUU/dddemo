package server

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"dddemo/pkg/meinlogger"
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

	// ---
	root "dddemo"
)

type Server struct {
	Service *servicehttp.Handler
	// ShopUC    *shopusecase.DishUseCase
	// KitchenUC *kitchenusecase.UseCase

	KitchenHandler *kitchenhttp.Handler
	ShopHandler    *shophttp.Handler

	// ---
	domains []Domain
	// ---
	templates meintemplate.Templates
	// ---
	embedWeb     embed.FS
	embedDomains embed.FS
}

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
		// ---
		embedWeb:     root.EmbedWeb,
		embedDomains: root.EmbedDomains,
	}

	server.templates, err = server.parseTemplates("web/templates")
	if err != nil {
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
	e.Use(meinlogger.LogrusMiddleware)
	// e.Logger = &meinlogger.MeinLogger{Logger: logrus.New()}

	// e.Use(middleware.Static("./static"))
	// e.Static("/web/static", "./web/static")
	// e.Static("/", "./")

	{ 
		// https://echo.labstack.com/cookbook/embed-resources/
		
		fsys, err := fs.Sub(s.embedWeb, "web/static")
		if err != nil {
			panic(err)
		}
		assetHandler := http.FileServer(http.FS(fsys))
		e.GET("/", echo.WrapHandler(assetHandler))
		e.GET("/web/static/*", echo.WrapHandler(http.StripPrefix("/web/static/", assetHandler)))

		for _, d := range s.domains {
			df := d.RootStaticFolder()
			dfsys, err := fs.Sub(s.embedDomains, df)
			if err != nil {
				panic(err)
			}
			e.GET(fmt.Sprint("/", df, "/*"), echo.WrapHandler(http.StripPrefix(fmt.Sprint("/", df, "/"), http.FileServer(http.FS(dfsys)))))
		}
	}

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
		Title:   "11111",
		Message: "222222",
	}

	err = s.templates.Render(ectx, 200, "hello", data)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
