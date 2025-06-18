package main

import (
	"errors"
	"fmt"
)

func main() {
	// 带间断语句的 if
	if num := 10; num > 5 {
		fmt.Println(num, "is greater than 5")
	}

	// 在条件中调用函数
	if result, err := doSomething_with_err(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	if result, err := doSomething_with_result(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// 条件中的逻辑晕眩
	age := 18
	var hasLicense bool = true
	if age >= 18 && hasLicense {
		fmt.Println("Can drive")
	} else if age > 16 {
		fmt.Println("Can learn to drive")
	} else {
		fmt.Println("Too young to drive")
	}

	// 无表达式的switch (相当于if-else链)
	var score float32 = 99.5
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}

	// 类型 switch
	var i interface{} = "avc"
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

	// fallthrough 使用
	switch n := 2; n {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		fallthrough
	case 4:
		fmt.Println("Three")
	}

}

func doSomething_with_err() (string, error) {
	return "", errors.New("this is an error")
}

func doSomething_with_result() (string, error) {
	return "abc", nil
}
