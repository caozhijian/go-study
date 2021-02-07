package main

import "fmt"

//省略号的用途 1.函数参数不定长 2.将slice打散 3.初始化数组

func main() {
	// 通过省略号去动态设置多个参数值
	fmt.Println(add(1, 2, 3, 4, 5))
	slice := []int{2, 3, 4, 5, 6}
	fmt.Println(add(slice...)) //将slice打散

	//slice是一种类型，还是引用传递，所以需要谨慎使用
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(add2(slice2))
	fmt.Println(slice2)

	arr := [...]int{1, 2, 3}
	fmt.Printf("%T", arr)
}

func add(params ...int) (sum int) {
	for _, param := range params {
		sum += param
	}
	return
}
func add2(params []int) (sum int) {
	for _, param := range params {
		sum += param
	}
	params[0] = 9
	return
}
