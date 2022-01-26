package db

import (
	"context"
	"fmt"

	"comm/config"

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
	var db = []Conf{}
	v, err := config.Get("db")
	if err != nil {
		return nil, err
	}
	err = v.Scan(&db)
	if err != nil {
		return nil, err
	}
	c := Conf{}
	if len(i) > 0 {
		c = db[i[0]]
	} else {
		c = db[0]
	}
	return gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", c.User, c.Passwd, c.Host, c.Port, c.DB), // data source name
		DefaultStringSize:         256,                                                                                                              // default size for string fields
		DisableDatetimePrecision:  true,                                                                                                             // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                                                             // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                                                             // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                                                            // auto configure based on currently MySQL version
	}), &gorm.Config{})
}
