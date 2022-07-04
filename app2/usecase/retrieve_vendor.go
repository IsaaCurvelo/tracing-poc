package usecase

import "app2/domain"

type (
	VendorRepository interface {
		FindByID(string) (*domain.Vendor, error)
	}

	retrieveVendorUseCase struct {
		vendorRepository VendorRepository
	}
)

func NewRetrieveVendorUseCase(vendorRepository VendorRepository) *retrieveVendorUseCase {
	return &retrieveVendorUseCase{vendorRepository: vendorRepository}
}

func (r *retrieveVendorUseCase) Execute(ID string) (*domain.Vendor, error) {
	return r.vendorRepository.FindByID(ID)
}
