package main

import "fmt"

func main() {
	fmt.Println("wrongExample:")
	wrongExample()
	fmt.Println("goodExample:")
	goodExample()
}

func wrongExample() {
	x := 5
	increment(x)

	fmt.Println(x)
	/*
		stiull prints 5
		because the increment function
		received a copy of x
	*/
}

func increment(x int) {
	x++
}

func goodExample() {
	sendsSoFar := 430
	const sendsToAdd = 25
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("You've sent", sendsSoFar, "messages")
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}
