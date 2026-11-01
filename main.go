package main

import "fmt"

type Point struct {
	X, Y int
}

// Receiver function aka Method
func (p Point) Squared_Distance_Between_Two_Points(other Point) int {
	dx := p.X - other.X
	dy := p.Y - other.Y
	return dx*dx + dy*dy
}

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1)
	fmt.Scan(&y1)
	fmt.Scan(&x2)
	fmt.Scan(&y2)
	pOne := Point{X: x1, Y: y1}
	pTwo := Point{X: x2, Y: y2}
	fmt.Println(pOne.Squared_Distance_Between_Two_Points(pTwo))
}
