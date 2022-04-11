package main

import "fmt"

func main() {
	FourArray1()
	FourArray2()
	FourArray3()
	FourArray4()
}

func FourArray1() {
	var a, N int
	var array []int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		array = append(array, a)
	}
	for i := 0; i < len(array); i += 2 {
		fmt.Printf("%d ", array[i])
	}
	fmt.Println("End of first programm")

}
func FourArray2() {
	var n int
	fmt.Scan(&n)

	m := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m[i])
	}

	for i := 0; i < n; i += 2 {
		fmt.Print(m[i], " ")
	}
	fmt.Println("End of second programm")

}
func FourArray3() {
	var l, v int
	fmt.Scan(&l)
	for i := 0; i < l; i++ {
		fmt.Scan(&v)
		if i%2 == 0 {
			fmt.Print(v, " ")
		}
	}
	fmt.Println("End of third programm")

}
func FourArray4() {
	var size, inputData int
	fmt.Scan(&size)
	for i := 0; i < size; i++ {
		if fmt.Scan(&inputData); i%2 == 0 {
			fmt.Print(inputData, " ")
		}
	}
	fmt.Println("End of four programm")

}
