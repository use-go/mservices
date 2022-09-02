package service

var name string

// SetName defined todo
func SetName(n string) string {
	name = n
	return n
}

// GetName defined todo
func GetName() string {
	return name
}
