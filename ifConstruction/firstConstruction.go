package main

import "fmt"

func main() {
	FirstIf()
	SecondIf()
	ThirdIf()
}

func FirstIf() {
	var a int
	fmt.Scan(&a)
	if a < 0 {
		fmt.Println("Число отрицательное")
	} else if a > 0 {
		fmt.Println("Число положительное")
	} else {
		fmt.Println("Ноль")
	}
	fmt.Println("End of first programm")
}
func SecondIf() {
	var c int
	fmt.Scan(&c)
	switch {
	case c > 0:
		fmt.Println("Число положительное")
	case c < 0:
		fmt.Println("Число отрицательное")
	default:
		fmt.Println("Ноль")
	}
	fmt.Println("End of second programm")

}
func ThirdIf() {
	var j float64
	fmt.Scan(&j)
	if j > 0 {
		fmt.Println("Число положительное")
	}
	if j < 0 {
		fmt.Println("Число отрицательное")
	}
	if j == 0 {
		fmt.Println("Ноль")
	}
	fmt.Println("End of third programm")

}
