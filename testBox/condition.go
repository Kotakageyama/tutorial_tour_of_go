package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Printf("Good morning")
	case t.Hour() < 17:
		fmt.Printf("Good afternoon")
	default:
		fmt.Printf("Hello, world")
	}
}
