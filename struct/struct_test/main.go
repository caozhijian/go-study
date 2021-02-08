package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var c Course = Course{
		Name:  "语文",
		Price: 100,
		Url:   "https://www.google.com",
	}
	fmt.Printf("课程的名称:%v,课程的价格:%v,课程的地址:%v\n",
		c.Name, c.Price, c.Url)

	c2 := Course{"数学", 150, "https://www.google.com"}
	fmt.Println(c2.Name)

	//指向结构体的指针
	c3 := &Course{"化学", 150, "https://www.google.com"}
	fmt.Printf("%T\n", c3) //*main.Course
	//fmt.Println((*c3).Name,(*c3).Price,(*c3).Url)
	//go语言内部会将c3.Name转换成(*c3).Name
	fmt.Println(c3.Name, c3.Price, c3.Url)

	//零值
	c4 := Course{}
	fmt.Println(c4.Price)

	fmt.Println("========")
	//多种方式零值初始化结构体
	var c5 Course = Course{}
	var c6 Course
	var c7 *Course = new(Course)
	fmt.Println(c5.Price)
	fmt.Println(c6.Price)
	fmt.Println(c7.Price)

	//结构体是值类型
	c8 := Course{"政治", 100, "https://www.google.com"}
	c9 := c8
	c9.Name = "生物"
	fmt.Printf("c8=%v\n", c8)
	fmt.Printf("c9=%v\n", c9)

	//结构体占用内存的大小，可以使用Sizeof来查看对象占用的类型
	fmt.Println(unsafe.Sizeof(1))        //8
	fmt.Println(unsafe.Sizeof("golang")) //16
	fmt.Println(unsafe.Sizeof(c8))       //16+8+16=40

}

type Course struct {
	Name  string
	Price int
	Url   string
}
