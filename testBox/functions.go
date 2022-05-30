package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func pow(x, y float64) float64 {
	return math.Pow(x, y)
}
func main() {
	fmt.Println(add(42, 13))
	fmt.Println(pow(4, 2.0))
}
