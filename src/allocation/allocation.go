package allocation

import (
	"slices"

	"github.com/Kaniek99/AIbasics/src/genotype"
)

type Processor struct {
	CoreMultipliers []float32 // index is coreID, value is multiplier
	RequiredTime    float32
	TasksTime       []int // index is taskID, value is required time
	TaskAllocation  genotype.Genotype
}

func NewProcessor(coreMultipliers []float32, requiredTime float32, tasksTime []int, taskAllocation genotype.Genotype) *Processor {
	return &Processor{coreMultipliers, requiredTime, tasksTime, taskAllocation}
}

func (p *Processor) CalculateFitness() float32 {
	requierdTime := make([]float32, len(p.CoreMultipliers))
	assignment := p.TaskAllocation.GetGenesSequence()
	for i, coreID := range assignment {
		requierdTime[coreID] += p.CoreMultipliers[coreID] * float32(p.TasksTime[i])
	}

	return slices.Max(requierdTime)
}

type AllocationProblem struct {
	FitnessCoefficient float32
	Solution           *Processor
	StagnationCounter  int
}

func NewAllocationProblem(solution *Processor) *AllocationProblem {
	fitnessCoefficient := solution.CalculateFitness()
	return &AllocationProblem{fitnessCoefficient, solution, 0}
}

func (ap *AllocationProblem) Solve() {
	for ap.StagnationCounter < 100 {
		childSolution := ap.Solution.TaskAllocation.Mutate()
		child := NewProcessor(ap.Solution.CoreMultipliers, ap.FitnessCoefficient, ap.Solution.TasksTime, childSolution)
		childFitness := child.CalculateFitness()
		if childFitness < ap.FitnessCoefficient {
			ap.FitnessCoefficient = childFitness
			ap.Solution = child
			ap.StagnationCounter = 0
		} else {
			ap.StagnationCounter++
		}
	}
}
