package db

import (
	"context"
	"fmt"
	"sync"

	"comm/config"
	"comm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbs []*gorm.DB
var once sync.Once

type Conf struct {
	User   string `json:"user"`
	Passwd string `json:"passwd"`
	Host   string `json:"host"`
	Port   int64  `json:"port"`
	DB     string `json:"db"`
}

func scan(ctx context.Context) func() {
	return func() {
		cfs := []Conf{}
		v, err := config.Get("db")
		if err != nil {
			return
		}

		err = v.Scan(&cfs)
		if err != nil {
			return
		}

		if len(dbs) == 0 {
			logger.Errorf(ctx, "not found db")
			return
		}

		for i := range cfs {
			v := cfs[i]
			dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", v.User, v.Passwd, v.Host, v.Port, v.DB)
			logger.Infof(ctx, "%v", dsn)
			db, err := gorm.Open(mysql.New(mysql.Config{
				DSN:                       dsn,   // data source name
				DefaultStringSize:         256,   // default size for string fields
				DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
				DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
				DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
				SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
			}), &gorm.Config{})
			if err != nil {
				logger.Errorf(ctx, "%v", err)
				continue
			}
			dbs = append(dbs, db)
		}
	}
}

func InitDb(ctx context.Context, i ...uint64) (*gorm.DB, error) {
	once.Do(scan(ctx))
	idx := uint64(0)
	if len(i) > 0 && len(dbs) >= int(i[0]) {
		idx = i[0]
	}
	return dbs[idx], nil
}
