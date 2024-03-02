package main

import "fmt"

func main() {
	const premiumPlanName = "Preminum Plan"
	// premiumPlanName = "Basic Plan" // !Error
	const basicPlanName = "Basic Plan"

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
}
