package main

import "fmt"

func main() {
	First8Check()
	Second8Check()
}
func First8Check() {
	var n, num int
	fmt.Scan(&n)

	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&num)

		if num >= 10 && num < 100 && num%8 == 0 {
			sum += num
		}
	}
	fmt.Println(sum)
	fmt.Println("End of first programm")

}

func Second8Check() {
	var n, sum int
	for fmt.Scan(&n); n > 0; n-- {
		var x int
		if fmt.Scan(&x); x > 9 && x < 100 && x%8 == 0 {
			sum += x
		}
	}
	fmt.Println(sum)
	fmt.Println("End of second programm")

}
