package utils

func deleteElement(slice [][2]int, index int) [][2]int {
	return append(slice[:index], slice[index+1:]...)
}
