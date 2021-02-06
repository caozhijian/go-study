package main

import "fmt"

func main() {
	a := 10
	b := 20
	//a,b=b,a
	fmt.Printf("a=%v,b=%v\n", a, b)
	fmt.Printf("变量a的内存地址是%p\n", &a) //变量有地址
	//现在有一种特殊的变量类型，这个变量只能保存地址值
	var ip *int //这个变量里面就只能保存地址类型这种值
	ip = &a
	*ip = 30
	fmt.Printf("现在变量a的值是%v\n", a)
	fmt.Printf("ip指向的内存空间地址是%p,内存中的值是:%d\n", ip, *ip)
	swap(&a, &b)
	fmt.Printf("交换后a的值是%v,b的值是%v\n", a, b)

	//arr:=[3]int{1,2,3}
	//var ip2 *[3]int = &arr
	//
	//var ptrs [3]*int //创建能存放三个指针变量的数组;指针数组
	//指针的默认值是nil
}
func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}
