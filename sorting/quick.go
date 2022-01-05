package sorting

func QuickSortInt(data []int) []int {
	return quickSort(data, 0, len(data)-1)
}

func quickSort(data []int, low, high int) []int {
	if low > high {
		return data
	}

	// 1. Начальное партиционирование массива
	var p int
	data, p = partition(data, low, high)

	// 2. Сортировка партиции со значениями меньше опорного
	data = quickSort(data, low, p-1)

	// 3. Сортировка партиции со значениями больше опорного
	data = quickSort(data, p+1, high)

	return data
}

// Опорным элементом считаем последний элемент в патриции
func partition(data []int, low, high int) ([]int, int) {
	pivot := data[high]
	point := low

	for j := point; j < high; j++ {
		if data[j] < pivot {
			data[j], data[point] = data[point], data[j]
			point++
		}
	}

	data[point], data[high] = data[high], data[point]

	return data, point
}
