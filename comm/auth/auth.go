package auth

import (
	"context"

	"github.com/2637309949/micro/v3/service/auth"
)

func FromContext(ctx context.Context) (*auth.Account, bool) {
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return nil, false
	}
	return acc, true
}
