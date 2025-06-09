package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	first := mergeSort(arr[:len(arr)/2])
	second := mergeSort(arr[len(arr)/2:])

	return merge(first, second)
}

func merge(first []int, second []int) []int {
	var final []int
	i := 0
	j := 0

	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			final = append(final, first[i])
			i++
		} else {
			final = append(final, second[j])
			j++
		}
	}

	for i < len(first) {
		final = append(final, first[i])
		i++
	}

	for j < len(second) {
		final = append(final, second[j])
		j++
	}

	return final
}

func main() {
	unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	fmt.Println(unsorted)

	sorted := mergeSort(unsorted)
	fmt.Println(sorted)
}
