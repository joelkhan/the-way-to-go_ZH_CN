package main

import (
	"fmt"
)

var sl []byte = make([]byte, 10, 100) // init sl
var start = 0

var data1 = [7]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
var data2 = [3]byte{'E', 'O', 'F'}

func main() {
	// Ex7.5
	fmt.Println("Ex7.5")
	fmt.Printf("sl len: %d, cap: %d\n", len(sl), cap(sl))
	fmt.Printf("first pos available is: %d\n", start)

	// 1st append
	fmt.Println("1st...")
	fmt.Println("view sl BEFORE Append()")
	viewSlice(sl)
	fmt.Printf("curr start pos: %d\n", start)
	sl = Append(sl, data1[:])
	fmt.Println("view sl AFTER Append()")
	viewSlice(sl)
	fmt.Printf("curr start pos: %d\n", start)

	// 2nd append
	fmt.Println("2nd...")
	sl = Append(sl, data2[:])
	fmt.Println("view sl AFTER Append()")
	viewSlice(sl)
	fmt.Printf("curr start pos: %d\n", start)

	// Ex7.6
	fmt.Println("Ex7.6")
}

func Append(slice, data []byte) []byte {
	// for _, val := range data {
	// 	fmt.Printf("%c ", val)
	// }
	// fmt.Println()

	avai := len(slice) - start
	if avai >= len(data) {
		for i := 0; i < len(data); i++ {
			slice[start+i] = data[i]
		}
		start += len(data)
	} else {
		fmt.Println("No enough spaces.")
	}
	return slice
}

func viewSlice(sl []byte) {
	for i := 0; i < len(sl); i++ {
		fmt.Printf("%c(%d) ", sl[i], sl[i])
	}
	fmt.Println()
}
