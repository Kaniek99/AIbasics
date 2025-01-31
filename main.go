package main

import (
	"fmt"

	kp "github.com/Kaniek99/AIbasics/src/knapsack"
)

const (
	itemMinWeight     = 10
	itemMaxWeight     = 90
	itemMinValue      = 10
	itemMaxValue      = 90
	numberOfItems     = 100
	knapsackMaxWeight = 2500
)

func main() {
	knap := kp.NewRandomKnapsackProblem(itemMinWeight, itemMaxWeight, itemMinValue, itemMaxValue, numberOfItems, knapsackMaxWeight)
	knap.Solve()
	fmt.Println("Solution: ", knap.Knapsack.Solution.Sequence)
	fmt.Println("Solution value: ", knap.FitnessCoefficient)
}
