package genotype

import (
	"errors"
	"math/rand"
	"time"
)

type Genotype interface {
	GetGenesSequence() []int // same there but I'm not sure if returning a slice of anys is the best soulution
	Mutate() Genotype
	// Crossover(Genotype) Genotype
}

// is it even needed to have another struct for just int sequence?
type BinarySequence struct {
	GenesSequence []int
	Length        int
}

func GenerateBinarySequence(len int) (BinarySequence, error) {
	if len <= 0 {
		return BinarySequence{}, errors.New("length should be greater than 0")
	}

	bs := BinarySequence{make([]int, len), len}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < len; i++ {
		bs.GenesSequence[i] = r.Intn(2)
	}

	return bs, nil
}

func (bs BinarySequence) GetGenesSequence() []int {
	return bs.GenesSequence
}

func (bs BinarySequence) Mutate() Genotype {
	newSequence := make([]int, bs.Length)
	copy(newSequence, bs.GenesSequence)
	index := rand.Intn(bs.Length)
	newSequence[index] = 1 - newSequence[index]
	return BinarySequence{newSequence, bs.Length}
}

// func (bs *BinarySequence) Crossover(other *BinarySequence) {
// 	index := rand.Intn(bs.length)
// 	for i := index; i < bs.length; i++ {
// 		bs.Sequence[i], other.Sequence[i] = other.Sequence[i], bs.Sequence[i]
// 	}
// }

type IntSequence struct {
	GenesSequence []int
	Length        int
	MaxNumber     int
}

func GenerateIntSequence(len, maxNum int) (IntSequence, error) {
	if len <= 0 {
		return IntSequence{}, errors.New("length should be greater than 0")
	}

	is := IntSequence{make([]int, len), len, maxNum}

	for i := 0; i < len; i++ {
		is.GenesSequence[i] = rand.Intn(maxNum)
	}

	return is, nil
}

func (is IntSequence) GetGenesSequence() []int {
	return is.GenesSequence
}

func (is IntSequence) Mutate() Genotype {
	newSequence := make([]int, is.Length)
	copy(newSequence, is.GenesSequence)
	index := rand.Intn(is.Length)
	newSequence[index] = (newSequence[index] + index) % is.MaxNumber
	return IntSequence{newSequence, is.Length, is.MaxNumber}
}
