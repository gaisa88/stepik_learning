package main

import "fmt"

func main() {
	Challenge1()
	Challenge2()
	Challenge3()
}
func Challenge1() {
	var a int
	var b int
	var c int
	var d int
	fmt.Scan(&a)
	b = a / 100
	c = a / 10 % 10
	d = a % 10
	fmt.Println(b + c + d)
	fmt.Println("End of first programm")

}
func Challenge2() {
	var a int
	fmt.Scan(&a)
	sum := 0
	for a > 0 {
		sum += a % 10
		a /= 10
	}
	fmt.Println(sum)
	fmt.Println("End of second programm")

}
func Challenge3() {
	var n, sum int

	for fmt.Scan(&n); n > 0; n /= 10 {
		sum += n % 10
	}

	fmt.Println(sum)
	fmt.Println("End of third programm")

}
