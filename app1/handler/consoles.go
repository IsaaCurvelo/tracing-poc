package handler

import (
	"app1/domain"
	"context"
	"github.com/gin-gonic/gin"
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
	allConsoles, err := ch.createConsoleUseCase.Execute(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, struct {
			message string
		}{message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, allConsoles)
}
