package dynamic

import "fmt"

var ErrInvalidPosition = fmt.Errorf("invalid position in subset")
var ErrNegativeSum = fmt.Errorf("negative sum is not allowed")

func IsSubsetSum(array []int, sum int) (bool, error) {
	if sum < 0 {
		return false, ErrNegativeSum
	}

	arraySize := len(array)
	subset := make([][]bool, arraySize)

	for i := 0; i <= arraySize; i++ {
		subset[i] = make([]bool, sum+1)
	}

	for i := 0; i <= arraySize; i++ {
		subset[i][0] = true
	}

	for i := 1; i <= sum; i++ {
		subset[0][i] = false
	}

	for i := 1; i <= arraySize; i++ {
		for j := 1; j <= sum; j++ {
			if array[i-1] > j {
				subset[i][j] = subset[i-1][j]
			}

			if array[i-1] <= j {
				if j-array[i-1] < 0 || j-array[i-1] > sum {

					return false, ErrInvalidPosition
				}

				subset[i][j] = subset[i-1][j] || subset[i-1][j-array[i-1]]
			}
		}
	}

	return subset[arraySize][sum], nil
}
