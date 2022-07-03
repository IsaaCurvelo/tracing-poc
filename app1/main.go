package main

import (
	"app1/handler"
	"app1/integration"
	"app1/repository"
	"app1/usecase"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// spin up components
	consolesRepository := repository.NewConsolesRepository()
	vendorsIntegration := integration.NewVendorsIntegration()
	createConsoleUseCase := usecase.NewCreateConsoleUseCase(consolesRepository, vendorsIntegration)
	consolesHandler := handler.NewConsolesHandler(createConsoleUseCase)

	//create gin engine
	engine := gin.New()
	engine.Use(gin.Recovery())

	engine.POST("/consoles", consolesHandler.HandleCreateConsole)

	httpServer := http.Server{Handler: engine, Addr: ":8081"}

	err := httpServer.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
