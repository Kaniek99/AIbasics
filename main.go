package main

import (
	"fmt"
	"math"
	"math/rand"

	pc "github.com/Kaniek99/AIbasics/src/allocation"
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

	coreMultipliers := []float32{1, 1.25, 1.5, 1.75}
	tasksTime := []int{}
	for i := 0; i < numberOfItems; i++ {
		tasksTime = append(tasksTime, rand.Intn(itemMaxValue-itemMinValue+1)+itemMinValue)
	}

	fmt.Println(tasksTime)
	taskAllocation, err := genotype.GenerateIntSequence(len(tasksTime), len(coreMultipliers))
	if err != nil {
		fmt.Println(err)
		return
	}

	processor := pc.NewProcessor(coreMultipliers, math.MaxFloat32, tasksTime, &taskAllocation)
	alloc := pc.NewAllocationProblem(processor)
	alloc.Solve()

	fmt.Println(alloc.Solution.TaskAllocation.GetGenesSequence())
	fmt.Println("Time required", alloc.FitnessCoefficient)
}
