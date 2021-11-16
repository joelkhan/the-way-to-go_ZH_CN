package main

import "fmt"

func main() {
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // Note the explicit address-of operator to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f\n", x)
	y := Sum2(&array)
	fmt.Printf("The sum2 of the array is: %f\n", y)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range a { // can also with dereferencing *a to get back to the array
		sum += v
	}
	return
}

func Sum2(a *[3]float64) (sum float64) {
	for _, v := range *a { // can also with dereferencing *a to get back to the array
		sum += v
	}
	return
}

// Output: The sum of the array is: 24.600000
