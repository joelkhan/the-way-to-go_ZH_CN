package main

import "fmt"

func main() {
	// ex7.9
	ex79()

	// ex7.10 see filter_slice.go
	// ex7.11 参考7.6.7 append 函数常见操作 之 7
	// ex7.12 参考7.6.7 append 函数常见操作 之 4
}

func ex79() {
	s := []int{1, 2, 3, 4, 5}
	unitLen := len(s)
	factor := 3
	fmt.Printf("slice len: %d, cap: %d, factor: %d\n", len(s), cap(s), factor)
	for i := 0; i < factor; i++ {
		//slice1 := make([]int, unitLen, unitLen)
		//fmt.Printf("slice111 len: %d, cap: %d\n", len(slice1), cap(slice1))
		s = append(s, make([]int, unitLen)...)
		//s = append(s, slice1...)
		fmt.Printf("slice len: %d, cap: %d\n", len(s), cap(s))
	}
	//fmt.Printf("slice len: %d, cap: %d\n", len(s), cap(s))
}
