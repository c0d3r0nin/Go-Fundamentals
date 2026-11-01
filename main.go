package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// TODO: choose the right label for n instead of always printing it.
	//fmt.Println(n)
	switch {
	case n%3 == 0 && n%5 == 0:
		fmt.Println("FizzBuzz")
	case n%3 == 0:
		fmt.Println("Fizz")
	case n%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println(n)
	}
}
