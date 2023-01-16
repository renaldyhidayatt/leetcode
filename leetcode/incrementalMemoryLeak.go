package main

func memLeak(memory1 int, memory2 int) []int {
	i := 0
	for i = 0; ; i++ {
		var himem *int
		if memory1 >= memory2 {
			himem = &memory1
		} else {
			himem = &memory2
		}
		if *himem < i {
			break
		}
		*himem -= i
	}
	result := []int{i, memory1, memory2}
	return result
}
