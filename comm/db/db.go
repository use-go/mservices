package db

import (
	"context"

	"comm/config"

	"github.com/jinzhu/gorm"
)

type Conf struct {
	User   string `json:"user"`
	Passwd string `json:"passwd"`
	Host   string `json:"host"`
	Port   int64  `json:"port"`
	DB     string `json:"db"`
}

func InitDb(ctx context.Context, index ...uint64) (*gorm.DB, error) {
	var db = []Conf{}
	v, err := config.Get("db")
	if err != nil {
		return nil, err
	}
	err = v.Scan(&db)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
