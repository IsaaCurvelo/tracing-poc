package usecase

import (
	"app1/domain"
	"context"
)

type (
	VendorsIntegration interface {
		GetByID(context.Context, string) (*domain.Vendor, error)
	}

	ConsolesRepository interface {
		GetAll(context.Context) ([]domain.Console, error)
	}

	GetAllConsoles struct {
		VendorsIntegration VendorsIntegration
		ConsolesRepository ConsolesRepository
	}
)

func NewCreateConsoleUseCase(cr ConsolesRepository, vi VendorsIntegration) *GetAllConsoles {
	return &GetAllConsoles{
		VendorsIntegration: vi,
		ConsolesRepository: cr,
	}
}

func (c *GetAllConsoles) Execute(context context.Context) ([]domain.Console, error) {
	consoles, err := c.ConsolesRepository.GetAll(context)
	if err != nil {
		return nil, err
	}

	for i, console := range consoles {
		vendor, err := c.VendorsIntegration.GetByID(context, console.VendorID)
		if err != nil {
			return nil, err
		}

		consoles[i].Vendor = vendor
	}

	return consoles, nil
}
