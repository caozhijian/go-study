package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//定义布尔类型
	//var gender bool = false
	gender := true
	fmt.Println(gender)

	var age int = 18
	fmt.Println(unsafe.Sizeof(age))
}
