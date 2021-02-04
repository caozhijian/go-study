package main

import (
	"fmt"
	"strconv"
)

func main() {
	//基本的类型转换
	//a := int(3.0)
	//fmt.Println(a)
	//1.go语言不支持变量间的隐式类型转换
	//var b int = 5.0 //常量到变量会进行隐式转换
	//fmt.Println(b)  //Output: 5

	//c := 5.6
	//fmt.Printf("%T\n", c)

	//var d int = c //会报错,Cannot use 'c' (type float64) as type int
	//fmt.Println(d)

	//var d int = int(c)
	//fmt.Println(d)

	//int转字符串，Itoa
	//var t int = 3
	//fmt.Printf("%T\n", t)
	//s := strconv.Itoa(t)
	//fmt.Printf("%T\n", s)

	//string转int，Atoi
	//var s2 string = "3"
	//fmt.Printf("%T\n", s2)
	//i, _ := strconv.Atoi(s2)
	//fmt.Printf("%T\n", i)

	//parse类的函数
	b, _ := strconv.ParseBool("true")
	fmt.Println(b)

	f, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Printf("%T", f)
}
