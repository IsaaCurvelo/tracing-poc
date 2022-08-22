package handler

import (
	"app1/domain"
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"net/http"
)

type (
	CreateConsoleUseCase interface {
		Execute(context.Context) ([]domain.Console, error)
	}
	consolesHandler struct {
		createConsoleUseCase CreateConsoleUseCase
	}
)

func NewConsolesHandler(useCase CreateConsoleUseCase) *consolesHandler {
	return &consolesHandler{createConsoleUseCase: useCase}
}

func (ch *consolesHandler) HandleGetAllConsoles(ctx *gin.Context) {
	tracingContext, span := otel.Tracer("app1").Start(
		ctx.Request.Context(),
		"consolesHandler.HandleGetAllConsoles",
	)
	defer span.End()

	allConsoles, err := ch.createConsoleUseCase.Execute(tracingContext)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, struct {
			message string
		}{message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, allConsoles)
}
