package trace

import (
	"comm/logger"
	"comm/util/encode"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/2637309949/micro/v3/service/context/metadata"
)

// Debug defined TODO
func Debug(ctx context.Context, action string, req, rsp interface{}) func() {
	startTime := time.Now()
	reqByte := encode.MustMarshal(req)
	reqStr := strings.Replace(strings.Replace(string(reqByte), " ", "", -1), "\n", "", -1)
	srv, _ := metadata.Get(ctx, "Micro-From-Service")
	upstream, downstream := "", ""
	if len(srv) > 0 {
		upstream = fmt.Sprintf("%v>>%v", srv, action)
	} else {
		upstream = fmt.Sprintf("%v", action)
	}
	if len(srv) > 0 {
		downstream = fmt.Sprintf("%v<<%v", srv, action)
	} else {
		downstream = fmt.Sprintf("%v", action)
	}
	logger.Init(logger.WithCallerSkipCount(2))
	defer logger.Init(logger.WithCallerSkipCount(logger.DefaultCallerSkipCount))
	logger.Infof(ctx, ">>>>> Received %v request = %v", upstream, reqStr)
	return func() {

		rspByte := encode.MustMarshal(rsp)
		rspStr := strings.Replace(strings.Replace(string(rspByte), " ", "", -1), "\n", "", -1)
		cost := int(time.Since(startTime) / time.Microsecond)
		logger.Init(logger.WithCallerSkipCount(2))
		defer logger.Init(logger.WithCallerSkipCount(logger.DefaultCallerSkipCount))
		logger.Infof(ctx, "<<<<< Finished %v %vms response = %v", downstream, cost, rspStr)
	}
}
