package store

import (
	"github.com/2637309949/micro/v3/service/store"
)

type (
	Record     = store.Record
	ReadOption = store.ReadOption
	ListOption = store.ListOption
)

func Read(key string, opts ...ReadOption) ([]*store.Record, error) {
	return store.Read(key, opts...)
}

func Write(r *Record) error {
	return store.Write(r)
}

func Delete(key string) error {
	return store.Delete(key)
}

func List(opts ...ListOption) ([]string, error) {
	return store.List(opts...)
}
