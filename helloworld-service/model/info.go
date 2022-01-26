package model

import (
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Info struct {
	Id        uint32         `gorm:"column:id" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (h *Info) TableName() string {
	return "info"
}

func (h *Info) Marshal(o interface{}) error {
	err := copier.Copy(o, h)
	if err != nil {
		return err
	}
	return nil
}

func (h *Info) Unmarshal(o interface{}) error {
	err := copier.Copy(h, o)
	if err != nil {
		return err
	}
	return nil
}

func InfoMarshalLst(toValue interface{}, o interface{}) error {
	err := copier.Copy(toValue, o)
	if err != nil {
		return err
	}
	return nil
}

func InfoUnmarshalLst(o interface{}, toValue interface{}) error {
	err := copier.Copy(toValue, o)
	if err != nil {
		return err
	}
	return nil
}
