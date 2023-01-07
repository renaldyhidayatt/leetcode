package main

import "fmt"

func checkPrime(n int) int {
	if n == 1 || n == 0 {
		return 0
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return 0
		}
	}

	return 1
}

func getPrime(n int) int {
	if n%2 == 0 {
		n = n + 1
	}

	for checkPrime(n) == 0 {
		n += 2
	}

	return n
}

func hashFunction(key int) int {
	capacity := getPrime(10)

	return key % capacity
}

func insertData(hashTable map[int][]int, key int, data int) {
	index := hashFunction(key)
	hashTable[index] = []int{key, data}
}

func removeData(hashTable map[int][]int, key int) {
	index := hashFunction(key)
	delete(hashTable, index)
}

func HashTableMain() {
	hashTable := make(map[int][]int)

	insertData(hashTable, 123, 10)
	insertData(hashTable, 432, 20)
	insertData(hashTable, 213, 30)
	insertData(hashTable, 654, 40)

	fmt.Println(hashTable)

	removeData(hashTable, 123)

	fmt.Println(hashTable)

}
