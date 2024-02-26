package dynamic

import "fmt"

var errCatalan = fmt.Errorf("can't have a negative n-th catalan number")

func NthCatalanNumber(n int) (int64, error) {
	if n < 0 {
		return 0, errCatalan
	}

	var catalanNumberList []int64
	catalanNumberList = append(catalanNumberList, 1)

	for i := 1; i <= n; i++ {
		catalanNumberList = append(catalanNumberList, 0)

		for j := 0; j < i; j++ {
			catalanNumberList[i] += catalanNumberList[j] * catalanNumberList[i-j-1]
		}
	}

	return catalanNumberList[n], nil
}
