package usecase

import (
	"app1/domain"
	"context"
	"go.opentelemetry.io/otel"
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

func (c *GetAllConsoles) Execute(ctx context.Context) ([]domain.Console, error) {
	ctx, span := otel.Tracer("app1").Start(ctx, "GetAllConsoles.Execute")
	defer span.End()

	consoles, err := c.ConsolesRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for i, console := range consoles {
		vendor, err := c.VendorsIntegration.GetByID(ctx, console.VendorID)
		if err != nil {
			return nil, err
		}

		consoles[i].Vendor = vendor
	}

	return consoles, nil
}
