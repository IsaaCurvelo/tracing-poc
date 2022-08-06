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
	return &exclusiveTitlesRepository{exclusiveTitles: map[string][]domain.ExclusiveTitle{
		"1": {
			{
				ID:       "gow",
				Name:     "God of War",
				VendorID: "1",
			},
			{
				ID:       "tlou",
				Name:     "The Last of Us",
				VendorID: "1",
			},
			{
				ID:       "uchtd",
				Name:     "Uncharted",
				VendorID: "1",
			},
		},
		"2": {
			{
				ID:       "botw",
				Name:     "The Legend of Zelda: Breath of the Wild",
				VendorID: "2",
			},
			{
				ID:       "dkykng",
				Name:     "Donkey Kong Country: Tropical Freeze",
				VendorID: "2",
			},
		},
		"3": {
			{
				ID:       "halo",
				Name:     "Halo",
				VendorID: "3",
			},
			{
				ID:       "forza",
				Name:     "Forza",
				VendorID: "3",
			},
			{
				ID:       "geow",
				Name:     "Gears of War",
				VendorID: "3",
			},
		},
	}}
}

func (vr *exclusiveTitlesRepository) GetByVendorID(_ context.Context, vendorID string) ([]domain.ExclusiveTitle, error) {
	if exclusiveTitles, ok := vr.exclusiveTitles[vendorID]; !ok {
		return nil, fmt.Errorf("could not find titles for vendorID %v", vendorID)
	} else {
		return exclusiveTitles, nil
	}
}
