package math

func Lerp(v0, v1, t float64) float64 {
	return (1-t)*v0 + t*v1
}
