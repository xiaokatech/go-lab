package main

import "fmt"

func main() {
	// 传统循环
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // skip even number
		}
		fmt.Println("Odd:", i)
	}

	// 类似 while 的循环
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("Sum:", sum)

	// 遍历 数组/切片
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// 遍历 map
	m := map[string]int{"a": 1, "b": 2}
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// 忽略索引或值
	for _, value := range numbers {
		fmt.Println("Value:", value)
	}

	for key := range m {
		fmt.Println("Key:", key)
	}

	// 遍历 字符串
	str := "Hello, 世界"
	for i, r := range str {
		fmt.Printf("Position: %d, Rune: %c with type %T\n", i, r, r)
		fmt.Println(r)
	}

	// break 跳出循环
	for i := 0; ; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// 带标签的break
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i*j == 6 {
				break outerLoop // 跳出外层循环
			}
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}

	// continue 跳过当前迭代
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}
}
