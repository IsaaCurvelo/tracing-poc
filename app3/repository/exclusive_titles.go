package repository

import (
	"app3/domain"
	"context"
	"fmt"
)

type exclusiveTitlesRepository struct {
	exclusiveTitles map[string][]domain.ExclusiveTitle
}

func NewExclusiveTitlesRepository() *exclusiveTitlesRepository {
	return &exclusiveTitlesRepository{exclusiveTitles: make(map[string][]domain.ExclusiveTitle)}
}

func (vr *exclusiveTitlesRepository) GetByVendorID(context context.Context, vendorID string) ([]domain.ExclusiveTitle, error) {
	if exclusiveTitles, ok := vr.exclusiveTitles[vendorID]; !ok {
		return nil, fmt.Errorf("could not find titles for vendorID %v", vendorID)
	} else {
		return exclusiveTitles, nil
	}
}
