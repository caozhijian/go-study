package main

import "fmt"

func main() {
	var c Course = Course{
		Name:  "语文",
		Price: 100,
		Url:   "https://www.google.com",
	}
	fmt.Printf("课程的名称:%v,课程的价格:%v,课程的地址:%v\n",
		c.Name, c.Price, c.Url)
}

type Course struct {
	Name  string
	Price int
	Url   string
}
