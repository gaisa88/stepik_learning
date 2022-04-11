package main

import "fmt"

func main() {
	FiveArray1()
	FiveArray2()
	FiveArray3()
}
func FiveArray1() {
	var a, b, N int
	var array []int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		array = append(array, a)
		if array[i] > 0 {
			b++
		}
	}
	fmt.Println(b)
	fmt.Println("End of first programm")
}
func FiveArray2() {
	var n, s int
	fmt.Scan(&n)
	num := make([]int, n)
	for _, n = range num {
		fmt.Scan(&n)
		if n > 0 {
			s++
		}
	}
	fmt.Println(s)
	fmt.Println("End of second programm")

}
func FiveArray3() {
	var n, c int
	fmt.Scan(&n)
	arr := make([]int, n, n)

	for i := 0; i < len(arr); i++ {
		var j int
		fmt.Scan(&j)
		arr[i] = j
		if arr[i] > 0 {
			c++
		}
	}
	fmt.Println(c)
	fmt.Println("End of third programm")

}
