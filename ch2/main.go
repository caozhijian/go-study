package main

import "fmt"

func main() {
	//变量的定义
	//1.最基础的变量定义
	//var i int
	//i = 20
	//fmt.Println(i)
	//定义并初始化
	//var i int = 10
	//fmt.Println(i)
	//2. 根据值自行推断变量类型
	//var i = 100 //没有设置类型
	//fmt.Println(i)
	//3. 省略var
	//i :=100
	//fmt.Println(i)
	//var a int = 1
	//var b = 2
	//c := 3
	//fmt.Println(a,b,c)

	//多变量定义
	//var a,b,c int
	//a,b,c = 10,20,30
	//fmt.Println(a,b,c)

	//var a,b,c int = 10,20,30
	//fmt.Println(a,b,c)

	//集合类型
	//var (
	//	a int
	//	name string
	//)

	//常量
	//const PI=3.1415926
	const (
		UNKNOWN = 0
		Female  = 1
		Male    = 2
	)
	const (
		x int = 1
		y
		s = "abc"
		z
	)
	fmt.Println(x, y, s, z) // Output: 1 1 abc abc

}
