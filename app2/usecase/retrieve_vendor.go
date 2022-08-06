package usecase

import (
	"app2/domain"
	"context"
)

type (
	VendorRepository interface {
		FindByID(context.Context, string) (*domain.Vendor, error)
	}

	ExclusiveTitlesIntegration interface {
		GetByVendorID(context context.Context, vendorID string) ([]domain.ExclusiveTitle, error)
	}

	retrieveVendorUseCase struct {
		vendorRepository           VendorRepository
		exclusiveTitlesIntegration ExclusiveTitlesIntegration
	}
)

func NewRetrieveVendorUseCase(vendorRepository VendorRepository, exclusiveTitlesIntegration ExclusiveTitlesIntegration) *retrieveVendorUseCase {
	return &retrieveVendorUseCase{vendorRepository: vendorRepository, exclusiveTitlesIntegration: exclusiveTitlesIntegration}
}

func (r *retrieveVendorUseCase) Execute(context context.Context, ID string) (*domain.Vendor, error) {
	vendor, err := r.vendorRepository.FindByID(context, ID)
	if err != nil {
		return nil, err
	}

	exclusiveTitles, err := r.exclusiveTitlesIntegration.GetByVendorID(context, vendor.ID)
	if err != nil {
		return nil, err
	}

	vendor.ExclusiveTitles = exclusiveTitles

	return vendor, nil
}
