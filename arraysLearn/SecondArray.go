package main

import "fmt"

func main() {
	SecondArray1()
	SecondArray2()
	SecondArray3()
}
func SecondArray1() {
	var N int
	var n []int

	fmt.Scan(&N)
	if N >= 4 {
		var a int
		for i := 0; i < N; i++ {
			fmt.Scan(&a)
			n = append(n, a)
		}
	}
	fmt.Println(n[3])
	fmt.Println("End of first programm")

}
func SecondArray2() {
	var n int

	fmt.Scan(&n)
	m := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m[i])
	}
	fmt.Println(m[3])
	fmt.Println("End of second programm")

}
func SecondArray3() {
	var length int
	fmt.Scan(&length)

	mySlice := make([]int, length)
	for i := range mySlice {
		fmt.Scan(&mySlice[i])
	}
	fmt.Println(mySlice[3])
	fmt.Println("End of third programm")
}
