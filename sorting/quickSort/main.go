package main

import "fmt"

func partition(arr []int, startIndex, endIndex int) int {
	pivot := arr[(startIndex + endIndex) / 2]
	i := startIndex
	j := endIndex

	for {
		// Эта проверка используется для завершения цикла for.
		// Все элементы слева от i меньше или равны пивоту,
		// а все элементы справа от j больше или равны пивоту.
		if i >= j {
			// Возвращается индекс i, который указывает на начало правой части массива.
			// Это используется для разделения массива на две части для дальнейшей рекурсивной сортировки.
			return i
		}

		for arr[i] < pivot {
			i++
		}
		
		for arr[j] > pivot {
			j--
		}

		arr[i], arr[j] = arr[j], arr[i]
	}
}

func quickSort(arr []int, low, high int) {
	// Если low меньше high, это означает, что в массиве есть более одного элемента, и его нужно сортировать.
	// Если low не меньше high, это означает, что массив либо пуст, либо состоит из одного элемента, и сортировка не требуется.
	if low < high {
		// Если условие выполняется, функция quickSort рекурсивно вызывается для левой и правой частей массива.
		pivotIndex := partition(arr, low, high)

		// Этот вызов сортирует подмассив, который находится слева от пивота.
		// Чтобы избежать повторной сортировки этого элемента, мы вызываем quickSort для подмассива,
		// который заканчивается на pivotIndex-1
		quickSort(arr, low, pivotIndex-1)

		// Этот вызов сортирует подмассив, который находится справа от пивота.
		// Аналогично, элемент на позиции pivotIndex уже отсортирован, поэтому мы
		// начинаем сортировку правой части массива с pivotIndex+1
		quickSort(arr, pivotIndex+1, high)
	}
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Исходный массив:", arr)

	quickSort(arr, 0, len(arr)-1)

	fmt.Println("Отсортированный массив:", arr)
}