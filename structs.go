package main

import "fmt"

type Person struct {
	name     string
	id       int
	age      int
	mobileno int64
	address  string
}

func main() {
	pointers()
	arrays()
	slices()
}

func pointers() {
	p := &Person{"Soumya", 123, 23, 8987654, "USA"}
	m := &Person{"Gokul", 234, 24, 3663722, "UK"}
	fmt.Println("Person1 name is:", (*p).name)
	fmt.Println("Person1 age is:", (*p).age)
	fmt.Println("Person2 name is:", (*m).name)
	fmt.Println("Person2 age is:", (*m).age)
}

func arrays() {
	members := [3]Person{
		{"Naveena", 654, 26, 23347, "US"},
		{"Sohan", 567, 34, 234567, "India"},
		{"Shanmukh", 897, 25, 987543, "Australia"},
	}

	slice := members[:2]
	fmt.Println("Array :", members)
	fmt.Println("slice :", slice)
	fmt.Println("length and capacity of slice:", len(slice), cap(slice))
}

func slices() {

	s := []string{"Soumya", "Naveena", "vedav", "surya", "Phani"}

	s = s[1:4]
	fmt.Println(s) // It gives the values of index from 1 to 3

	s = s[:3]
	fmt.Println(s) // It gives values of index 0 to 2

	fmt.Println(s, len(s), cap(s))
}
