package matrix

import (
	"context"
	"errors"
	"sync"
)

func (m1 Matrix[T]) Add(m2 Matrix[T]) (Matrix[T], error) {
	if !m1.MatchDimensions(m2) {
		return Matrix[T]{}, errors.New("matrices are not compatible for addition")
	}

	var zeroVal T
	result := New[T](m1.Rows(), m1.Columns(), zeroVal)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var wg sync.WaitGroup
	errCh := make(chan error, 1)

	for i := 0; i < m1.rows; i++ {
		i := i
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < m1.columns; j++ {
				select {
				case <-ctx.Done():
					return
				default:
				}

				sum := m1.elements[i][j] + m2.elements[i][j]

				err := result.Set(i, j, sum)

				if err != nil {
					cancel()

					select {
					case errCh <- err:
					default:

					}

					return
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	if err := <-errCh; err != nil {
		return Matrix[T]{}, err
	}

	return result, nil
}
