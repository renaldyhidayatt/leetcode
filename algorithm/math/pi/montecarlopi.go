package pi

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func MonteCarloPi(randomPoints int) float64 {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	inside := 0

	for i := 0; i < randomPoints; i++ {
		x := rnd.Float64()
		y := rnd.Float64()

		if x*x+y*y <= 1 {
			inside += 1
		}
	}

	pi := float64(inside) / float64(randomPoints) * 4

	return pi
}

func MonteCarloPiConcurrent(n int) (float64, error) {
	numCpu := runtime.GOMAXPROCS(0)

	c := make(chan int, numCpu)

	pointToDraw, err := splitInt(n, numCpu)

	if err != nil {
		return 0, err
	}

	for _, p := range pointToDraw {
		go drawPoints(p, c)
	}

	inside := 0

	for i := 0; i < numCpu; i++ {
		inside += <-c
	}

	return float64(inside) / float64(n) * 4, nil
}

func drawPoints(n int, c chan<- int) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	inside := 0

	for i := 0; i < n; i++ {
		x, y := rnd.Float64(), rnd.Float64()

		if x*x+y*y <= 1 {
			inside++
		}
	}
	c <- inside
}

func splitInt(x int, n int) ([]int, error) {
	if x < n {
		return nil, fmt.Errorf("x must be < n - given values are x=%d, n=%d", x, n)
	}

	split := make([]int, n)

	if x%n == 0 {
		for i := 0; i < n; i++ {
			split[i] = x / n
		}
	} else {
		limit := x % n

		for i := 0; i < limit; i++ {
			split[i] = x/n + 1
		}

		for i := limit; i < n; i++ {
			split[i] = x / n
		}
	}

	return split, nil
}
