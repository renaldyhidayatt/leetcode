package main

func arraySign(nums []int) int {
	negative := 0
	for index, element := range nums {
		if (element==0) return 0;
		if (element < 0) negative++;
	}
}
