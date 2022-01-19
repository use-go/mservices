package model

import "encoding/json"

func UnmarshalInfo(data []byte) (Info, error) {
	var r Info
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Info) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Info struct {
	Id   uint32 `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

// TableName sets the insert table name for this struct type
func (h *Info) TableName() string {
	return "info"
}
