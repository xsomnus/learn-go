package main

import "fmt"

func main() {
	// 基本使用
	var age int
	fmt.Println("my age is ", age)
	//类型推断
	var age2 = 20
	fmt.Println("my age2 is", age2)
	//可以声明多个变量
	var width, height int
	fmt.Println("width is", width, "height is ", height)

	//同时声明不同的类型的变量
	var (
		name = "x7e"
		age3 = 29
		height2 int
	)
	fmt.Println("my name is", name, ",age is ", age3, "and height is ", height2)

	/*
	简短声明
	*	1.要求 := 操作符的左边至少有一个变量是尚未声明的
	*/
	name2, age4 := "x7e", 25
	fmt.Println("my name is", name2, "age is ", age4)
}
