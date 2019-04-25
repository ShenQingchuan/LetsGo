package main

import (
	"fmt"
	"strings"
)

func main() {
	// 实验一：Go 和 Python 字符串读取简便度大pk！Go 能否通过"[]" 的方式读取字符串内的字符呢？
	var s string = "hello"
	// 结果1：得到了 'l' 的 Unicode 编码值，
	/* 这是因为: Golang 的 string 类型实际上是二进制串，如果单独去取，得到的也只会是元素的二进制值 */
	fmt.Println("Exp1:\nfmt.Println: ", s[2])
	// 结果2：由于加上了格式控制，Go 在此就会把该元素按 rune字符形式 来输出
	fmt.Printf("fmt.Printf: %c\n", s[1])

	// 实验二：Golang 用 string 存二进制数据与对字符串的遍历、切割、拼合
	/* 2.1 字符串是直接呈现的 UTF-8 编码 */
	stra := "\xe8\x8b\x9f"
	fmt.Printf("\nExp2: \nstring a: %s", stra)

	/* 2.2 拆了又装！ */
	strb := "1999/11/06"
	bparts := strings.Split(strb, "/")
	for i, ele := range bparts {
		fmt.Printf("\nin string b part[%d]: %s", i, ele)
	}
	strc := strings.Join(bparts, "-")
	fmt.Println("\n\nJoined string c: ", strc)

}
