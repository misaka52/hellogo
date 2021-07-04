package basic

import "math"

func CalTriangle(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}
