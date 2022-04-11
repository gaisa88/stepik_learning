package main

import "fmt"

func main() {
	FirstMinutes()
	SecondMinutes()
	ThirdMinutes()
}

func FirstMinutes() {
	var a int
	var h int
	var m int
	fmt.Scan(&a)
	h = a / 30
	m = 2 * (a % 30)
	fmt.Println("It is ", h, " hours ", m, " minutes.")
	fmt.Println("End of first programm")

}
func SecondMinutes() {
	var d int
	fmt.Scan(&d)
	angle := d * 2
	fmt.Println("It is", angle/60, "hours", angle%60, "minutes.")
	fmt.Println("End of second programm")

}
func ThirdMinutes() {
	var a int
	fmt.Scan(&a)
	fmt.Println("It is", a/30, "hours", (a%30)*2, "minutes.")
	fmt.Println("End of third programm")

}
