package main

import "fmt"

func main() {
	//基本的类型转换
	a := int(3.0)
	fmt.Println(a)
	//1.go语言不支持变量间的隐式类型转换
	var b int = 5.0 //常量到变量会进行隐式转换
	fmt.Println(b)  //Output: 5

	c := 5.6
	fmt.Printf("%T\n", c)

	//var d int = c //会报错,Cannot use 'c' (type float64) as type int
	//fmt.Println(d)

	var d int = int(c)
	fmt.Println(d)
}
