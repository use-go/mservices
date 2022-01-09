package handler

import (
	"context"

	"comm/logger"
	"comm/auth"

	user "proto/user"
)

// Call is a single request handler called via client.Call or the generated client code
func (h *Handler) Call(ctx context.Context, req *user.Request, rsp *user.Response) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Call", acc.Name)
	}
	logger.Info(ctx, "Received User.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
