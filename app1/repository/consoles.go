package repository

import "app1/domain"

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

func (cr *consolesRepository) Upsert(console *domain.Console) error {
	cr.consoles[console.ID] = console
	return nil
}
