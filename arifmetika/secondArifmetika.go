package main

import (
	"fmt"
	"math"
)

func main() {
	FirstMulty()
	SecondMulty()
	ThirdMulty()
}
func FirstMulty() {

	var a int
	var b int
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	fmt.Scan(&b) // считаем переменную 'b' с консоли

	a *= a
	b *= b
	c := a + b
	fmt.Println(c)
	fmt.Println("End of first programm")

}

func SecondMulty() {

	var a, b int
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	fmt.Scan(&b) // считаем переменную 'b' с консоли

	fmt.Println((a * a) + (b * b))
	fmt.Println("End of second programm")

}

func ThirdMulty() {
	var number1, number2 float64
	fmt.Scan(&number1, &number2)
	fmt.Println(math.Pow(number1, 2) + math.Pow(number2, 2))
	fmt.Println("End of third programm")

}
