package bitwiseorofsubarrays

func subarrayBitwiseORs(A []int) int {
	res, t := map[int]bool{}, map[int]bool{}
	for _, num := range A {
		r := map[int]bool{}
		r[num] = true
		for n := range t {
			r[(num | n)] = true
		}
		t = r
		for n := range t {
			res[n] = true
		}
	}
	return len(res)
}
