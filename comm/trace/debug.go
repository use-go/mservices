package trace

import (
	"comm/logger"
	"comm/util/json"
	"context"
	"reflect"
	"strings"
	"time"
)

// panics if no exit `RequestId``
func Debug(ctx context.Context, action string, req, rsp interface{}) func() {
	traceID := ExtractTraceID(ctx)
	startTime := time.Now()
	reqByte := json.MustByte(req)
	reqStr := strings.Replace(strings.Replace(string(reqByte), " ", "", -1), "\n", "", -1)
	logger.Init(logger.WithCallerSkipCount(2))
	defer logger.Init(logger.WithCallerSkipCount(logger.DefaultCallerSkipCount))
	logger.Infof(">>>>> Received %v request %v\n%v", action, traceID, reqStr)
	v := reflect.Indirect(reflect.ValueOf(rsp))
	reqv := v.FieldByName("RequestId")
	codeV := v.FieldByName("Code")
	if reqv.CanSet() {
		reqv.SetString(traceID)
	}
	if codeV.CanSet() {
		codeV.SetInt(200)
	}
	return func() {
		rspByte := json.MustByte(rsp)
		rspStr := strings.Replace(strings.Replace(string(rspByte), " ", "", -1), "\n", "", -1)
		cost := int(time.Since(startTime) / time.Microsecond)
		logger.Init(logger.WithCallerSkipCount(2))
		defer logger.Init(logger.WithCallerSkipCount(logger.DefaultCallerSkipCount))
		logger.Infof("<<<<< Finished %v request %vms %v\n%v", action, cost, traceID, rspStr)
	}
}
