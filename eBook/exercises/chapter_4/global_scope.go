package main

import "fmt"

var a = "G" // global scope

func main() {
	n()
	m()
	n()
}

func n() {
	print(a)
	fmt.Println()
}

func m() {
	a = "O" // simple assignment: global a gets a new value
	print(a)
	fmt.Println()
}

// GOO
