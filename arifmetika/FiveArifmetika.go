package main

import "fmt"

func main() {
	FirstDecNumber()
	SecondDecNumber()
	ThirdDecNumber()
}
func FirstDecNumber() {
	var a int
	var b int
	fmt.Scan(&a)
	b = a / 10 % 10
	fmt.Println(b)
	fmt.Println("End of first programm")
}
func SecondDecNumber() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a % 100 / 10)
	fmt.Println("End of second programm")

}

func ThirdDecNumber() {
	var a string
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	//fmt.Print(a)
	fmt.Println(string(a[len(a)-2]))
	fmt.Println("End of third programm")

}
