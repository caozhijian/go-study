package main

import "fmt"

func main() {
	//fmt.Println("test1")
	//defer之后只能是函数调用，不能是表达式，比如a++
	//多个defer同时出现，会按照先进后出的顺序执行
	//defer fmt.Println("defer test1")
	//defer fmt.Println("defer test2")
	//defer fmt.Println("defer test3")
	//fmt.Println("test2")

	//defer语句执行时的拷贝机制
	//test:=func(){
	//	fmt.Println("testA")
	//}
	//defer test()
	//test=func() {
	//	fmt.Println("testB")
	//}
	//fmt.Println("testC")

	//x:=10
	//defer func(a int) {fmt.Printf("a的值是%v\n",a)}(x)
	//x++
	//
	//y:=10
	//defer func(b *int) {fmt.Printf("b的值是%v\n",*b)}(&y)
	//y++
	//
	//z:=10
	//defer func() {fmt.Printf("z的值是%v\n",z)}()
	//z++

	fmt.Println(f1())  //Output: 50
	fmt.Println(*f2()) //Output: 11
}

func f1() int {
	x := 50
	defer func() {
		x++
	}()
	return x
}

func f2() *int {
	a := 10
	b := &a
	defer func() {
		*b++
	}()
	return b
}
