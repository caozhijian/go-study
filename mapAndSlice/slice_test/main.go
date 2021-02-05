package main

import "fmt"

func main() {
	courses := []string{"语文", "英语", "历史"}
	fmt.Printf("courses的类型是%T\n", courses) //Output: courses的类型是[]string

	//使用make对slice进行声明与初始化
	courses2 := make([]string, 5)
	courses2 = []string{"语文", "英语", "历史"}
	fmt.Println(courses2)

	//通过数组的方式生成切片
	course3 := [5]string{"数学", "物理", "化学", "生物", "小提琴"}
	subCourse3 := course3[1:4]
	fmt.Println(subCourse3)
	fmt.Printf("sunCourse3的类型是%T\n", subCourse3)

	//数组是值传递，但切片是引用传递
	replace(subCourse3)
	fmt.Println(subCourse3) //Output: [逗你玩 化学 生物]
}

func replace(mySlice []string) {
	mySlice[0] = "逗你玩"
}
