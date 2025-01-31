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
	//TODO: make soulution a Genotype, not necessarily a BinarySequence
	Solution *(utils.BinarySequence)
}

// TODO: make soulution a Genotype, not necessarily a BinarySequence
func NewKnapsack(maxWeight int, solution *utils.BinarySequence) Knapsack {
	return Knapsack{maxWeight, solution}
}

func NewRandomKnapsack(maxWeight, numberOfItems int) Knapsack {
	//TODO: make soulution a Genotype, not necessarily a BinarySequence
	solution, err := utils.GenerateBinarySequence(numberOfItems)

	if err != nil {
		fmt.Println(err)
		return Knapsack{}
	}

	return NewKnapsack(maxWeight, &solution)
}

func (k *Knapsack) Mutate() Knapsack {
	//TODO: make mutated a Genotype, not necessarily a BinarySequence
	mutated := utils.BinarySequence{Sequence: make([]int, k.Solution.Length), Length: k.Solution.Length}
	copy(mutated.Sequence, k.Solution.Sequence)
	mutated = mutated.Mutate()
	return NewKnapsack(k.MaxWeight, &mutated)
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

func NewRandomKnapsackProblem(itemMinWeight, itemMaxWeight, itemMinValue, itemMaxValue, numberOfItems, knapsackMaxWeight int) *KnapsackProblem {
	items := []*Item{}
	for i := 0; i < numberOfItems; i++ {
		items = append(items, NewRandomItem(itemMinWeight, itemMaxWeight, itemMinValue, itemMaxValue))
	}

	knapsack := NewRandomKnapsack(knapsackMaxWeight, numberOfItems)

	return NewKnapsackProblem(items, &knapsack)
}

func CalculateFitness(items []*Item, knapsack *Knapsack) int {
	weight := calculateWeight(items, knapsack.Solution)
	if weight > knapsack.MaxWeight {
		return knapsack.MaxWeight - weight
	}

	// value := calculateValue(items, knapsack.Solution)
	values := []int{}
	for _, item := range items {
		values = append(values, item.Value)
	}

	value := knapsack.Solution.Fitness(values)
	return value
}

// TODO: make soulution a Genotype, not necessarily a BinarySequence
// func calculateValue(items []*Item, solution *(utils.BinarySequence)) int {
// 	value := 0
// 	for i, item := range items {
// 		value += item.Value * solution.Sequence[i]
// 	}

// 	return value
// }

// TODO: make soulution a Genotype, not necessarily a BinarySequence
func calculateWeight(items []*Item, solution *(utils.BinarySequence)) int {
	weight := 0

	for i, item := range items {
		weight += item.Weight * solution.Sequence[i]
	}

	return weight
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
		child := *kp.Knapsack
		child = child.Mutate()
		childFitness := CalculateFitness(kp.Items, &child)
		if childFitness > kp.FitnessCoefficient {
			kp.Knapsack = &child
			kp.FitnessCoefficient = childFitness
			kp.StagnationCounter = 0
		} else {
			kp.StagnationCounter++
		}
	}
}
