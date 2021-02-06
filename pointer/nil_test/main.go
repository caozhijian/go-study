package main

import "fmt"

func main() {
	//make与new函数
	//new函数的用法
	var p *int = new(int) //go的编译器会申请一个内存空间，并将其值设置为0
	*p = 10
	//make函数
	var info map[string]string = make(map[string]string)
	info["name"] = "hello"

	//new函数返回的是这个值的地址（指针），make函数返回的是指定类型的实例
	//nil的一些细节
	var info2 map[string]string
	if info2 == nil {
		fmt.Println("map的默认值是nil")
	}
	var slice []string
	if slice == nil {
		fmt.Println("slice的默认值是nil")
	}
	var err error
	if err == nil {
		fmt.Println("error的默认值是nil")
	}
	//nil是唯一可以用来表示部分类型的零值的标识符，它可以代表许多不同内存布局的值

}
