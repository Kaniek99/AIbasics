package knapsack

import (
	"fmt"
	"math/rand"

	"github.com/Kaniek99/AIbasics/utils"
)

type Item struct {
	Weight int
	Value  int
}

func NewItem(weight, value int) *Item {
	return &Item{weight, value}
}

func NewRandomItem(minWeight, maxWeight, minValue, maxValue int) *Item {
	weight := rand.Intn(maxWeight-minWeight+1) + minWeight
	value := rand.Intn(maxValue-minValue+1) + minValue
	return NewItem(weight, value)
}

type Knapsack struct {
	MaxWeight int
	Solution  utils.Genotype
}

func NewKnapsack(maxWeight int, solution utils.Genotype) *Knapsack {
	return &Knapsack{maxWeight, solution}
}

type KnapsackProblem struct {
	Items              []*Item
	Knapsack           *Knapsack
	FitnessCoefficient int
	StagnationCounter  int
}

func NewKnapsackProblem(items []*Item, knapsack *Knapsack) *KnapsackProblem {
	fitnessCoefficient := CalculateFitness(items, knapsack)
	return &KnapsackProblem{items, knapsack, fitnessCoefficient, 0}
}

func CalculateFitness(items []*Item, knapsack *Knapsack) int {
	weights := []int{}
	for _, item := range items {
		weights = append(weights, item.Weight)
	}

	values := []int{}
	for _, item := range items {
		values = append(values, item.Value)
	}

	return knapsack.Solution.CalculateFitness(values, weights, knapsack.MaxWeight)
}

func (kp *KnapsackProblem) Solve() {
	valueOfItems, weightOfItems := 0, 0
	fmt.Print("Items: ")

	for _, item := range kp.Items {
		fmt.Print(item, " ")
		valueOfItems += item.Value
		weightOfItems += item.Weight
	}
	fmt.Printf("\nValue: %d, Weight: %d\n", valueOfItems, weightOfItems)

	for kp.StagnationCounter < 100 {
		childSolution := kp.Knapsack.Solution.Mutate()
		child := NewKnapsack(kp.Knapsack.MaxWeight, childSolution)
		childFitness := CalculateFitness(kp.Items, child)
		if childFitness > kp.FitnessCoefficient {
			kp.Knapsack = child
			kp.FitnessCoefficient = childFitness
			kp.StagnationCounter = 0
		} else {
			kp.StagnationCounter++
		}
		// Print the solution and the child for each iteration
		// fmt.Println("Solution:", kp.Knapsack.Solution.GetGenesSequence())
		// fmt.Println("   Child:", child.Solution.GetGenesSequence())
	}
}
