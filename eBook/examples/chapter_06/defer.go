package main

import "fmt"

func main() {
	Function1()
	a()
	f()
}

func Function1() {
	fmt.Printf("In Function1 at the top\n")
	//defer Function2()
	Function2() // try without defer
	fmt.Printf("In Function1 at the bottom!\n")
}

func Function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!\n")
}


func a() {
	i := 0
	defer fmt.Println(i)
	i++
	fmt.Println(i)
	return
}

func f() {
	defer fmt.Println()
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

