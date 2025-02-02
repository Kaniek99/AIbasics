package knapsack

import (
	"fmt"
	"math/rand"

	"github.com/Kaniek99/AIbasics/src/genotype"
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
	Solution  genotype.Genotype
}

func NewKnapsack(maxWeight int, solution genotype.Genotype) *Knapsack {
	return &Knapsack{maxWeight, solution}
}

func (k Knapsack) CalculateFitness(values, weights []int, maxWeight int) int {
	totalValue := 0
	totalWeight := 0

	genesSequence := k.Solution.GetGenesSequence()

	for i, selected := range genesSequence {
		if selected == 1 {
			totalValue += values[i]
			totalWeight += weights[i]
		}
	}

	if totalWeight > maxWeight {
		return maxWeight - totalWeight
	}

	return totalValue
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

	return knapsack.CalculateFitness(values, weights, knapsack.MaxWeight)
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
