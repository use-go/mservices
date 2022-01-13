package encode

func MustSerialize(v interface{}) []byte {
	bytes, _ := Serialize(v)
	return bytes
}

func MustDeserialize(data []byte, v interface{}) {
	Deserialize(data, v)
}

func MustMarshal(v interface{}) []byte {
	bytes, _ := Marshal(v)
	return bytes
}

func MustUnmarshal(data []byte, v interface{}) {
	Unmarshal(data, v)
}
