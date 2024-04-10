package handlers

import "strconv"

func lenStr[T any](list []T) string {
	return toString(len(list))
}

func toString(i int) string {
	return strconv.Itoa(i)
}
