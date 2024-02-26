package matrix

import (
	"context"
	"errors"
	"sync"
)

func (m1 Matrix[T]) Multiply(m2 Matrix[T]) (Matrix[T], error) {
	if m1.Columns() != m2.Rows() {
		return Matrix[T]{}, errors.New("matrices cannot be multiplied: column count of the first matrix must match row count of the second matrix")
	}

	var zeroVal T
	result := New[T](m1.Rows(), m2.Columns(), zeroVal)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var wg sync.WaitGroup

	errCh := make(chan error, 1)

	for i := 0; i < m1.Rows(); i++ {
		for j := 0; j < m2.Columns(); j++ {
			i, j := i, j

			wg.Add(1)

			go func() {
				defer wg.Done()

				dotProduct := zeroVal

				for k := 0; k < m1.Columns(); k++ {
					select {
					case <-ctx.Done():
						return
					default:
					}
					val1, err := m1.Get(i, k)

					if err != nil {
						cancel()

						select {
						case errCh <- err:
						default:
						}

						return
					}

					val2, err := m2.Get(k, j)
					if err != nil {
						cancel()
						select {
						case errCh <- err:
						default:
						}
						return
					}
					dotProduct += val1 * val2
				}

				err := result.Set(i, j, dotProduct)

				if err != nil {
					cancel()
					select {
					case errCh <- err:
					default:

					}
					return
				}

			}()
		}
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
