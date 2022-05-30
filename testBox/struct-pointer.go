package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	var v Vertex
	v = Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
	fmt.Println(1e11)
	fmt.Println(1e9)
}
