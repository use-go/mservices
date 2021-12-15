package trace

import (
	"context"

	"github.com/2637309949/micro/v3/service/debug/trace"
)

func Extract(ctx context.Context) (traceID string, parentSpanID string, isFound bool) {
	return trace.FromContext(ctx)
}

func ExtractTraceID(ctx context.Context) string {
	traceID, _, _ := Extract(ctx)
	return traceID
}
