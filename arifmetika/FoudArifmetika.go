package main

import "fmt"

func main() {
	FirstEndNumber()
	SecondEndNumber()
	ThirdEndNumber()
}

func FirstEndNumber() {
	var a int
	var b int
	fmt.Scan(&a)
	b = a % 10
	fmt.Println(b)
	fmt.Println("End of first programm")

}
func SecondEndNumber() {
	var n int
	fmt.Scan(&n)
	fmt.Println(n % 10)
	fmt.Println("End of second programm")

}
func ThirdEndNumber() {
	var a string
	fmt.Scan(&a)
	fmt.Println(string(a[len(a)-1]))
	fmt.Println("End of third programm")

}
