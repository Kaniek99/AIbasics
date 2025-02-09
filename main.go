package main

import (
	"fmt"
	"math"
	"math/rand"

	pc "github.com/Kaniek99/AIbasics/src/allocation"
	"github.com/Kaniek99/AIbasics/src/genotype"
	kp "github.com/Kaniek99/AIbasics/src/knapsack"
	sal "github.com/Kaniek99/AIbasics/src/salesman"
)

const (
	minValue          = 10
	maxValue          = 90
	numberOfItems     = 100
	knapsackMaxWeight = 2500
)

func main() {
	items := []*kp.Item{}
	for i := 0; i < numberOfItems; i++ {
		items = append(items, kp.NewRandomItem( /* Weight limits: */ minValue, maxValue /* Value limits: */, minValue, maxValue))
	}

	valueOfItems, weightOfItems := 0, 0
	fmt.Print("Items: ")

	for _, item := range items {
		fmt.Printf("{weight: %d, value: %d} ", item.Value, item.Weight)
		valueOfItems += item.Value
		weightOfItems += item.Weight
	}
	fmt.Printf("\nValue of all items: %d, Weight of all items: %d\n", valueOfItems, weightOfItems)

	bs, err := genotype.GenerateBinarySequence(numberOfItems)
	if err != nil {
		fmt.Println(err)
		return
	}

	knapsack := kp.NewKnapsack(knapsackMaxWeight, &bs)

	knap := kp.NewKnapsackProblem(items, knapsack)
	knap.Solve()
	fmt.Println("Solution:", knap.Knapsack.Solution.GetGenesSequence())
	fmt.Println("Solution value:", knap.FitnessCoefficient)
	fmt.Println()

	coreMultipliers := []float32{1, 1.25, 1.5, 1.75}
	tasksTime := []int{}
	for i := 0; i < numberOfItems; i++ {
		tasksTime = append(tasksTime, rand.Intn(maxValue-minValue+1)+minValue)
	}

	fmt.Println("Core multipliers:", coreMultipliers)
	fmt.Println("Tasks time:", tasksTime)
	taskAllocation, err := genotype.GenerateIntSequence(len(tasksTime), len(coreMultipliers))
	if err != nil {
		fmt.Println(err)
		return
	}

	processor := pc.NewProcessor(coreMultipliers, math.MaxFloat32, tasksTime, &taskAllocation)
	alloc := pc.NewAllocationProblem(processor)
	alloc.Solve()

	fmt.Println("Solution:", alloc.Solution.TaskAllocation.GetGenesSequence())
	fmt.Println("Time required", alloc.FitnessCoefficient)

	cd := [][]float64{}
	for i := 0; i < numberOfItems; i++ {
		cityDistances := make([]float64, numberOfItems)
		cd = append(cd, cityDistances)
	}

	for i := 0; i < numberOfItems; i++ {
		for j := i + 1; j < numberOfItems; j++ {
			cd[i][j] = float64(rand.Intn(maxValue-minValue+1) + minValue)
		}
		for j := 0; j < numberOfItems; j++ {
			cd[j][i] = cd[i][j]
		}
	}
	fmt.Println()
	// fmt.Println("Matrix of city distances:", cd)
	salesmanRoute, err := genotype.GenerateIntSequenceWithUniqueValues(numberOfItems)
	if err != nil {
		fmt.Println(err)
		return
	}

	salesman := sal.NewSalesman(cd, salesmanRoute)
	salesmanProblem := sal.NewSalesmanProblem(salesman)
	salesmanProblem.Solve()
	fmt.Println("Solution:", salesmanProblem.Salesman.Solution.GetGenesSequence())
	fmt.Println("Total distance:", salesmanProblem.FitnessCoefficient)
}
