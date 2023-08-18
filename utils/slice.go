package utils

import "log"

func DeleteElement[K any](slice []K, index int) []K {
	return append(slice[:index], slice[index+1:]...)
}

func IndexOf(array []int, value int) int {
	for idx, num := range array {
		if num == value {
			return idx
		}
	}
	log.Fatalf("could not find value %d\n", value)
	return 0
}

func RemoveAt[K any](array []K, index int) []K {
	return append(array[:index], array[index+1:]...)
}
