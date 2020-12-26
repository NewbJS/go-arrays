package arrays


// Includes checks if an array or slice contains a specified v value.
func Includes(arr []interface{}, v interface{}) bool {
	found := 0

	for _, element := range arr {
		if element == v {
			found++
		}
	}

	if found >= 1 {
		return true
	}

	return false
}