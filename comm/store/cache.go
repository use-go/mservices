package store

import (
	"comm/util/encode"
	"context"
	"reflect"
	"time"

	"github.com/2637309949/micro/v3/service/client"
)

// Cache is the interface of a cache backend
type Cache interface {
	// Get retrieves an item from the cache. Returns the item or nil, and a bool indicating
	// whether the key was found.
	Get(ctx context.Context, key string, value interface{}, opts ...client.CallOption) error

	// Set sets an item to the cache, replacing any existing item.
	Set(ctx context.Context, key string, value interface{}, expire time.Duration, opts ...client.CallOption) error

	// Add adds an item to the cache only if an item doesn't already exist for the given
	// key, or if the existing item has expired. Returns an error otherwise.
	Add(ctx context.Context, key string, value interface{}, expire time.Duration, opts ...client.CallOption) error

	// Replace sets a new value for the cache key only if it already exists. Returns an
	// error if it does not.
	Replace(ctx context.Context, key string, data interface{}, expire time.Duration, opts ...client.CallOption) error

	// Delete removes an item from the cache. Does nothing if the key is not in the cache.
	Delete(ctx context.Context, key string, opts ...client.CallOption) error

	// Increment increments a real number, and returns error if the value is not real
	Increment(ctx context.Context, key string, data uint64, opts ...client.CallOption) (uint64, error)

	// Decrement decrements a real number, and returns error if the value is not real
	Decrement(ctx context.Context, key string, data uint64, opts ...client.CallOption) (uint64, error)

	// Flush seletes all items from the cache.
	Flush(ctx context.Context, opts ...client.CallOption) error
}

type cache struct {
	srv interface{}
}

func (c *cache) Get(ctx context.Context, key string, value interface{}, opts ...client.CallOption) error {
	fk := reflect.ValueOf(c.srv).MethodByName("Get")
	getReq := reflect.New(fk.Type().In(1).Elem())
	getReq.Elem().FieldByName("Key").SetString(key)

	optsv := []reflect.Value{
		reflect.ValueOf(ctx),
		getReq,
	}
	for _, o := range opts {
		optsv = append(optsv, reflect.ValueOf(o))
	}
	out := fk.Call(optsv)
	if e, ok := out[1].Interface().(error); ok && e != nil {
		return e
	}
	v, _ := out[0].Elem().FieldByName("Value").Interface().([]byte)
	encode.MustDeserialize(v, value)
	return nil
}

func (c *cache) Set(ctx context.Context, key string, value interface{}, expire time.Duration, opts ...client.CallOption) error {
	fk := reflect.ValueOf(c.srv).MethodByName("Set")
	setReq := reflect.New(fk.Type().In(1).Elem())
	setReq.Elem().FieldByName("Key").SetString(key)
	setReq.Elem().FieldByName("Value").SetBytes(encode.MustSerialize(value))
	setReq.Elem().FieldByName("Expire").SetInt(int64(expire))
	optsv := []reflect.Value{
		reflect.ValueOf(ctx),
		setReq,
	}
	for _, o := range opts {
		optsv = append(optsv, reflect.ValueOf(o))
	}
	out := fk.Call(optsv)
	if e, ok := out[1].Interface().(error); ok && e != nil {
		return e
	}
	return nil
}

func (c *cache) Add(ctx context.Context, key string, value interface{}, expire time.Duration, opts ...client.CallOption) error {
	fk := reflect.ValueOf(c.srv).MethodByName("Add")
	addReq := reflect.New(fk.Type().In(1).Elem())
	addReq.Elem().FieldByName("Key").SetString(key)
	addReq.Elem().FieldByName("Value").SetBytes(encode.MustSerialize(value))
	addReq.Elem().FieldByName("Expire").SetInt(int64(expire))

	optsv := []reflect.Value{
		reflect.ValueOf(ctx),
		addReq,
	}
	for _, o := range opts {
		optsv = append(optsv, reflect.ValueOf(o))
	}
	out := fk.Call(optsv)
	if e, ok := out[1].Interface().(error); ok && e != nil {
		return e
	}
	return nil
}

