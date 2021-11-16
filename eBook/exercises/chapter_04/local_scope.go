package main

import "fmt"

var a = "G" // global (package) scope

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
	a := "O" // new local variable a is declared
	print(a)
	fmt.Println()
}

// GOG
