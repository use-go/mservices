package trace

import (
	"bytes"
	"comm/logger"
	"comm/util/encode"
	"context"
	"strings"
	"time"
)

// Debug defined TODO
func Debug(ctx context.Context, action string, req, rsp interface{}) func() {
	startTime := time.Now()
	reqByte := encode.MustMarshal(req)
	reqStr := strings.Replace(strings.Replace(string(reqByte), " ", "", -1), "\n", "", -1)
	logger.Init(logger.WithCallerSkipCount(2))
	defer logger.Init(logger.WithCallerSkipCount(logger.DefaultCallerSkipCount))
	logger.Infof(ctx, ">>>>> Received %v request = %v", action, reqStr)
	return func() {
		if rspBuffer, ok := rsp.(*bytes.Buffer); ok {
			rsp = rspBuffer.Bytes()
		}
		rspByte := encode.MustMarshal(rsp)
		rspStr := strings.Replace(strings.Replace(string(rspByte), " ", "", -1), "\n", "", -1)
		cost := int(time.Since(startTime) / time.Microsecond)
		logger.Init(logger.WithCallerSkipCount(2))
		defer logger.Init(logger.WithCallerSkipCount(logger.DefaultCallerSkipCount))
		logger.Infof(ctx, "<<<<< Finished %v %vms response = %v", action, cost, rspStr)
	}
}
