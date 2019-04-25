package main

import "fmt"

func lengthOfNonRepeatingSubstr(s string) int {
	lastOccrred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {

		if lastIndex, ok := lastOccrred[ch]; ok && lastIndex >= start {
			start = lastOccrred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccrred[ch] = i
	}

	return maxLength
}

func main() {
	// 实验一：计算最长不含有重复字符的子串长度：
	fmt.Println("Exp1：")
	fmt.Println("[\"abcabcbb\"] :", lengthOfNonRepeatingSubstr("abcabcbb"))
	fmt.Println("[\"bbbbbb\"] :", lengthOfNonRepeatingSubstr("bbbbbb"))
	fmt.Println("[\"pwwkew\"] :", lengthOfNonRepeatingSubstr("pwwkew"))
	fmt.Println("[\"空字符串\"] :", lengthOfNonRepeatingSubstr(""))
	fmt.Println("[\"一二三三二一\"] :", lengthOfNonRepeatingSubstr("一二三三二一"))
	fmt.Println("[\"这里是跟我Go\"] :", lengthOfNonRepeatingSubstr("这里是跟我Go"))

}
