package main

import "fmt"

func main() {
	a := 12
	b := 0
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获到异常")
		}
		fmt.Println("Hello")
	}()
	fmt.Println(div(a, b))
}
func div(a, b int) (int, error) {
	if b == 0 {
		panic("被除数不能为0")
	}
	return a / b, nil
}
