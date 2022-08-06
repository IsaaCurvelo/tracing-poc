package repository

import (
	"app2/domain"
	"context"
	"fmt"
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

func (vr *vendorsRepository) FindByID(_ context.Context, ID string) (*domain.Vendor, error) {
	if value, ok := vr.vendors[ID]; !ok {
		return nil, fmt.Errorf("could not find vendor of id %v", ID)
	} else {
		return value, nil
	}
}
