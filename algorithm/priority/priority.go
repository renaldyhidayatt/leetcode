package priority

import "fmt"

func Heapify(arr []int, n, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && arr[i] < arr[l] {
		largest = l
	}

	if r < n && arr[largest] < arr[r] {
		largest = r
	}

	// Swap and continue heapifying if root is not largest
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		Heapify(arr, n, largest)
	}
}

func Insert(array []int, newNum int) []int {
	size := len(array)
	if size == 0 {
		array = append(array, newNum)
	} else {
		array = append(array, newNum)
		for i := (size / 2) - 1; i >= 0; i-- {
			Heapify(array, size, i)
		}

	}
	return array
}

func DeleteNode(array []int, num int) []int {
	size := len(array)
	i := 0
	for i = 0; i < size; i++ {
		if num == array[i] {
			break
		}
	}

	array[i], array[size-1] = array[size-1], array[i]

	array = array[:size-1]

	for i := (len(array) / 2) - 1; i >= 0; i-- {
		Heapify(array, len(array), i)
	}

	return array
}

func PriorityMain() {
	arr := []int{}

	arr = Insert(arr, 3)
	arr = Insert(arr, 4)
	arr = Insert(arr, 9)
	arr = Insert(arr, 5)
	arr = Insert(arr, 2)

	fmt.Println("Max-Heap array:", arr)

	arr = DeleteNode(arr, 4)
	fmt.Println("After deleting an element:", arr)
}
