package util

import "errors"

func RemoveIndex[T comparable](collection []T, index int) []T {
	return append(collection[:index], collection[index+1:]...)
}
func IndexOf[T comparable](collection []T, el T) (int, error) {
	for i, x := range collection {
		if x == el {
			return i, nil
		}
	}
	return -1, errors.New("Element index not found")
}
