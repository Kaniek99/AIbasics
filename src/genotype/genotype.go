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
