package salesman

import (
	"math"

	"github.com/Kaniek99/AIbasics/src/genotype"
)

type Coordinates struct {
	X float64
	Y float64
}

func (c *Coordinates) DistanceTo(destination Coordinates) float64 {
	return math.Sqrt(math.Pow(destination.X-c.X, 2) + math.Pow(destination.Y-c.Y, 2))
}

type City struct {
	Coordinates []float64 // Coordinates
	ID          int
}

type Salesman struct {
	CitiesDistances [][]float64
	Solution        genotype.Genotype
}

func NewSalesman(cd [][]float64, solution genotype.Genotype) *Salesman {
	return &Salesman{cd, solution}
}

func (s *Salesman) CalculateFitness() float64 {
	sequence := s.Solution.GetGenesSequence()
	totalDistance := 0.0

	for i := 1; i < len(sequence); i++ {
		firstCity, secondCity := sequence[i-1], sequence[i]
		totalDistance += s.CitiesDistances[firstCity][secondCity]
	}

	return totalDistance
}

type SalesmanProblem struct {
	FitnessCoefficient float64
	Salesman           *Salesman
	StagnationCounter  int
}

func NewSalesmanProblem(salesman *Salesman) *SalesmanProblem {
	fc := salesman.CalculateFitness()
	return &SalesmanProblem{fc, salesman, 0}
}

func (sp *SalesmanProblem) Solve() {
	for sp.StagnationCounter < 100 {
		childSolution := sp.Salesman.Solution.Swap()
		child := NewSalesman(sp.Salesman.CitiesDistances, childSolution)
		childFitness := child.CalculateFitness()
		if childFitness > sp.FitnessCoefficient {
			sp.Salesman = child
			sp.FitnessCoefficient = childFitness
			sp.StagnationCounter = 0
		} else {
			sp.StagnationCounter++
		}
	}
}
