package strain

func Keep[T any](items []T, pred func(T) bool) []T {
	var kept []T

	for _, item := range items {
		if pred(item) {
			kept = append(kept, item)
		}
	}

	return kept
}

func Discard[T any](items []T, pred func(T) bool) []T {
	var kept []T

	for _, item := range items {
		if !pred(item) {
			kept = append(kept, item)
		}
	}

	return kept
}
