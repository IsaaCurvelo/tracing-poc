package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"net/http"
)

type contextKey string

const (
	ReqIDHeader            = "x-request-id"
	RequestID   contextKey = "request-id"
)

func HandleTracingHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := otel.GetTextMapPropagator()
		tracingContext := t.Extract(ctx.Request.Context(), propagation.HeaderCarrier(ctx.Request.Header))
		ctx.Request = ctx.Request.WithContext(tracingContext)

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
