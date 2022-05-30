package main

import "fmt"

func main() {
	q := []int{1, 2, 3, 4}
	r := []bool{true, false, true, false, false, false, false, false, false, false}
	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, false},
	}
	fmt.Println(q, r, s)
}
