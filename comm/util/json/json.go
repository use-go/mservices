package json

import "encoding/json"

func MustByte(v interface{}) []byte {
	bytes, _ := json.Marshal(v)

	return bytes
}
