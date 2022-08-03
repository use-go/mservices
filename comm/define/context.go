package define

import (
	"context"

	"github.com/2637309949/micro/v3/util/ctx"
)

var (
	TODO       = ctx.FromContext(context.TODO())
	Background = ctx.FromContext(context.Background())
)
