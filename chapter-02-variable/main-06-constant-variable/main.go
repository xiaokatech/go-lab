package main

import "fmt"

func main() {
	const premiumPlanName = "Preminum Plan"
	// premiumPlanName = "Basic Plan" // !Error
	const basicPlanName = "Basic Plan"

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)

	fmt.Println("======")

	calculateTime()
}

func calculateTime() {
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = minutesInHour * secondsInMinute

	fmt.Println("number of seconds in an hour:", secondsInHour)
}
