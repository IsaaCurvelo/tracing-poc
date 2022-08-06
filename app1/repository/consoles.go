package repository

import (
	"app1/domain"
	"context"
)

type consolesRepository struct {
	consoles map[string]*domain.Console
}

func NewConsolesRepository() *consolesRepository {
	return &consolesRepository{consoles: map[string]*domain.Console{
		"1": {
			ID:         "1",
			VendorID:   "1",
			Name:       "Playstation 5",
			Generation: 9,
		},
		"2": {
			ID:         "2",
			VendorID:   "2",
			Name:       "Nintendo Switch",
			Generation: 8,
		},
		"3": {
			ID:         "3",
			VendorID:   "3",
			Name:       "Xbox Series X",
			Generation: 9,
		},
	}}
}

func (cr *consolesRepository) GetAll(_ context.Context) ([]domain.Console, error) {
	allConsoles := make([]domain.Console, len(cr.consoles))

	i := 0
	for _, v := range cr.consoles {
		allConsoles[i] = *v
		i++
	}

	return allConsoles, nil
}
