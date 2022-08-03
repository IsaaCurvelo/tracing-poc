package main

import (
	"app3/handler"
	"app3/pb/exclusive_titles_pb"
	"app3/repository"
	"app3/usecase"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	// spin up dependencies
	exclusiveTitlesRepository := repository.NewExclusiveTitlesRepository()
	retrieveVendorUseCase := usecase.NewGetExclusiveTitlesByVendorIDUsecase(exclusiveTitlesRepository)
	exclusiveTitlesHandler := handler.NewExclusiveTitlesHandler(retrieveVendorUseCase)

	//register gRPC service in the server
	server := grpc.NewServer()
	exclusive_titles_pb.RegisterExclusiveTitlesServiceServer(server, exclusiveTitlesHandler)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)
	go func() {
		listener, err := net.Listen("tcp", ":8083")
		if err != nil {
			log.Fatalf(err.Error())
		}
		errCh <- server.Serve(listener)
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
