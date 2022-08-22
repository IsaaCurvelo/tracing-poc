package repository

import (
	"app2/domain"
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
)

type vendorsRepository struct {
	vendors map[string]*domain.Vendor
}

func NewVendorsRepository() *vendorsRepository {
	return &vendorsRepository{vendors: map[string]*domain.Vendor{
		"1": {
			ID:            "1",
			Name:          "Sony",
			OriginCountry: "Japan",
		},
		"2": {
			ID:            "1",
			Name:          "Nintendo",
			OriginCountry: "Japan",
		},
		"3": {
			ID:            "1",
			Name:          "Xbox",
			OriginCountry: "United States of America",
		},
	}}
}

func (vr *vendorsRepository) FindByID(ctx context.Context, ID string) (*domain.Vendor, error) {
	ctx, span := otel.Tracer("app2").Start(ctx, "vendorsRepository.FindByID")
	defer span.End()

	if value, ok := vr.vendors[ID]; !ok {
		return nil, fmt.Errorf("could not find vendor of id %v", ID)
	} else {
		return value, nil
	}
}
