package main

import (
	"app3/handler"
	"app3/pb/exclusive_titles_pb"
	"app3/repository"
	"app3/usecase"
	"context"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.9.0"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func tracerProvider(url string) (*trace.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("app3"),
		)),
	)
	return tp, nil
}

func main() {
	tp, err := tracerProvider("http://jaeger:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(b3.New())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer func(ctx context.Context) {
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	// spin up dependencies
	exclusiveTitlesRepository := repository.NewExclusiveTitlesRepository()
	retrieveVendorUseCase := usecase.NewGetExclusiveTitlesByVendorIDUsecase(exclusiveTitlesRepository)
	exclusiveTitlesHandler := handler.NewExclusiveTitlesHandler(retrieveVendorUseCase)

	//register gRPC service in the server
	server := grpc.NewServer(grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))
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
