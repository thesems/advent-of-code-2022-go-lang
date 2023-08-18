package utils

import (
	"log"
	"testing"
)

func TestIndexOf(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	if IndexOf(arr, 5) != 5 {
		log.Fatal("wrong!")
	}
}

func TestInsertAt(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	arr = InsertAt(arr, 5, -1)
	if arr[4] != 4 || arr[5] != -1 || arr[6] != 5 {
		log.Fatal("wrong!")
	}
}

func TestRemoveAt(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	arr = RemoveAt(arr, 1)
	if IndexOf(arr, 1) == 1 {
		log.Fatal("wrong!")
	}
}
