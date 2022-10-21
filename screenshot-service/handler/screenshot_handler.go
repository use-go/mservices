package handler

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"comm/auth"
	"comm/errors"
	"comm/logger"
	"comm/mark"
	"comm/service"
	"proto/id"
	"proto/screenshot"
)

var (
	screenshotPath = "/usr/src/app"
)

// Screenshot defined todo
func (h *Handler) Screenshot(ctx context.Context, req *screenshot.ScreenshotRequest, rsp *screenshot.ScreenshotResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "Screenshot")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Screenshot", acc.Name)
	}

	gtResp, err := h.IdService.Generate(ctx, &id.GenerateRequest{Type: "uuid"})
	if err != nil {
		logger.Error(ctx, err.Error())
		return errors.InternalServerError(service.GetName(), "generate id error")
	}
	imageName := gtResp.Id + ".png"
	imagePath := filepath.Join(screenshotPath, imageName)
	defer func() { os.Remove(imagePath) }()
	width, height := "800", "600"
	if req.Width != 0 {
		width = fmt.Sprintf("%v", req.Width)
	}
	if req.Height != 0 {
		height = fmt.Sprintf("%v", req.Height)
	}
	cmd := exec.Command("/usr/bin/chromium-browser",
		"--headless", "--window-size="+width+","+height, "--no-sandbox", "--screenshot="+imagePath,
		"--hide-scrollbars", "--disable-setuid-sandbox", "--single-process", "--no-zygote", "--disable-gpu", req.Url)

	outp, err := cmd.CombinedOutput()
	timemark.Mark("Command")
	logger.Info(ctx, string(outp))
	if err != nil {
		logger.Error(ctx, string(outp)+err.Error())
		return errors.InternalServerError(service.GetName(), "error taking screenshot")
	}
	file, err := ioutil.ReadFile(imagePath)
	if err != nil {
		logger.Errorf(ctx, "error reading file %s", err)
		return errors.InternalServerError(service.GetName(), "error taking screenshot")
	}
	base := base64.StdEncoding.EncodeToString(file)
	rsp.Data = []byte("data:image/png;base64, " + base)
	return nil
}
