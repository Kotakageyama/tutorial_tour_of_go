package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}
func Sqrt(x float64) (float64, error) {

	z := 1.0
	old_z := 0.0
	if x < 0 {
		return -1, ErrNegativeSqrt(x)
	}
	for i := 0; i < 10; i++ {
		old_z = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		if math.Abs(z-old_z) < 0.0000001 {
			break
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
