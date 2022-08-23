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
	ReqIDHeader                   = "x-request-id"
	TraceIDHeader                 = "x-b3-traceid"
	SpanIDHeader                  = "x-b3-spanid"
	ParentSpanIDHeader            = "x-b3-parentspanid"
	SampledIDHeader               = "x-b3-sampled"
	TraceCTXKey        contextKey = "tracing-context"
)

type Tracing struct {
	RequestID string
	B3Tracing
}

type B3Tracing struct {
	TraceID      string
	SpanID       string
	ParentSpanID string
	SampleID     string
}

func Context() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := otel.GetTextMapPropagator()
		tracingContext := t.Extract(ctx.Request.Context(), propagation.HeaderCarrier(ctx.Request.Header))
		ctx.Request = ctx.Request.WithContext(tracingContext)

		ctx.Request = BindRequestContext(ctx.Request)
		ctx.Next()
	}
}

func FromContext(ctx context.Context) Tracing {
	v := ctx.Value(TraceCTXKey)
	if v == nil {
		return Tracing{}
	}
	return v.(Tracing)
}

func WithTracing(ctx context.Context, tracing Tracing) context.Context {
	return context.WithValue(ctx, TraceCTXKey, tracing)
}

func BindRequestContext(r *http.Request) *http.Request {
	var requestID = r.Header.Get(ReqIDHeader)
	if requestID == "" {
		requestID = uuid.New().String()
	}

	t := Tracing{
		RequestID: requestID,
		B3Tracing: B3Tracing{
			TraceID:      r.Header.Get(TraceIDHeader),
			SpanID:       r.Header.Get(SpanIDHeader),
			ParentSpanID: r.Header.Get(ParentSpanIDHeader),
			SampleID:     r.Header.Get(SampledIDHeader),
		},
	}

	return r.WithContext(WithTracing(r.Context(), t))
}
