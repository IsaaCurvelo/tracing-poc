package repository

import (
	"app2/domain"
	"fmt"
)

type vendorsRepository struct {
	vendors map[string]*domain.Vendor
}

func NewVendorsRepository() *vendorsRepository {
	return &vendorsRepository{vendors: make(map[string]*domain.Vendor)}
}

func (vr *vendorsRepository) findByID(ID string) (*domain.Vendor, error) {
	if value, ok := vr.vendors[ID]; !ok {
		return nil, fmt.Errorf("could not find vendor of id %v", ID)
	} else {
		return value, nil
	}
}
