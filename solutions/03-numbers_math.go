package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	quotient := a / b

	remainder := a % b

	fmt.Println(quotient)
	fmt.Println(remainder)

}
