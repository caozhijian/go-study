package main

import "fmt"

func main() {
	//数组的声明与初始化
	//var courses [10] string
	//var courses = [5]string{"golang","python"}
	courses := [...]string{"golang", "python"}
	fmt.Println(courses)

	e := [5]int{4: 100}
	fmt.Println(e)
}
