package main

import "fmt"

func main() {
	var i int = 1
	i = 2
	defer show(i) 

	defer fmt.Println("I am closing the file") // defer exec when your func stops or return
	fmt.Println("hello")
	return
	//panic("I am panicking") // stops your function
	fmt.Println("Bye")

}
func show(i int) {

	fmt.Println(i)
}
