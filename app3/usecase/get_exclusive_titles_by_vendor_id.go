package usecase

import (
	"app3/domain"
	"context"
)

type (
	ExclusiveTitlesRepository interface {
		GetByVendorID(context.Context, string) ([]domain.ExclusiveTitle, error)
	}

	getExclusiveTitlesByVendorID struct {
		exclusiveTitlesRepository ExclusiveTitlesRepository
	}
)

func NewGetExclusiveTitlesByVendorIDUsecase(exclusiveTitlesRepository ExclusiveTitlesRepository) *getExclusiveTitlesByVendorID {
	return &getExclusiveTitlesByVendorID{exclusiveTitlesRepository: exclusiveTitlesRepository}
}

func (r *getExclusiveTitlesByVendorID) Execute(context context.Context, vendorID string) ([]domain.ExclusiveTitle, error) {
	return r.exclusiveTitlesRepository.GetByVendorID(context, vendorID)
}
