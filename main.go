package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	words := strings.Fields(strings.TrimSpace(line))
	counts := make(map[string]int)

	// TODO: put every word of words into counts, then set these two numbers.

	for _, word := range words {
		counts[word]++
	}

	// Count words that appear only once.
	once := 0
	for _, frequency := range counts {
		if frequency == 1 {
			once++
		}
	}

	// Count the number of different words.
	different := len(counts)

	fmt.Println(different)
	fmt.Println(once)
}
