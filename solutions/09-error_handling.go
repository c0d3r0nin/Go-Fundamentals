package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")

	// TODO: parse `line` with strconv.Atoi, which hands back two values:
	// the number it read and an error. Work out which of the two tells you
	// whether the parse actually worked, then print the success line or the
	// failure word the task asks for. Replace the placeholder below.

	if num, err := strconv.Atoi(line); err != nil {
		fmt.Println("bad")
	} else {
		fmt.Println("ok", num)
	}

	// if error != nil {
	// 	fmt.Printf("%v", error)
	// }
	// fmt.Printf("ok %v", str)

}
