package main

import (
	"fmt"
)

func main() {
	initialIncome := 250.0
	targetIncome := 32000.0
	annualIncrease := 2.0
	currentIncome := initialIncome
	years := 0

	for currentIncome < targetIncome {
		fivePercent := currentIncome * 0.05
		fmt.Printf("Year %d: Income = %.2f dollars, 5%% of Income = %.2f dollars\n", years+1, currentIncome, fivePercent)
		currentIncome *= annualIncrease
		years++
	}

	fmt.Printf("It will take %d years to reach or exceed the target income of %.2f dollars.\n", years, targetIncome)
}
