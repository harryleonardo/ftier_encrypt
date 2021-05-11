package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	SharedConfig "github.com/ftier-encrypt/shared/config"
	SharedContext "github.com/ftier-encrypt/shared/context"

	encryptHandler "github.com/ftier-encrypt/domain/encrypt/delivery/http"
	encryptUsecase "github.com/ftier-encrypt/domain/encrypt/usecase"
)

func main() {
	// - initialize echo labstack as a framework that i'm using;
	e := echo.New()

	// - initiate config;
	conf := SharedConfig.GetDefaultImmutableConfig()

	// - CORS
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	// - initialize customize context;
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ac := &SharedContext.ApplicationContext{
				Context: c,
				Config:  conf,
			}
			return h(ac)
		}
	})

	encryptUcase := encryptUsecase.NewEncryptUsecase()

	encryptHandler.EncryptHandler(e, encryptUcase)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.GetPort())))
}
