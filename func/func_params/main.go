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
	fmt.Printf("%T\n", arr)

	//函数是一等公民，可以作为参数、返回值、赋值给另一个变量
	myFunc := add
	fmt.Printf("myFunc的类型是%T\n", myFunc())
	//匿名函数
	myFunc2 := func(a, b int) int {
		return a + b
	}
	myFunc2(1, 2)

	result := func(a, b int) int {
		return a + b
	}(3, 4)
	fmt.Println(result)

	var mySub sub = subImpl
	var mySub2 sub = func(a, b int) int {
		return a - b
	}
	fmt.Println(mySub(9, 0), mySub2(8, 1))

	//函数可以作为其它函数的参数进行传递
	score := []int{90, 33, 69, 88, 100}
	fmt.Println(filter(score, func(i int) bool {
		if i >= 90 {
			return true
		} else {
			return false
		}
	}))
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

//自定义sub类型
type sub func(a, b int) int

func subImpl(a, b int) int {
	return a - b
}

func filter(scores []int, line func(int) bool) []int {
	reSlice := make([]int, 0)
	for _, score := range scores {
		if line(score) {
			reSlice = append(reSlice, score)
		}
	}
	return reSlice
}
