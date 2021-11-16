package main

import (
	"fmt"
)

const CAPACITY_UNIT = 20

var data1 = [7]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
var data2 = [3]byte{'E', 'O', 'F'}
var data3 = [26]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g',
	'h', 'i', 'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z'}

func main() {
	sl := make([]byte, 3, CAPACITY_UNIT)
	// init slice
	initSlice(sl)
	// view slice
	viewSlice(sl)

	fmt.Println("1st append...")
	sl = Append(sl, data1[:])
	viewSlice(sl)

	fmt.Println("2nd append...")
	sl = Append(sl, data2[:])
	viewSlice(sl)

	fmt.Println("3rd append...")
	sl = Append(sl, data1[:])
	viewSlice(sl)

	fmt.Println("4th append...")
	sl = Append(sl, data3[:])
	viewSlice(sl)
}

func initSlice(sl []byte) {
	for i := 0; i < len(sl); i++ {
		//fmt.Printf("%c(%d) ", sl[i], sl[i])
		sl[i] = byte(int('X') + i)
	}
}

func viewSlice(sl []byte) {
	fmt.Printf("sl len: %d, cap: %d\n", len(sl), cap(sl))
	for i := 0; i < len(sl); i++ {
		fmt.Printf("%c(%d) ", sl[i], sl[i])
	}
	fmt.Println()
}

func Append(slice, data []byte) []byte {
	avai := cap(slice) - len(slice)

	if avai < len(data) {
		fmt.Println("No enough spaces.") // enlarge slice capacity
		times := len(data)/CAPACITY_UNIT + 1
		fmt.Printf("need to enlarge: %d\n", times)
		biggerSl := make([]byte, len(slice), cap(slice)+CAPACITY_UNIT*times)
		for i := 0; i < len(biggerSl); i++ {
			biggerSl[i] = slice[i]
		}
		slice = biggerSl
	}
	oldLen := len(slice)
	slice = slice[:oldLen+len(data)]
	for i := 0; i < len(data); i++ {
		slice[oldLen+i] = data[i]
	}

	// if avai >= len(data) {
	// 	oldLen := len(slice)
	// 	slice = slice[:oldLen+len(data)]
	// 	for i := 0; i < len(data); i++ {
	// 		slice[oldLen+i] = data[i]
	// 	}
	// } else {
	// 	fmt.Println("No enough spaces.") // enlarge slice capacity
	// 	times := len(data)/CAPACITY_UNIT + 1
	// 	fmt.Printf("need to enlarge: %d\n", times)
	// 	biggerSl := make([]byte, len(slice), cap(slice)+CAPACITY_UNIT*times)
	// 	for i := 0; i < len(biggerSl); i++ {
	// 		biggerSl[i] = slice[i]
	// 	}
	// 	slice = biggerSl
	// }
	return slice
}
