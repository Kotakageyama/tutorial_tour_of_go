package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for i, _ := range image {
		image[i] = make([]uint8, dx)
	}
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			image[i][j] = uint8(i ^ j)
		}
	}
	return image
}

func main() {
	pic.Show(Pic)
	fmt.Printf("ok end\n")
}
