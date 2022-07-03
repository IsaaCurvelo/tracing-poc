package usecase

import "app1/domain"

type (
	VendorsIntegration interface {
		GetByID(ID string) (*domain.Vendor, error)
	}

	ConsolesRepository interface {
		Upsert(console *domain.Console) error
	}

	CreateConsoleUseCase struct {
		VendorsIntegration VendorsIntegration
		ConsolesRepository ConsolesRepository
	}
)

func NewCreateConsoleUseCase(cr ConsolesRepository, vi VendorsIntegration) *CreateConsoleUseCase {
	return &CreateConsoleUseCase{
		VendorsIntegration: vi,
		ConsolesRepository: cr,
	}
}

func (c *CreateConsoleUseCase) Execute(console *domain.Console) error {
	vendor, err := c.VendorsIntegration.GetByID(console.VendorID)
	if err != nil {
		return err
	}

	console.Vendor = vendor

	return c.ConsolesRepository.Upsert(console)
}
