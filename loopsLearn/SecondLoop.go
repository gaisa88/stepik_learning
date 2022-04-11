package main

import "fmt"

func main() {
	NumbersAdd1()
	NumbersAdd2()

}
func NumbersAdd1() {
	var a, b, c int
	fmt.Scan(&a, &b)
	for ; a <= b; a++ {
		c += a
	}
	fmt.Println(c)
	fmt.Println("End of first programm")
}
func NumbersAdd2() {
	var a, b int
	fmt.Scan(&a, &b)
	s := 0
	for a <= b {
		s += a
		a++
	}
	fmt.Println(s)
	fmt.Println("End of second programm")

}
