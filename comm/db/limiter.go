package db

import (
	"comm/define"
	"context"
	"reflect"

	"gorm.io/gorm"
)

var (
	OffsetDefault = 0
	LimitDefault  = 20
	PageNoDefault = 1
)

type Offset interface {
	GetOffset() int
	GetLimit() int
}

type Pager interface {
	GetPageSize() int
	GetPageNo() int
}

func SetZero(t interface{}, key string, v interface{}) {
	field := reflect.ValueOf(t).Elem().FieldByName(key)
	if field.CanSet() {
		if reflect.Zero(reflect.TypeOf(v)).Interface() == field.Interface() {
			field.Set(reflect.ValueOf(v))
		}
	}
}

func SetLimit(_ context.Context, db *gorm.DB, limiter interface{}) *gorm.DB {
	if l, ok := limiter.(Offset); ok {
		SetZero(limiter, "Offset", OffsetDefault)
		SetZero(limiter, "Limit", LimitDefault)
		db = db.Offset(l.GetOffset()).Limit(l.GetLimit())
	} else if l, ok := limiter.(Pager); ok {
		SetZero(limiter, "PageNo", PageNoDefault)
		SetZero(limiter, "PageSize", LimitDefault)
		db = db.Limit(l.GetPageSize()).Offset(l.GetPageSize() * (l.GetPageNo() - 1))
	}
	return db
}

func SetOrder(_ context.Context, db *gorm.DB, o order, tb ...string) *gorm.DB {
	strOrder := o.GetOrderCol()
	if len(strOrder) > 0 {
		if len(tb) > 0 {
			strOrder = tb[0] + "." + strOrder
		}
		switch o.GetOrderType() {
		case define.ORDER_ASC:
			strOrder += " ASC"
		case define.ORDER_DESC:
			strOrder += " DESC"
		default:
			strOrder += " ASC"
		}

		db = db.Order(strOrder)
	}

	return db
}
