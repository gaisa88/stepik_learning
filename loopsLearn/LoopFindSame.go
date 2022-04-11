package main

import "fmt"

func main() {
	FindSame1()
	FindSame2()
	FindSame3()
}
func FindSame1() {
	var x int
	var y int
	var k int
	fmt.Scan(&x)
	fmt.Scan(&y)
	for j := 1000; x/j >= 0; j = (j / 10) {
		if x/j == 0 {
			continue
		}
		k = (x / j) % 10
		for i := 1000; y/i >= 0; i = (i / 10) {
			if y/i == 0 {
				continue
			}
			if k == (y/i)%10 {
				fmt.Print(k, " ")
			}
			if i == 1 {
				break
			}
		}
		if j == 1 {
			break
		}
	}
	fmt.Println("End of first programm")

}
func FindSame2() {
	var a, b, an, bn int
	fmt.Scan(&a, &b)

	for i := 1000; i >= 1 && a > 0; i = i / 10 {
		if a/i == 0 {
			continue
		}
		an = a / i % 10
		for i := 1; i < 10000 && b > 0; i = i * 10 {
			bn = b / i % 10
			if bn == an {
				fmt.Print(an, " ")
				break
			}
		}
	}
	fmt.Println("End of second programm")

}
func FindSame3() {
	var a, b string
	fmt.Scan(&a, &b)
	for i := range a {
		for j := range b {
			if b[j] == a[i] {
				fmt.Print(string(b[j]) + " ")
			}
		}
	}
	fmt.Println("End of third programm")

}
