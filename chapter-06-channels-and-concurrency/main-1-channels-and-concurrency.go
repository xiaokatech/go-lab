package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("abx")

	emails := []Email{
		{date: time.Now()},
		// {date: time.Now()},
		// {date: time.Now()},
	}
	filteroldEmails(emails)
}

type Email struct {
	date time.Time
}

func filteroldEmails(emails []Email) {
	isOldChain := make(chan bool)

	go func() {
		defer close(isOldChain) // 确保关闭channel
		for _, e := range emails {
			if e.date.After(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
				fmt.Println("hit condition")
				isOldChain <- true
				continue
			}
			isOldChain <- false
		}
	}()

	isOld := <-isOldChain
	fmt.Println("email 1 is old:", isOld)
	isOld = <-isOldChain
	fmt.Println("email 2 is old:", isOld)
	isOld = <-isOldChain
	fmt.Println("email 3 is old:", isOld)
}
