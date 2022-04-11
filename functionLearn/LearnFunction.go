package main

import "fmt"

func main() {
	x := minimumFromFour()
	fmt.Print(x)
}
func minimumFromFour() int {
	var array [4]int
	var a int
	for i := 0; i < len(array); i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	amax := array[0]
	for _, value := range array {
		if value <= amax {
			amax = value
		}
	}
	return amax
}
