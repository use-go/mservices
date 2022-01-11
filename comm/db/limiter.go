package db

import (
	"context"
	"reflect"

	"github.com/jinzhu/gorm"
)

type Offset interface {
	GetOffset() int32
	GetLimit() int32
}

type Pager interface {
	GetPageSize() int32
	GetPageNo() int32
}

func SetZero(t interface{}, key string, v interface{}) {
	field := reflect.ValueOf(t).Elem().FieldByName(key)
	if field.CanSet() {
		if reflect.Zero(reflect.TypeOf(v)).Interface() == field.Interface() {
			field.Set(reflect.ValueOf(v))
		}
	}
}

func SetLimit(ctx context.Context, db *gorm.DB, limiter interface{}) *gorm.DB {
	if l, ok := limiter.(Offset); ok {
		SetZero(limiter, "Offset", 0)
		SetZero(limiter, "Limit", 20)
		db = db.Offset(l.GetOffset()).Limit(l.GetLimit())
	} else if l, ok := limiter.(Pager); ok {
		SetZero(limiter, "PageNo", 1)
		SetZero(limiter, "PageSize", 20)
		db = db.Limit(l.GetPageSize()).Offset(l.GetPageSize() * (l.GetPageNo() - 1))
	}
	return db
}
