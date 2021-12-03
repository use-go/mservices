package model

import "encoding/json"

func UnmarshalHelloworld(data []byte) (Helloworld, error) {
	var r Helloworld
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Helloworld) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Helloworld struct {
	Name string `gorm:"column:name" json:"name"`
}

// TableName sets the insert table name for this struct type
func (h *Helloworld) TableName() string {
	return "helloworld"
}
