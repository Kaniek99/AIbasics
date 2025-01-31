package utils

import (
	"errors"
	"math/rand"
	"time"
)

type Genotype interface {
	Mutate() Genotype
	Fitness([]any) int
	// Crossover(Genotype) Genotype
}

// Each genotype is a binary sequence
type BinarySequence struct {
	Sequence []int
	Length   int
}

func GenerateBinarySequence(len int) (BinarySequence, error) {
	if len <= 0 {
		return BinarySequence{}, errors.New("length should be greater than 0")
	}

	bs := BinarySequence{make([]int, len), len}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < len; i++ {
		// bs.Sequence[i] = rand.Intn(2)
		bs.Sequence[i] = r.Intn(2)
	}

	return bs, nil
}

func (bs BinarySequence) Mutate() BinarySequence {
	// rand.Seed(time.Now().UnixNano())
	index := rand.Intn(bs.Length)
	bs.Sequence[index] = 1 - bs.Sequence[index]
	return bs
}

func (bs BinarySequence) Fitness(values []int) int {
	fitnessCoefficient := 0
	for i := 0; i < bs.Length; i++ {
		fitnessCoefficient += bs.Sequence[i] * values[i]
	}
	return fitnessCoefficient
}

// func (bs *BinarySequence) Crossover(other *BinarySequence) {
// 	index := rand.Intn(bs.length)
// 	for i := index; i < bs.length; i++ {
// 		bs.Sequence[i], other.Sequence[i] = other.Sequence[i], bs.Sequence[i]
// 	}
// }

func SumOfTheProductOfTwoSlices(first, second []int) (int, error) {
	sum := 0

	if len(first) != len(second) {
		return sum, errors.New("slices should have the same length")
	}

	for i := 0; i < len(first); i++ {
		sum += first[i] * second[i]
	}

	return sum, nil
}
