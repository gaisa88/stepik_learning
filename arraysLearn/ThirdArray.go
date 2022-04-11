package main

import (
	"fmt"
	"sort"
)

func main() {
	ThirdArray1()
	ThirdArray2()
	ThirdArray3()
	ThirdArray4()
}

func ThirdArray1() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	amax := array[0]
	for _, value := range array {
		if value >= amax {
			amax = value
		}
	}
	fmt.Println(amax)
	fmt.Println("End of first programm")

}
func ThirdArray2() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	for i := 0; i < 5; i++ {
		if array[i] > a {
			a = array[i]
		}
	}
	fmt.Println(a)
	fmt.Println("End of second programm")

}
func ThirdArray3() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	slc := array[:]
	sort.Ints(slc)
	fmt.Println(slc[4])
	fmt.Println("End of third programm")

}
func ThirdArray4() {
	array := [5]int{}
	var a int
	var max int = -999999999
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
		if a > max {
			max = a
		}
	}
	fmt.Println(max)
	fmt.Println("End of four programm")

}
