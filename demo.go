package main

import "fmt"

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main()  {
	// 使用 defer 后的返回值分析
	// 1、方法中的 return 不像看到的那样直接返回对应变量的值
	// 2、而是分解为 对返回变量进行赋值 >> 执行 defer 函数 >> 返回变量值
	fmt.Println(f())	// 1
	fmt.Println(f2())	// 5
	fmt.Println(f3())	// 1
}

