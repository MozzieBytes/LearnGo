package main

import (
	"reflect"
	"sync"
	"testing"

	"golang.org/x/tour/tree"
)

func Test_Walk_TreeMatch(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	got := make([]int, 0)
	t1, c := tree.New(1), make(chan int, 10)
	Walk(t1, c)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range c {
			got = append(got, v)
		}
	}()
	wg.Wait()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`Walk(tree.New(1), channel int) = %v, want %v`, got, want)
	}
}

func Test_Same_SameTree(t *testing.T) {
	rslt := Same(tree.New(1), tree.New(1))
	if !rslt {
		t.Fatalf(`Same(tree.New(1), tree.New(1)) = %v, want true`, rslt)
	}
}

func Test_Same_DiffTree(t *testing.T) {
	rslt := Same(tree.New(1), tree.New(2))
	if rslt {
		t.Fatalf(`Same(tree.New(1), tree.New(1)) = %v, want false`, rslt)
	}
}
