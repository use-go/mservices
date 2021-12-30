package json

import "encoding/json"

func MustMarshal(v interface{}) []byte {
	if vb, ok := v.([]byte); ok {
		return vb
	}
	bytes, _ := json.Marshal(v)
	return bytes
}

func MustUnmarshal(data []byte, v interface{}) interface{} {
	json.Unmarshal(data, v)
	return v
}
