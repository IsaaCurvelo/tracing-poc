package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type contextKey string

const (
	ReqIDHeader            = "x-request-id"
	RequestID   contextKey = "request-id"
)

func BindRequestIDToCtx() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Request = BindRequestContext(ctx.Request)
		ctx.Next()
	}
}

func BindRequestContext(r *http.Request) *http.Request {
	var requestID = r.Header.Get(ReqIDHeader)
	if requestID == "" {
		requestID = uuid.New().String()
	}

	return r.WithContext(context.WithValue(r.Context(), RequestID, requestID))
}
