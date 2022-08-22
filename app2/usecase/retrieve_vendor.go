package usecase

import (
	"app2/domain"
	"context"
	"go.opentelemetry.io/otel"
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

func (r *retrieveVendor) Execute(ctx context.Context, ID string) (*domain.Vendor, error) {
	ctx, span := otel.Tracer("app2").Start(ctx, "retrieveVendor.Execute")
	defer span.End()

	vendor, err := r.vendorRepository.FindByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	exclusiveTitles, err := r.exclusiveTitlesIntegration.GetByVendorID(ctx, vendor.ID)
	if err != nil {
		return nil, err
	}

	vendor.ExclusiveTitles = exclusiveTitles

	return vendor, nil
}
