package arrays


// RemoveIndex removes an element that is in the specified index position of an array or slice.
func RemoveIndex(s []interface{}, index int) []interface{} {
	return append(s[:index], s[index+1:]...)
}
