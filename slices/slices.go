package slices

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	{
		// 实验1：切片的长度和容量有什么区别？
		sa := a[2:5]
		fmt.Println("\nExp1:\n[sd]: len = ", len(sa), "cap =", cap(sa))
	}

	{
		// 实验2：请问 切片之于数组，在结构上的区别是什么？ 又是如何扩展的呢？
		sb := a[:5]

		fmt.Println("\nExp2: \nBegin: [sb]: ", sb, "len =", len(sb), ", cap =", cap(sb))
		fmt.Println()

		for i := 0; i < 20; i++ {
			fmt.Println("running: [sb]: len =", len(sb), "cap =", cap(sb))
			sb = append(sb, 2*i+1)
		}

		fmt.Println("\nEnding: [sb]:", sb)
	}

	{
		// 实验3：两个切片涵盖的区域有重叠会怎么样？
		sc := a[:7]
		sd := a[3:]
		sc[5] = 0
		fmt.Println("\nExp3:\nsc: ", sc, "\nsd: ", sd)
	}
}
