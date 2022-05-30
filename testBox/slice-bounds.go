package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	a := s[:2]
	fmt.Println(a)
	a[1] = 1110
	fmt.Println(a, s)

	s = s[1:]
	fmt.Println(s)
}
