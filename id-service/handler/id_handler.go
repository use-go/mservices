package handler

import (
	"context"
	"fmt"

	"comm/auth"
	"comm/errors"
	"comm/logger"
	"comm/mark"
	"proto/id"

	"github.com/google/uuid"
	"github.com/teris-io/shortid"
)

func (h *Handler) Generate(ctx context.Context, req *id.GenerateRequest, rsp *id.GenerateResponse) error {
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "Generate")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Generate", acc.Name)
	}
	if len(req.Type) == 0 {
		req.Type = "uuid"
	}

	switch req.Type {
	case "uuid":
		rsp.Type = "uuid"
		rsp.Id = uuid.New().String()
	case "snowflake":
		id, err := h.Snowflake.Mint()
		if err != nil {
			logger.Errorf(ctx, "Failed to generate snowflake id: %v", err)
			return errors.InternalServerError("id.generate", "failed to mint snowflake id")
		}
		rsp.Type = "snowflake"
		rsp.Id = fmt.Sprintf("%v", id)
	case "bigflake":
		id, err := h.Bigflake.Mint()
		if err != nil {
			logger.Errorf(ctx, "Failed to generate bigflake id: %v", err)
			return errors.InternalServerError("id.generate", "failed to mint bigflake id")
		}
		rsp.Type = "bigflake"
		rsp.Id = fmt.Sprintf("%v", id)
	case "shortid":
		id, err := shortid.Generate()
		if err != nil {
			logger.Errorf(ctx, "Failed to generate shortid id: %v", err)
			return errors.InternalServerError("id.generate", "failed to generate short id")
		}
		rsp.Type = "shortid"
		rsp.Id = id
	default:
		return errors.BadRequest("id.generate", "unsupported id type")
	}

	return nil
}

func (h *Handler) Types(ctx context.Context, req *id.TypesRequest, rsp *id.TypesResponse) error {
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "Types")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Types", acc.Name)
	}
	rsp.Types = []string{
		"uuid",
		"shortid",
		"snowflake",
		"bigflake",
	}
	return nil
}
