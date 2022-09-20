package integration

import (
	"app1/domain"
	"context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"io"
	"log"
	"net/http"
)

type vendorsIntegration struct {
}

func NewVendorsIntegration() *vendorsIntegration {
	return &vendorsIntegration{}
}

func (vi *vendorsIntegration) GetByID(ctx context.Context, ID string) (*domain.Vendor, error) {
	ctx, span := otel.Tracer("app1").Start(ctx, "vendorsIntegration.GetByID")
	defer span.End()

	header := http.Header{}
	propagator := otel.GetTextMapPropagator()
	propagator.Inject(ctx, propagation.HeaderCarrier(header))

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8082/vendors/%v", ID), nil)
	if err != nil {
		return nil, err
	}

	request.Header = header

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal()
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	retrievedVendor := &domain.Vendor{}
	err = json.Unmarshal(responseBody, retrievedVendor)
	if err != nil {
		return nil, err
	}

	return retrievedVendor, nil
}
