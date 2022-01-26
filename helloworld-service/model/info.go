package model

import (
	"encoding/json"
	"time"
)

func UnmarshalInfo(data []byte) (Info, error) {
	var r Info
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Info) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Info struct {
	Id        uint32     `gorm:"column:id" json:"id"`
	Name      string     `gorm:"column:name" json:"name"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (h *Info) TableName() string {
	return "info"
}
