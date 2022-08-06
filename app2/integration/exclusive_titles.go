package integration

import (
	"app2/domain"
	"app2/pb/exclusive_titles_pb"
	"context"
	"google.golang.org/grpc"
)

type (
	exclusiveTitlesPB interface {
		GetByVendorID(context.Context, *exclusive_titles_pb.ExclusiveTitlesRequest, ...grpc.CallOption) (
			*exclusive_titles_pb.ExclusiveTitlesResponse, error)
	}
	exclusiveTitlesIntegration struct {
		exclusiveTitlesPB exclusiveTitlesPB
	}
)

func NewExclusiveTitlesIntegration(exclusiveTitlesPB exclusiveTitlesPB) *exclusiveTitlesIntegration {
	return &exclusiveTitlesIntegration{exclusiveTitlesPB: exclusiveTitlesPB}
}

func (e exclusiveTitlesIntegration) GetByVendorID(context context.Context, vendorID string) ([]domain.ExclusiveTitle, error) {
	pbResponse, err := e.exclusiveTitlesPB.GetByVendorID(context, &exclusive_titles_pb.ExclusiveTitlesRequest{VendorId: vendorID})
	if err != nil {
		return nil, err
	}

	exclusiveTitles := make([]domain.ExclusiveTitle, len(pbResponse.ExclusiveTitles))
	for i, et := range pbResponse.ExclusiveTitles {
		exclusiveTitles[i] = domain.ExclusiveTitle{
			ID:   et.Id,
			Name: et.Name,
		}
	}

	return exclusiveTitles, nil
}
