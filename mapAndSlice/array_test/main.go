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

	//数组操作
	//求长度
	array_test := [5]int{1, 2, 3, 4}
	//fmt.Println(len(array_test))
	//fmt.Println(cap(array_test))
	//使用for-range遍历数组
	for i, value := range array_test {
		fmt.Println(i, value)
	}
	//使用for-range求和
	sum := 0
	for _, value := range array_test {
		sum += value
	}
	fmt.Println(sum)

	//数组是值类型
	courseA := [3]string{"语文", "英语", "历史"}
	courseB := [...]string{"数学", "物理", "化学", "生物"}
	fmt.Printf("CourseA的类型是%T\n", courseA)
	fmt.Printf("CourseB的类型是%T\n", courseB)
	//在函数传递参数时，数组作为参数，实际调用的时候是值传递
	printArray(courseB)
	fmt.Println(courseB) //Output:[数学 物理 化学 生物]

}
func printArray(toPrint [4]string) {
	toPrint[0] = "体育"
	fmt.Println(toPrint) //Output: [体育 物理 化学 生物]
}
