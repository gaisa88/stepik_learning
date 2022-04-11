package main

import "fmt"

func main() {
	FirstArray1()
	FirstArray2()
	FirstArray3()
}
func FirstArray1() {
	var workArray [10]uint8
	var indexArray [6]uint8

	for i := range workArray {
		fmt.Scan(&workArray[i])
	}
	for i := range indexArray {
		fmt.Scan(&indexArray[i])
	}
	for i := 0; i < len(indexArray); i += 2 {
		tempI := workArray[indexArray[i]]
		workArray[indexArray[i]] = workArray[indexArray[i+1]]
		workArray[indexArray[i+1]] = tempI
	}
	for i := 0; i < len(workArray); i++ {
		fmt.Print(workArray[i], " ")
	}
	fmt.Println("End of first programm")

}
func FirstArray2() {
	workArray := [10]uint8{}
	for i := range workArray {
		fmt.Scan(&workArray[i])
	}
	var x, y int
	for n := 0; n < 3; n++ {
		fmt.Scan(&x, &y)
		workArray[x], workArray[y] = workArray[y], workArray[x]
	}
	for i := range workArray {
		fmt.Printf("%d ", workArray[i])
	}
	fmt.Println("End of second programm")

}
func FirstArray3() {
	var workArray [10]uint8
	var swapFirst, swapSecond, i uint8
	for ; i < 13; i++ {
		if i > 9 {
			fmt.Scan(&swapFirst, &swapSecond)
			workArray[swapFirst], workArray[swapSecond] = workArray[swapSecond], workArray[swapFirst]
		} else {
			fmt.Scan(&workArray[i])
		}
	}
	for _, val := range workArray {
		fmt.Print(val, " ")
	}
	fmt.Println("End of third programm")

}
