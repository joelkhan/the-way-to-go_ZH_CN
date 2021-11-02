// leaving out the length 4 in [4] uses  slice instead of an array
// which is generally better for performance
package main

import "fmt"

func main() {
	// var a = [4]float32 {1.0,2.0,3.0,4.0}
	var a = []float32{1.0, 2.0, 3.0, 4.0}
	fmt.Printf("The sum of the array is: %f\n", Sum(a))
	var b = []int{1, 2, 3, 4, 5}
	sum, average := SumAndAverage(b)
	fmt.Printf("The sum of the array is: %d, and the average is: %f\n", sum, average)

	// 问题 7.5
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
	}
	for i := 0; i < len(items); i++ {
		fmt.Printf("%d ", items[i])
	}
	fmt.Println()

	for ix := range items {
		items[ix] *= 2
	}
	for i := 0; i < len(items); i++ {
		fmt.Printf("%d ", items[i])
	}
	fmt.Println()

	// 问题 7.7
	var arr = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var asl = arr[5:7]
	fmt.Printf("%d %d\n", len(asl), cap(asl))
	asl = asl[0:4]
	fmt.Printf("%d %d\n", len(asl), cap(asl))
	asl = asl[1:1]
	fmt.Printf("%d %d\n", len(asl), cap(asl))
	asl = asl[1:2]
	fmt.Printf("%d %d\n", len(asl), cap(asl))
}

/*
func Sum(a [4]float32) (sum float32) {
	for _, item := range a {
		sum += item
	}
	return
}
*/

func Sum(a []float32) (sum float32) {
	for _, item := range a {
		sum += item
	}
	return
}

func SumAndAverage(a []int) (int, float32) {
	sum := 0
	for _, item := range a {
		sum += item
	}
	return sum, float32(sum / len(a))
}
