// filter_slice.go
package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	s = Filter(s, even)
	fmt.Println(s)
}

// Filter returns a new slice holding only
// the elements of s that satisfy f()
func Filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, i := range s {
		if fn(i) {
			fmt.Printf("p len: %d, cap: %d\n", len(p), cap(p))
			p = append(p, i)
		}
	}
	return p
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

/* [0 2 4 6 8] */
