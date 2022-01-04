package sorting

func BubbleSortInt(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	for i := len(data); i > 0; i-- {
		for j := 1; j < i; j++ {
			if data[j-1] > data[j] {
				data[j-1], data[j] = data[j], data[j-1]
			}
		}
	}

	return data
}
