package utils

func DeleteElement[K any](slice []K, index int) []K {
	return append(slice[:index], slice[index+1:]...)
}
