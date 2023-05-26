package util

import (
	"errors"
	"math/rand"
	"time"
)

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

func Randomize[T comparable](collection []T) []T {
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(collection), func(i, j int) {
		collection[i], collection[j] = collection[j], collection[i]
	})
	return collection
}
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
