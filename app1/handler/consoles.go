package handler

import (
	"app1/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	CreateConsoleUseCase interface {
		Execute(console *domain.Console) error
	}
	consolesHandler struct {
		createConsoleUseCase CreateConsoleUseCase
	}
)

func NewConsolesHandler(useCase CreateConsoleUseCase) *consolesHandler {
	return &consolesHandler{createConsoleUseCase: useCase}
}

func (ch *consolesHandler) HandleCreateConsole(ctx *gin.Context) {
	consoleRequest := &domain.Console{}

	err := ctx.ShouldBindJSON(consoleRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, struct {
			message string
		}{message: err.Error()})
		return
	}

	err = ch.createConsoleUseCase.Execute(consoleRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, struct {
			message string
		}{message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, consoleRequest)
}
