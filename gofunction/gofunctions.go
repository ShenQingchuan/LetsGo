package main

import (
	"fmt"
	"time"
)

func traditionalSwap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func newSwap(a *int, b *int) (int, int) {
	return *b, *a
}

func tictok(innerf func(a *int, b *int)) func(_a *int, _b *int) {
	return func(_a *int, _b *int) {
		start := time.Now()
		time.Sleep(time.Millisecond)
		innerf(_a, _b)
		end := time.Since(start).Seconds()
		fmt.Println("Time spent: ", end*1000000-1, " ns.")
	}
}

func main() {
	// 实验1：用函数实现两个整数的交换，最后探讨一下 Golang 中实现值 swap 交换终极杀招
	a, b := 4, 5
	fmt.Println("Exp1.(Traditional Swap):\nBefore => a: ", a, " b: ", b)
	traditionalSwap(&a, &b)
	fmt.Println("After => a: ", a, " b: ", b)

	c, d := 6, 7
	fmt.Println("\nExp1.(new Swap):\nBefore => c: ", c, " d: ", d)
	c, d = newSwap(&c, &d)
	fmt.Println("After => c: ", c, " d: ", d)

	e, f := 9, 0
	fmt.Println("\nExp1.(killer Swap):\nBefore => e: ", e, " f: ", f)
	e, f = f, e
	fmt.Println("After => e: ", e, " f: ", f)

	// 实验2：Golang 的 “装饰器” 写法：（函数的扩展 / 闭包）
	// 设计一个计时函数用来测试一下 Golang 的传统swap会用时多久呢？
	/* 函数 func 在 Golang 中是 “一等公民”！ 顾名思义，变量的值都可以是函数哦！*/
	g, h := 11, 12
	fmt.Println("\n\nExp2.(Time of Traditional Swap):\nBefore => g: ", g, " h: ", h)
	result := tictok(traditionalSwap)
	result(&g, &h)
	fmt.Println("After => g: ", g, " h: ", h)
}
