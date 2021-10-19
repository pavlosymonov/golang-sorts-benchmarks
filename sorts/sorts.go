package sorts

func BubbleSort(items []int) {
	for i := len(items); i > 0; i-- {
		for j := 1; j < i; j++ {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
		}
	}
}

func InsertionSort(items []int) {
	for i := 1; i < len(items); i++ {
		for j := i; j > 0; j-- {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
		}
	}
}

func SelectionSort(items []int) {
	n := len(items)

	for i := 0; i < n; i++ {
		minIdx := i

		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}
