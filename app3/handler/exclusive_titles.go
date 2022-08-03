package handler

import (
	"app3/domain"
	"app3/pb/exclusive_titles_pb"
	"context"
)

type (
	GetExclusiveTitlesByVendorIDUseCase interface {
		Execute(context.Context, string) ([]domain.ExclusiveTitle, error)
	}

	exclusiveTitlesHandler struct {
		exclusive_titles_pb.UnimplementedExclusiveTitlesServiceServer
		usecase GetExclusiveTitlesByVendorIDUseCase
	}
)

func NewExclusiveTitlesHandler(usecase GetExclusiveTitlesByVendorIDUseCase) *exclusiveTitlesHandler {
	return &exclusiveTitlesHandler{
		usecase: usecase,
	}
}

func (e exclusiveTitlesHandler) GetByVendorID(context context.Context, exclusiveTitlesRequest *exclusive_titles_pb.ExclusiveTitlesRequest) (
	*exclusive_titles_pb.ExclusiveTitlesResponse, error,
) {
	exclusiveTitles, err := e.usecase.Execute(context, exclusiveTitlesRequest.VendorId)

	if err != nil {
		return nil, err
	}

	exclusiveTitlesRsp := &exclusive_titles_pb.ExclusiveTitlesResponse{}

	exclusiveTitlesRsp.ExclusiveTitles = make([]*exclusive_titles_pb.ExclusiveTitle, len(exclusiveTitles))

	for i, et := range exclusiveTitles {
		exclusiveTitlesRsp.ExclusiveTitles[i] = &exclusive_titles_pb.ExclusiveTitle{
			Id:       et.ID,
			Name:     et.Name,
			VendorId: et.VendorID,
		}
	}

	return exclusiveTitlesRsp, nil
}
