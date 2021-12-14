package encoding

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"

	"github.com/micro/micro/v3/service/debug/trace"
)

// see https://stackoverflow.com/questions/28595664/how-to-stop-json-marshal-from-escaping-and/28596225
func JSONMarshal(ctx context.Context, t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	traceID, _, _ := trace.FromContext(ctx)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	rsp := bytes.TrimRight(buffer.Bytes(), "\n")
	if strings.HasPrefix(string(rsp), "{") {
		rsp = []byte(strings.Replace(
			strings.Replace(string(rsp), "{", "{\"code\": 200,", 1),
			"{",
			"{\"request_id\": \""+traceID+"\",", 1),
		)
	}
	return rsp, err
}
