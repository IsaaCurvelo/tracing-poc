package main

import (
	"app2/handler"
	"app2/repository"
	"app2/usecase"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	// spin up dependencies
	vendorsRepository := repository.NewVendorsRepository()
	retrieveVendorUseCase := usecase.NewRetrieveVendorUseCase(vendorsRepository)
	vendorsHandler := handler.NewVendorsHandler(retrieveVendorUseCase)

	//create gin engine
	engine := gin.New()
	engine.Use(gin.Recovery())

	engine.POST("/vendors/:vendor-id", vendorsHandler.HandleRetrieveVendor)

	httpServer := http.Server{Handler: engine, Addr: ":8080"}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)
	go func() {
		errCh <- httpServer.ListenAndServe()
	}()

	select {
	case <-sigCh:
		log.Println("\ngoodbye")
		return
	case err := <-errCh:
		if err != nil {
			log.Fatal(err)
		}
	}

}
