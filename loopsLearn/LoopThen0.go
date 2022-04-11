package main

import "fmt"

func main() {
	FirstCheck0()
	SecondCheck0()
	ThirdCheck0()
}
func FirstCheck0() {
	var max = 0
	var count = 1
	var n int
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		if max < n {
			max = n
			count = 1
		} else if max == n {
			count++
		}
	}
	fmt.Println(count)
	fmt.Println("End of first programm")

}

func SecondCheck0() {
	var a, s, p int

	for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
		if a > s {
			s = a
			p = 0
		}
		if a == s {
			p++
		}
	}
	fmt.Println(p)
	fmt.Println("End of second programm")

}
func ThirdCheck0() {
	var a, max, count int

	for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
		switch {
		case a > max:
			max, count = a, 1
		case a == max:
			count++
		}
	}
	fmt.Println(count)
	fmt.Println("End of third programm")

}
