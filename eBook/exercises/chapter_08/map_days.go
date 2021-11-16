// map_days.go
package main

import (
	"fmt"
)

var Days = map[int]string{1: "monday",
	2: "tuesday",
	3: "wednesday",
	4: "thursday",
	5: "friday",
	6: "saturday",
	7: "sunday"}

func main() {
	fmt.Println(Days)
	// fmt.Printf("%v", Days)
	flagHolliday := false
	for k, v := range Days {
		if v == "thursday" || v == "holliday" {
			fmt.Println(v, " is the ", k, "th day in the week")
			if v == "holliday" {
				flagHolliday = true
			}
		}
	}
	if !flagHolliday {
		fmt.Println("holliday is not a day!")
	}

	// 问题 8.1
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}
}

/* Output:
thursday  is the  4 th day in the week
holliday is not a day!
*/
