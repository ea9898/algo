package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 10

	fmt.Println(binarySearch(arr, target))
}

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	var mid int

	for low <= high { // итерируемся до тех пор, пока левая и правая часть не пересеклись
		mid = (low + high) / 2

		if arr[mid] == target { // если искомое число больше, чем средний элемент массива, отбрасываем левую часть диапазона
			return mid
		}

		if target > arr[mid] {
			low = mid + 1 // смещаем границу low вправо на 1 элемент
		} else {
			high = mid - 1 // смещаем границу high влево на 1 элемент
		}
	}

	return -1
}
