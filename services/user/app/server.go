package app

import (
	"user/transport/http/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Start() {
	e := echo.New()

	// connect db

	// register handler

	// register routes
	routes.Private(e, handler)

	// Start server
	go func() {
		port := "8080"
		log.Info("start at port " + port)
		err := e.Start(":" + port)
		if err != nil {
			log.Error(err)
		}
	}()
}
