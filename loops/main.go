package main

import "fmt"

func main() {
	var j int
	fmt.Println("Enter a value:")
	fmt.Scan(&j)
	var n, ch string
	fmt.Println("choice:")
	fmt.Scan(&n)
	ch = n
	switch ch {
	case "even":
		even(j)
	case "prime":
		if prime(j) {
			fmt.Println(j, "is a prime number")
		} else {
			fmt.Println(j, "is not a prime number")
		}
	case "fibonacci":
		fmt.Println(fibonacci(j))
	case "palindrome":
		palindrome(j)
	default:
		fmt.Println("Invalid")
	}

}

func even(j int) {
	if j%2 == 0 {
		fmt.Println(j, "is even number")
	} else {
		fmt.Println(j, "is odd number")
	}
}

func prime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func fibonacci(j int) int {
	if j <= 1 {
		return j
	}
	return fibonacci(j-1) + fibonacci(j-2)
}

func palindrome(j int) {
	var rev, rem int

	temp := j

	for temp != 0 {
		rem = temp % 10
		rev = rev*10 + rem
		temp /= 10
	}

	if j == rev {
		fmt.Println("palindrome", j)
	} else {
		fmt.Println("not a palindrome", j)
	}
}
