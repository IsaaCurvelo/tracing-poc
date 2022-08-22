package handler

import (
	"app2/domain"
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"net/http"
)

type (
	RetrieveVendorUseCase interface {
		Execute(context.Context, string) (*domain.Vendor, error)
	}
	vendorsHandler struct {
		retrieveVendorUseCase RetrieveVendorUseCase
	}
)

func NewVendorsHandler(retrieveVendorUseCase RetrieveVendorUseCase) *vendorsHandler {
	return &vendorsHandler{retrieveVendorUseCase: retrieveVendorUseCase}
}

func (v vendorsHandler) HandleRetrieveVendor(ctx *gin.Context) {
	tracingContext, span := otel.Tracer("app2").Start(ctx.Request.Context(), "vendorsHandler.HandleRetrieveVendor")
	defer span.End()

	request := &struct {
		ID string `uri:"vendor-id"`
	}{}

	err := ctx.ShouldBindUri(request)
	if err != nil {
		return
	}

	vendor, err := v.retrieveVendorUseCase.Execute(tracingContext, request.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, struct {
			message string
		}{message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, vendor)

}
