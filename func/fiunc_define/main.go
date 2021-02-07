package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := div2(12, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		println(result)
	}
}

func add(a, b int) int {
	return a + b
}

func add2(a, b int) (sum int) {
	sum = a + b
	return sum
}

//返回多个值
func add3(a, b int) (sum int) {
	sum = a + b
	return
}

func div(a, b int) (int, error) {
	var result int
	var err error
	if b == 0 {
		err = errors.New("被除数不能为零")
	} else {
		result = a / b
	}
	return result, err
}

func div2(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("被除数不能为零")
	} else {
		result = a / b
	}
	return
}
