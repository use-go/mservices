package handler

import (
	"bytes"
	"comm/auth"
	"comm/errors"
	"comm/logger"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"proto/avatar"

	"github.com/o1egl/govatar"
)

// DeleteInfo defined TODO
func (h *Handler) Generate(ctx context.Context, req *avatar.GenerateRequest, rsp *avatar.GenerateResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Generate", acc.Name)
	}

	var gender govatar.Gender

	// gender, default is `male`
	if req.Gender == "male" {
		gender = govatar.MALE
	} else if req.Gender == "female" {
		gender = govatar.FEMALE
	} else {
		gender = govatar.MALE
	}

	// generate avatar
	var avatarImg image.Image
	var err error

	if req.Username == "" {
		avatarImg, err = govatar.Generate(gender)
	} else {
		avatarImg, err = govatar.GenerateForUsername(gender, req.Username)
	}
	if err != nil {
		return errors.InternalServerError("avatar.generate", "generate avatarImg error: %v", err)
	}

	// format avatar image, default is `jpeg`
	format := req.Format
	if format != "png" && format != "jpeg" {
		format = "jpeg"
	}

	buf := bytes.NewBuffer(nil)
	if format == "png" {
		err = png.Encode(buf, avatarImg)
	} else {
		err = jpeg.Encode(buf, avatarImg, nil)
	}
	if err != nil {
		return errors.InternalServerError("avatar.generate", "encode avatar image error: %v", err)
	}

	base64String := fmt.Sprintf("data:image/%s;base64,%s", format, base64.StdEncoding.EncodeToString(buf.Bytes()))

	if !req.Upload {
		rsp.Base64 = base64String
		return nil
	}
	return nil
}
