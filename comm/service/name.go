package service

var name string

func SetName(n string) string {
	name = n
	return n
}

func GetName() string {
	return name
}
