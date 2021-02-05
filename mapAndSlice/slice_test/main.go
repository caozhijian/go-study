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

	//append,可以向slice追加元素
	subCourse3 = append(subCourse3, "相声")
	fmt.Println(subCourse3)

	//copy函数,拷贝时，需要设置好目标对象长度
	//subCourse4 := []string{}
	subCourse4 := make([]string, len(subCourse3))
	copy(subCourse4, subCourse3)
	fmt.Println(subCourse4)

	//合并两个切片
	slice1 := []string{"slice1"}
	slice2 := []string{"slice2"}
	subCourse5 := append(slice1, slice2...)
	fmt.Println(subCourse5)

	//删除元素
	deleteCourses := [5]string{"数学", "物理", "化学", "生物", "小提琴"}
	courseSlice := deleteCourses[:]
	courseSlice = append(courseSlice[:1], courseSlice[2:]...) //删除物理
	fmt.Println(courseSlice)

}

func replace(mySlice []string) {
	mySlice[0] = "逗你玩"
}
