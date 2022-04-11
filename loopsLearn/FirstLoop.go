package main

import "fmt"

func main() {
	FirstCube()
	SecondCube()
	ThirdCube()
}

func FirstCube() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i * i)
	}
	fmt.Println("End of first programm")
}
func SecondCube() {
	intList := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range intList {
		fmt.Println(num * num)
	}
	fmt.Println("End of second programm")

}

func ThirdCube() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i * i)
	}
	fmt.Println("End of third programm")

}
