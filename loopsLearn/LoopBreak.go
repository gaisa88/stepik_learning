package main

import "fmt"

func main() {
	FirstBreak()
	SecondBreak()
	ThirdBreak()
}
func FirstBreak() {
	var n int
	var c int
	var d int
	fmt.Scan(&n)
	fmt.Scan(&c)
	fmt.Scan(&d)
	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
	}
	fmt.Println("End of first programm")

}

func SecondBreak() {
	var n, c, d int
	fmt.Scan(&n, &c, &d)
	for i := c; i <= n; i += c {
		if i%d != 0 {
			fmt.Println(i)
			return
		}
	}
	fmt.Println("End of second programm")

}
func ThirdBreak() {
	var n, c, d int
	fmt.Scan(&n, &c, &d)
	for i := c; i <= n; i += c {
		if i%d != 0 {
			fmt.Println(i)
			break
		}
	}
	fmt.Println("End of third programm")

}
