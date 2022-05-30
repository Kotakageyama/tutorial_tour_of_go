package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}
func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		value1, ok1 := <-ch1
		value2, ok2 := <-ch2
		if value1 != value2 {
			return false
		}
		if ok1 != ok2 {
			return false
		}
		if !ok1 && !ok2 {
			break
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for c := range ch {
		fmt.Println(c)
	}
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("same")
	} else {
		fmt.Println("false")
	}
	if Same(tree.New(1), tree.New(2)) {
		fmt.Println("same")
	} else {
		fmt.Println("false")
	}
}
