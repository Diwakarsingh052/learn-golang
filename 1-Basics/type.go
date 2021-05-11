package main

import "fmt"

type money float64

//type student struct {
//
//}

func main() {
	var rupee money = 100.4
	//var f float32 = float32(rupee)
	//fmt.Println(rupee, f)
	rupee.show()

}

func (m money) show() {
	fmt.Println(m)
}