func (c *cache) Replace(ctx context.Context, key string, data interface{}, expire time.Duration, opts ...client.CallOption) error {
	fk := reflect.ValueOf(c.srv).MethodByName("Replace")
	replaceReq := reflect.New(fk.Type().In(1).Elem())
	replaceReq.Elem().FieldByName("Key").SetString(key)
	replaceReq.Elem().FieldByName("Value").SetBytes(encode.MustSerialize(data))
	replaceReq.Elem().FieldByName("Expire").SetInt(int64(expire))

	optsv := []reflect.Value{
		reflect.ValueOf(ctx),
		replaceReq,
	}
	for _, o := range opts {
		optsv = append(optsv, reflect.ValueOf(o))
	}
	out := fk.Call(optsv)
	if e, ok := out[1].Interface().(error); ok && e != nil {
		return e
	}
	return nil
}

func (c *cache) Delete(ctx context.Context, key string, opts ...client.CallOption) error {
	fk := reflect.ValueOf(c.srv).MethodByName("Delete")
	deleteReq := reflect.New(fk.Type().In(1).Elem())
	deleteReq.Elem().FieldByName("Key").SetString(key)

	optsv := []reflect.Value{
		reflect.ValueOf(ctx),
		deleteReq,
	}
	for _, o := range opts {
		optsv = append(optsv, reflect.ValueOf(o))
	}
	out := fk.Call(optsv)
	if e, ok := out[1].Interface().(error); ok && e != nil {
		return e
	}
	return nil
}

func (c *cache) Increment(ctx context.Context, key string, data uint64, opts ...client.CallOption) (uint64, error) {
	fk := reflect.ValueOf(c.srv).MethodByName("Increment")
	incrementReq := reflect.New(fk.Type().In(1).Elem())
	incrementReq.Elem().FieldByName("Key").SetString(key)
	incrementReq.Elem().FieldByName("Value").SetInt(int64(data))

	optsv := []reflect.Value{
		reflect.ValueOf(ctx),
		incrementReq,
	}
	for _, o := range opts {
		optsv = append(optsv, reflect.ValueOf(o))
	}
	out := fk.Call(optsv)
	if e, ok := out[1].Interface().(error); ok && e != nil {
		return 0, e
	}
	v, _ := out[0].Elem().FieldByName("Value").Interface().(int64)
	return uint64(v), nil
}

func (c *cache) Decrement(ctx context.Context, key string, data uint64, opts ...client.CallOption) (uint64, error) {
	fk := reflect.ValueOf(c.srv).MethodByName("Decrement")
	decrementReq := reflect.New(fk.Type().In(1).Elem())
	decrementReq.Elem().FieldByName("Key").SetString(key)
	decrementReq.Elem().FieldByName("Value").SetInt(int64(data))

	optsv := []reflect.Value{
		reflect.ValueOf(ctx),
		decrementReq,
	}
	for _, o := range opts {
		optsv = append(optsv, reflect.ValueOf(o))
	}
	out := fk.Call(optsv)
	if e, ok := out[1].Interface().(error); ok && e != nil {
		return 0, e
	}
	v, _ := out[0].Elem().FieldByName("Value").Interface().(int64)
	return uint64(v), nil
}

func (c *cache) Flush(ctx context.Context, opts ...client.CallOption) error {
	fk := reflect.ValueOf(c.srv).MethodByName("Flush")
	optsv := []reflect.Value{
		reflect.ValueOf(ctx),
	}
	for _, o := range opts {
		optsv = append(optsv, reflect.ValueOf(o))
	}
	out := fk.Call(optsv)
	if e, ok := out[1].Interface().(error); ok && e != nil {
		return e
	}
	return nil
}

func CacheService(srv interface{}) Cache {
	return &cache{srv}
}
