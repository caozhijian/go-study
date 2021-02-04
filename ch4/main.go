package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	//定义布尔类型
	//var gender bool = false
	gender := true
	fmt.Println(gender)

	var age int = 18
	fmt.Println(unsafe.Sizeof(age))

	//float类型
	var weight float32 = 92.5
	fmt.Println(weight)
	fmt.Println(math.MaxFloat32)

	weightFree := 80.5
	fmt.Printf("%T", weightFree) //Output: float64
}
