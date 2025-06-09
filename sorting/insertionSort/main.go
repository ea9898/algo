package main

import "fmt"

func insertionSort(arr []int) []int {
	// Объявляем временную переменную для хранения текущего элемента
	var temp int

	// Внешний цикл проходит по всем элементам массива
	for idx := 0; idx < len(arr); idx++ {
		// Сохраняем текущий элемент во временную переменную
		temp = arr[idx]

		// Запоминаем текущий индекс для последующей вставки элемента
		currentIndex := idx

		// Внутренний цикл проходит по всем элементам слева от текущего
		for j := idx - 1; j >= 0; j-- {

			// Если текущий элемент меньше элемента слева
			if temp < arr[j] {
				// Сдвигаем элемент вправо, освобождая место для вставки
				arr[currentIndex] = arr[j]

				// Уменьшаем индекс для вставки
				currentIndex--
			}
		}

		// Вставляем сохраненный элемент в правильную позицию
		if currentIndex >= 0 {
			arr[currentIndex] = temp
		}
	}

	return arr
}

func main() {
	arr := []int{10, 1, 99, 0, 0, 0, 1}
	fmt.Println(arr)

	insertionSort(arr)
	fmt.Println(arr)
}
