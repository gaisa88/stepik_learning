package main

import "fmt"

func main() {
	LoopPersent1()
	LoopPersent2()
	LoopPersent3()
}
func LoopPersent1() {
	var x int
	var p int
	var y int
	var i = 0
	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)
	for x < y {
		x += x * p / 100
		i++
	}
	fmt.Println(i)
	fmt.Println("End of first programm")
}
func LoopPersent2() {
	var x, p, y, i int
	fmt.Scan(&x, &p, &y)
	for ; x < y; x += x * p / 100 {
		i++
	}
	fmt.Println(i)
	fmt.Println("End of second programm")
}
func LoopPersent3() {
	var x, p, y float32
	fmt.Scan(&x, &p, &y)
	years := 0
	for x <= y {
		x += x / 100 * p
		years++
	}
	fmt.Println(years)
	fmt.Println("End of third programm")

}
