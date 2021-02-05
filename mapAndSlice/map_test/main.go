package main

import "fmt"

func main() {
	//1.字面值
	m1 := map[string]string{
		"m1": "v1",
	}
	fmt.Printf("%v\n", m1)
	//2.make函数
	m2 := make(map[string]string)
	m2["m2"] = "v2"
	fmt.Printf("%v\n", m2)
	//3. 定义一个空的map
	m3 := map[string]string{}
	m3["m3"] = "v3"
	fmt.Printf("%v\n", m3)

	//map的基本操作
	m := map[string]string{
		"a": "va",
		"b": "vb",
		//"d":"",
	}
	//1.进行增加、修改
	m["c"] = "vc"
	m["b"] = "ab"
	fmt.Printf("%v\n", m)

	//2.查询
	v, ok := m["d"]
	//fmt.Println(v,ok)
	if ok {
		fmt.Printf("找到了，值是%v", v)
	} else {
		fmt.Println("没有找到")
	}

	//3.删除
	delete(m, "a")
	fmt.Println(m) //Output: map[b:ab c:vc]

	//4.遍历
	for k, v := range m {
		fmt.Println(k, v)
	}
}
