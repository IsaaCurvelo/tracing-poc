package integration

import (
	"app1/domain"
	"context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel"
	"io/ioutil"
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

	response, err := http.Get(fmt.Sprintf("http://localhost:8082/vendors/%v", ID))
	if err != nil {
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
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
