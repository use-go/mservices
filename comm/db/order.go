package db

type order interface {
	GetOrderType() int32
	GetOrderCol() string
}
