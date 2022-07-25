package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	old_z := 0.0
	// 最大10回計算するが、途中で差が0.0000001以下になれば終了
	for i := 0; i < 10; i++ {
		old_z = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		if math.Abs(z-old_z) < 0.0000001 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2), math.Sqrt(2))
}
