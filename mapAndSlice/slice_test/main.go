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
	fmt.Println("===============")

	//slice原理
	//第一种现象
	a := make([]int, 0)
	b := []int{1, 2, 3}
	fmt.Println(copy(a, b))
	fmt.Println(a) //不会去扩展a的空间

	//第二种现象
	c := b[:]
	c[0] = 8
	fmt.Println(b) //Output:[8 2 3]
	fmt.Println(c) //Output:[8 2 3],修改元素值，影响的是同一块儿内存空间

	//第三种现象
	c = append(c, 4)
	fmt.Println(b) //Output:[8 2 3]
	fmt.Println(c) //Output:[8 2 3 4]
	c[1] = 9
	fmt.Println(b) //Output:[8 2 3]
	fmt.Println(c) //Output:[8 9 3 4];
	// 为什么append后执行c[1]=9就不会影响原来的数组
	//涉及到了切片的扩容，一旦产生扩容，切片就会指向新的内存地址

	//第四种现象
	fmt.Println(len(c))
	fmt.Println(cap(c))

	fmt.Println("===================")
	//使用make方式初始化，len与cap分别是多少;默认不会有多余的预留空间，len==cap
	sliceTest := make([]int, 5)
	//sliceTest的len是5,cap是5
	fmt.Printf("sliceTest的len是%d,cap是%d\n", len(sliceTest), cap(sliceTest))
	sliceTest2 := make([]int, 5, 8)
	//sliceTest2的len是5,cap是8
	fmt.Printf("sliceTest2的len是%d,cap是%d\n", len(sliceTest2), cap(sliceTest2))

	//通过数组取切片
	arrayTest := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceTest3 := arrayTest[2:4]
	//fmt.Println(sliceTest3)
	for index, value := range sliceTest3 {
		fmt.Println(index, value)
	}
	//sliceTest3的len是2,cap是8
	fmt.Printf("sliceTest3的len是%d,cap是%d\n", len(sliceTest3), cap(sliceTest3))

	sliceTest4 := []int{1, 2, 3}
	//sliceTest4的len是3,cap是3
	fmt.Printf("sliceTest4的len是%d,cap是%d\n", len(sliceTest4), cap(sliceTest4))

	//切片扩容,扩容时会影响速度
	sliceTest5 := make([]int, 0)
	fmt.Printf("sliceTest5的len是%d,cap是%d\n", len(sliceTest5), cap(sliceTest5))
	sliceTest5 = append(sliceTest5, 1)
	//sliceTest5的len是1,cap是1
	fmt.Printf("sliceTest5的len是%d,cap是%d\n", len(sliceTest5), cap(sliceTest5))
	sliceTest5 = append(sliceTest5, 2)
	//sliceTest5的len是2,cap是2
	fmt.Printf("sliceTest5的len是%d,cap是%d\n", len(sliceTest5), cap(sliceTest5))
	sliceTest5 = append(sliceTest5, 3)
	//sliceTest5的len是3,cap是4
	fmt.Printf("sliceTest5的len是%d,cap是%d\n", len(sliceTest5), cap(sliceTest5))
	sliceTest5 = append(sliceTest5, 4)
	//sliceTest5的len是4,cap是4
	fmt.Printf("sliceTest5的len是%d,cap是%d\n", len(sliceTest5), cap(sliceTest5))
	sliceTest5 = append(sliceTest5, 5)
	//sliceTest5的len是5,cap是8
	fmt.Printf("sliceTest5的len是%d,cap是%d\n", len(sliceTest5), cap(sliceTest5))

}

func replace(mySlice []string) {
	mySlice[0] = "逗你玩"
}
