package usecase

import (
	"app3/domain"
	"context"
)

type (
	ExclusiveTitlesRepository interface {
		GetByVendorID(context.Context, string) ([]domain.ExclusiveTitle, error)
	}

	getExclusiveTitlesByVendorIDUsecase struct {
		exclusiveTitlesRepository ExclusiveTitlesRepository
	}
)

func NewRetrieveVendorUseCase(vendorRepository ExclusiveTitlesRepository) *getExclusiveTitlesByVendorIDUsecase {
	return &getExclusiveTitlesByVendorIDUsecase{exclusiveTitlesRepository: vendorRepository}
}

func (r *getExclusiveTitlesByVendorIDUsecase) Execute(context context.Context, vendorID string) ([]domain.ExclusiveTitle, error) {
	return r.exclusiveTitlesRepository.GetByVendorID(context, vendorID)
}
