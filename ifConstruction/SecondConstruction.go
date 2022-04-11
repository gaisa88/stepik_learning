package main

import "fmt"

func main() {
	Firstcomparison()
	SecondComparison()
	ThirdComparison()
}
func Firstcomparison() {
	var a int
	var b int
	var c int
	var d int
	fmt.Scan(&a)
	b = a / 100
	c = a / 10 % 10
	d = a % 10
	if b == c || c == d || b == d {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
	fmt.Println("End of first programm")

}
func SecondComparison() {
	var a string
	fmt.Scan(&a)

	if a[0] == a[1] || a[0] == a[2] || a[1] == a[2] {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
	fmt.Println("End of second programm")

}
func ThirdComparison() {
	var a, b, c, input int
	fmt.Scan(&input)
	a = input / 100
	b = input / 10 % 10
	c = input % 10
	if (a != b) && (b != c) && (c != a) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	fmt.Println("End of third programm")

}
