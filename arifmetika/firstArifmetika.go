package main

import "fmt"

func main() {
	FirstAdd()
	SecondAdd()
	ThirdAdd()
}

func FirstAdd() {
	var a int
	fmt.Scan(&a)
	b := a * 2
	c := b + 100
	fmt.Println(c)
	fmt.Println("End of first programm")
}

func SecondAdd() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a*2 + 100)
	fmt.Println("End of second programm")

}

func ThirdAdd() {
	var (
		inputNum int
		outNum   int
	)

	fmt.Scan(&inputNum)

	outNum = (inputNum * 2) + 100

	fmt.Println(outNum)
	fmt.Println("End of third programm")

}
