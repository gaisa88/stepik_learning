package main

import "fmt"

func main() {
	FirstLoopContinue()
	SecondLoopContinue()
	ThirdLoopContinue()
	FourLoopContinue()
}
func FirstLoopContinue() {
	var x int
	for fmt.Scan(&x); x <= 100; fmt.Scan(&x) {
		switch {
		case x < 10:
			continue
		case x > 100:
			break
		default:
			fmt.Println(x)
		}
	}
	fmt.Println("End of first programm")

}
func SecondLoopContinue() {
	var n int
	for fmt.Scan(&n); n <= 100; fmt.Scan(&n) {
		if n >= 10 {
			fmt.Println(n)
		}
	}
	fmt.Println("End of second programm")

}
func ThirdLoopContinue() {
	for i := 0; i <= 100; fmt.Scan(&i) {
		if i >= 10 {
			fmt.Println(i)
		}
	}
	fmt.Println("End of third programm")

}
func FourLoopContinue() {
	var n int

	for {
		fmt.Scan(&n)
		if n < 10 {
			continue
		}
		if n > 100 {
			break
		}
		fmt.Println(n)
	}
	fmt.Println("End of four programm")

}
