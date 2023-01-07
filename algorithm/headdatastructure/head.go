package headdatastructure

import "fmt"

func heapify(arr []int, n int, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && arr[i] < arr[l] {
		largest = l
	}

	if r < n && arr[largest] < arr[r] {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func insert(array []int, newNum int) {
	size := len(array)
	if size == 0 {
		array = append(array, newNum)
	} else {
		array = append(array, newNum)
		for i := (size / 2) - 1; i >= 0; i-- {
			heapify(array, size, i)
		}
	}
}

func deleteNode(array []int, num int) {
	size := len(array)
	i := 0
	for i = 0; i < size; i++ {
		if num == array[i] {
			break
		}
	}

	array[i], array[size-1] = array[size-1], array[i]

	array = append(array[:i], array[i+1:]...)

	for i := (len(array) / 2) - 1; i >= 0; i-- {
		heapify(array, len(array), i)
	}
}

func HeadDataMain() {
	arr := []int{}

	insert(arr, 3)
	insert(arr, 4)
	insert(arr, 9)
	insert(arr, 5)
	insert(arr, 2)

	fmt.Println("Max-Heap array:", arr)

	deleteNode(arr, 4)
	fmt.Println("After deleting an element:", arr)
}
