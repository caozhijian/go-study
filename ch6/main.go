package main

import "fmt"

func main() {
	//字符串基本操作
	//1.求字符串长度
	var name string = "Hello,少年"
	name_arr := []rune(name)
	fmt.Println(len(name_arr))
}
