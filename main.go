package main

import (
	"fmt"

	"github.com/Kaniek99/AIbasics/src/genotype"
	kp "github.com/Kaniek99/AIbasics/src/knapsack"
)

const (
	itemMinWeight     = 10
	itemMaxWeight     = 90
	itemMinValue      = 10
	itemMaxValue      = 90
	numberOfItems     = 10
	knapsackMaxWeight = 2500
)

func main() {
	items := []*kp.Item{}
	for i := 0; i < numberOfItems; i++ {
		items = append(items, kp.NewRandomItem(itemMinWeight, itemMaxWeight, itemMinValue, itemMaxValue))
	}

	// Generate a Genotype
	bs, err := genotype.GenerateBinarySequence(numberOfItems)
	if err != nil {
		fmt.Println(err)
		return
	}

	knapsack := kp.NewKnapsack(knapsackMaxWeight, &bs)

	knap := kp.NewKnapsackProblem(items, knapsack)
	knap.Solve()
	fmt.Println("Solution: ", knap.Knapsack.Solution.GetGenesSequence())
	fmt.Println("Solution value: ", knap.FitnessCoefficient)
}
