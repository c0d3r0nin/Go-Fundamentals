package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	// TODO: accumulate the sum of 1..n with a loop, and print that instead of the placeholder.
	//fmt.Println(0)

	sum := 0
	for index := 0; index <= n; index++ {
		sum += index

	}

	fmt.Println(sum)
}
