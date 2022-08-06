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
		GetByVendorID(context.Context, string) ([]domain.ExclusiveTitle, error)
	}

	retrieveVendor struct {
		vendorRepository           VendorRepository
		exclusiveTitlesIntegration ExclusiveTitlesIntegration
	}
)

func NewRetrieveVendorUseCase(vendorRepository VendorRepository, exclusiveTitlesIntegration ExclusiveTitlesIntegration) *retrieveVendor {
	return &retrieveVendor{vendorRepository: vendorRepository, exclusiveTitlesIntegration: exclusiveTitlesIntegration}
}

func (r *retrieveVendor) Execute(context context.Context, ID string) (*domain.Vendor, error) {
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
