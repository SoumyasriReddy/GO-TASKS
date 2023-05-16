package calculator

import "math"

func Add(a, b int) (sum int) {
	sum = a + b
	return sum
}

func Sub(a, b int) (sum int) {
	sum = a - b
	return sum
}

func Mul(a, b int) (sum int) {
	sum = a * b
	return sum
}

func Div(a, b int) (sum int) {
	sum = a / b
	return sum
}

func Mod(a, b int) (sum int) {
	sum = a % b
	return sum
}

func Sqrt(x float64) (sum float64) {
	sum = math.Sqrt(x)
	return sum
}

func Exp(x float64) (sum float64) {
	sum = math.Exp(x)
	return sum
}

func Sin(x float64) (sum float64) {
	sum = math.Sin(x)
	return sum
}

func Cos(x float64) (sum float64) {
	sum = math.Cos(x)
	return sum
}

func Pow(x, y float64) (sum float64) {
	sum = math.Pow(x, y)
	return sum
}

func All(x, y, c, m, n float64) (sum float64) {
	sum = x + y - c*m/n
	return sum
}
