package main

import "fmt"

func main() {
	type myByte = byte
	var b myByte
	fmt.Printf("%T\n", b) //uint8

	type myInt int
	var i myInt
	fmt.Printf("%T\n", i) //main.myInt

	type Course struct {
		name  string
		price int
	}

	type Callable interface {
	}

}
