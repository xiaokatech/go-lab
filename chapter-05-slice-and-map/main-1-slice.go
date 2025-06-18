package main

import "fmt"

func main_1() {
	var myInts [10]int
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("myInts:", myInts)
	fmt.Println("primes:", primes)

	mySlice := primes[1:4]
	fmt.Println("mySlice:", mySlice)

	// init with make
	mySlice_2 := make([]int, 5, 10)
	mySlice_3 := make([]int, 5)
	fmt.Println("mySlice_2:", mySlice_2)
	fmt.Println("mySlice_3:", mySlice_3)
	fmt.Println("cap(mySlice_2):", cap(mySlice_2))
	fmt.Println("cap(mySlice_3):", cap(mySlice_3))

	mySlice_4 := append(mySlice_2, mySlice_3...)
	fmt.Println("mySlice_4:", mySlice_4)
	fmt.Println("len(mySlice_4):", len(mySlice_4))
	fmt.Println("cap(mySlice_4):", cap(mySlice_4))

	var costs = []Cost{
		{day: 0, value: 3},
		{day: 3, value: 4},
		{day: 5, value: 9},
	}

	result := getCostsByDay(costs)
	fmt.Println("result:", result)
}

type Cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []Cost) []float64 {
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		// fmt.Println("cost.day", cost.day)
		// fmt.Println("costsByDay", costsByDay)
		// fmt.Println("len(costsByDay)", len(costsByDay))
		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[cost.day] += cost.value
	}
	return costsByDay
}
