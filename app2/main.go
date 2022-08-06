package main

import (
	"app2/handler"
	"app2/integration"
	"app2/pb/exclusive_titles_pb"
	"app2/repository"
	"app2/usecase"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	// spin up dependencies
	// grpc client
	dialOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial("localhost:8083", dialOption)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	client := exclusive_titles_pb.NewExclusiveTitlesServiceClient(conn)

	// application components
	vendorsRepository := repository.NewVendorsRepository()
	exclusiveTitlesIntegration := integration.NewExclusiveTitlesIntegration(client)
	retrieveVendorUseCase := usecase.NewRetrieveVendorUseCase(vendorsRepository, exclusiveTitlesIntegration)
	vendorsHandler := handler.NewVendorsHandler(retrieveVendorUseCase)

	//create gin engine
	engine := gin.New()
	engine.Use(gin.Recovery())

	engine.POST("/vendors/:vendor-id", vendorsHandler.HandleRetrieveVendor)

	httpServer := http.Server{Handler: engine, Addr: ":8082"}

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
