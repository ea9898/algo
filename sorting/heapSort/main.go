package main

import "fmt"

// heapSort выполняет сортировку массива методом пирамидальной сортировки (сортировки кучей)
func heapSort(arr []int) []int {
	n := len(arr)

	// Построение кучи (перегруппировка массива)
	// Начинаем с последнего родительского узла и идем к корню
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Извлечение элементов из кучи один за другим
	for i := n - 1; i > 0; i-- {
		// Перемещаем текущий корень (максимальный элемент) в конец
		arr[0], arr[i] = arr[i], arr[0]

		// Вызываем heapify на уменьшенной куче
		heapify(arr, i, 0)
	}

	return arr
}

// heapify преобразует поддерево с корнем i в max-кучу
// n - размер кучи, i - индекс корня поддерева
func heapify(arr []int, n, i int) {
	largest := i     // Инициализируем наибольший элемент как корень
	left := 2*i + 1  // Левый потомок
	right := 2*i + 2 // Правый потомок

	// Если левый потомок больше корня
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// Если правый потомок больше, чем самый большой элемент на данный момент
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// Если наибольший элемент не корень
	if largest != i {
		// Меняем местами корень и наибольший элемент
		arr[i], arr[largest] = arr[largest], arr[i]

		// Рекурсивно преобразуем затронутое поддерево
		heapify(arr, n, largest)
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println(arr)

	sortedArr := heapSort(arr)
	fmt.Println(sortedArr)
}