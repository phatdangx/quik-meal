package app

import (
	"context"
	"time"
	"user/handler"
	"user/repository/mongodb"
	"user/transport/http/routes"
	"user/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Start() {
	e := echo.New()

	// init db client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Error("can not connect mongodb")
	}

	// init repository
	repository := mongodb.NewUserMongoRepo(client)

	// register usecase
	usecase := usecase.NewUserUseCase(repository)

	// register handler
	privateHandler := handler.NewPrivateHandler(usecase)

	// register routes
	routes.Private(e, privateHandler)

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
