package db

import (
	"context"
	"fmt"

	"comm/config"
	"comm/errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Conf struct {
	User   string `json:"user"`
	Passwd string `json:"passwd"`
	Host   string `json:"host"`
	Port   int64  `json:"port"`
	DB     string `json:"db"`
}

func InitDb(ctx context.Context, i ...uint64) (*gorm.DB, error) {
	db := []Conf{}
	v, err := config.Get("db")
	if err != nil {
		return nil, err
	}
	err = v.Scan(&db)
	if err != nil {
		return nil, err
	}

	if len(db) == 0 {
		return nil, errors.InternalServerError("db not found")
	}

	c := Conf{}
	idx := uint64(0)

	if len(i) > 0 {
		idx = i[0]
	}
	c = db[idx]

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", c.User, c.Passwd, c.Host, c.Port, c.DB)
	cfg := mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}
	return gorm.Open(mysql.New(cfg), &gorm.Config{})
}
