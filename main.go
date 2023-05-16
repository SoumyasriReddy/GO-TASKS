package main

import (
	"fmt"
	"practice/calculator1"
)

func main() {

	var a, b int
	fmt.Println("Enter a value:")
	fmt.Scan(&a)
	fmt.Println("Enter b value:")
	fmt.Scan(&b)
	fmt.Println("Add:", calculator.Add(a, b))
	fmt.Println("Sub:", calculator.Sub(a, b))
	fmt.Println("Mul:", calculator.Mul(a, b))
	fmt.Println("Div:", calculator.Div(a, b))
	fmt.Println("Mod:", calculator.Mod(a, b))

	var x, y, c, m, n float64
	fmt.Println("Enter x value:")
	fmt.Scan(&x)
	fmt.Println("Sqrt:", calculator.Sqrt(x))
	fmt.Println("Exp:", calculator.Exp(x))
	fmt.Println("Sin:", calculator.Sin(x))
	fmt.Println("Cos:", calculator.Cos(x))

	fmt.Println("Enter x value:")
	fmt.Scan(&x)
	fmt.Println("Enter y value:")
	fmt.Scan(&y)
	fmt.Println("Pow:", calculator.Pow(x, y))

	fmt.Println("Enter x value:")
	fmt.Scan(&x)
	fmt.Println("Enter y value:")
	fmt.Scan(&y)
	fmt.Println("Enter c value:")
	fmt.Scan(&c)
	fmt.Println("Enter m value:")
	fmt.Scan(&m)
	fmt.Println("Enter n value:")
	fmt.Scan(&n)
	fmt.Println("All:", calculator.All(x, y, c, m, n))

}
