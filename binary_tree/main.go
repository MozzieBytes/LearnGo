package main

import (
	"fmt"
	"reflect"
	"sync"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t.Left != nil {
			walker(t.Left)
		}
		ch <- t.Value
		if t.Right != nil {
			walker(t.Right)
		}
	}
	walker(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	var sliceTree func(*tree.Tree, chan []int)
	sliceTree = func(t *tree.Tree, ch chan []int) {
		var wg sync.WaitGroup
		wg.Add(1)
		s := make([]int, 10)
		go func() {
			defer wg.Done()
			c := make(chan int, 10)
			Walk(t, c)
			for v := range c {
				s = append(s, v)
			}
		}()
		wg.Wait()
		ch <- s
	}

	var wg sync.WaitGroup
	wg.Wait()

	wg.Add(2)
	c1, c2 := make(chan []int), make(chan []int)
	go func() { defer wg.Done(); go sliceTree(t1, c1) }()
	go func() { defer wg.Done(); go sliceTree(t2, c2) }()
	wg.Wait()

	return reflect.DeepEqual(<-c1, <-c2)
}

func main() {
	c := make(chan int, 10)
	go Walk(tree.New(2), c)

	i := 1
	for v := range c {
		fmt.Printf("Value %d: %d\n", i, v)
		i++
	}
}
