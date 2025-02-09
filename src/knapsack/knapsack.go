package knapsack

import (
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

func (k Knapsack) CalculateFitness(items []*Item) int {
	weights := []int{}
	values := []int{}
	for _, item := range items {
		weights = append(weights, item.Weight)
		values = append(values, item.Value)
	}

	totalValue, totalWeight := 0, 0
	genesSequence := k.Solution.GetGenesSequence()

	for i, selected := range genesSequence {
		if selected == 1 {
			totalValue += values[i]
			totalWeight += weights[i]
		}
	}

	if totalWeight > k.MaxWeight {
		return k.MaxWeight - totalWeight
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
	fitnessCoefficient := knapsack.CalculateFitness(items)
	return &KnapsackProblem{items, knapsack, fitnessCoefficient, 0}
}

func (kp *KnapsackProblem) Solve() {
	for kp.StagnationCounter < 100 {
		childSolution := kp.Knapsack.Solution.Mutate()
		child := NewKnapsack(kp.Knapsack.MaxWeight, childSolution)
		childFitness := child.CalculateFitness(kp.Items)
		if childFitness > kp.FitnessCoefficient {
			kp.Knapsack = child
			kp.FitnessCoefficient = childFitness
			kp.StagnationCounter = 0
		} else {
			kp.StagnationCounter++
		}
	}
}
