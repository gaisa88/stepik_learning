package main

import (
	"fmt"
	"math"
)

func main() {
	FirstMultiply()
	SecondMultiply()
	ThirdMultiply()
}

func FirstMultiply() {
	var a int
	fmt.Scan(&a)
	a *= a
	fmt.Println(a)
	fmt.Println("End of first programm")
}
func SecondMultiply() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a * a)
	fmt.Println("End of second programm")

}
func ThirdMultiply() {
	var number float64
	fmt.Scan(&number)
	fmt.Println(math.Pow(number, 2))
	fmt.Println("End of third programm")

}
