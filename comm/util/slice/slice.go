package slice

func InSlice(val string, arr ...string) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}
