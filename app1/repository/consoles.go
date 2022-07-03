package repository

import "app1/domain"

type consolesRepository struct {
	consoles map[string]*domain.Console
}

func NewConsolesRepository() *consolesRepository {
	return &consolesRepository{consoles: make(map[string]*domain.Console)}
}

func (cr *consolesRepository) Upsert(console *domain.Console) error {
	cr.consoles[console.ID] = console
	return nil
}
