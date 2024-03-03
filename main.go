package main

import (
	"fmt"
	"time"

	tinytime "github.com/wagslane/go-tinytime"
	"github.com/xiaokatech/go-lab-package/mystrings"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(
		mystrings.Reverse("hello world"),
	)

	tt := tinytime.New(1585750374)

	tt = tt.Add(time.Hour * 48)
	fmt.Println(tt)
}
