package main

import (
	"app1/handler"
	"app1/handler/middleware"
	"app1/integration"
	"app1/repository"
	"app1/usecase"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.9.0"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	vendorsIntegrationHostEnv = "VENDORS_INTEGRATION_HOST"
	tracingCollectorHostEnv   = "TRACING_COLLECTOR_HOST"
	localhost                 = "localhost"
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
			semconv.ServiceNameKey.String("app1"),
		)),
	)
	return tp, nil
}

func main() {
	// retrieve env variables
	vendorsIntegrationHost := os.Getenv(vendorsIntegrationHostEnv)
	if vendorsIntegrationHost == "" {
		vendorsIntegrationHost = localhost
	}

	tracingCollectorHost := os.Getenv(tracingCollectorHostEnv)
	if tracingCollectorHost == "" {
		tracingCollectorHost = localhost
	}
	fmt.Printf("resolved tracing collector host to be %v\n", tracingCollectorHost)

	tp, err := tracerProvider(fmt.Sprintf("http://%v:14268/api/traces", tracingCollectorHost))
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

	// spin up components
	consolesRepository := repository.NewConsolesRepository()
	vendorsIntegration := integration.NewVendorsIntegration(vendorsIntegrationHost)
	createConsoleUseCase := usecase.NewCreateConsoleUseCase(consolesRepository, vendorsIntegration)
	consolesHandler := handler.NewConsolesHandler(createConsoleUseCase)

	//create gin engine
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.BindRequestIDToCtx())

	engine.GET("/consoles", consolesHandler.HandleGetAllConsoles)

	otelHandler := otelhttp.NewHandler(engine,
		"httpHandler.request_received",
		otelhttp.WithTracerProvider(tp),
		otelhttp.WithPropagators(b3.New()))

	httpServer := http.Server{Handler: otelHandler, Addr: ":8081"}

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
