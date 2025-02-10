package genotype

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
)

type Genotype interface {
	GetGenesSequence() []int // same there but I'm not sure if returning a slice of anys is the best soulution
	Mutate() Genotype
	Swap() Genotype
	Crossover(Genotype) (Genotype, error)
}

type GenotypeCreateLengthError struct {
	Length int
}

func (e GenotypeCreateLengthError) Error() string {
	return fmt.Sprintf("length should be greater than 0, got: %d", e.Length)
}

type GenotypesCompareLengthError struct {
	FirstGenotypeLength  int
	SecondGenotypeLength int
}

func (e GenotypesCompareLengthError) Error() string {
	return fmt.Sprintf("genotypes should have the same length, %d != %d", e.FirstGenotypeLength, e.SecondGenotypeLength)
}

type GenotypesCompareTypeError struct {
	TypeOfFirstGenotype  reflect.Type
	TypeOfSecondGenotype reflect.Type
}

func (e GenotypesCompareTypeError) Error() string {
	return fmt.Sprintf("genotypes should have the same type, %v != %v", e.TypeOfFirstGenotype, e.TypeOfSecondGenotype)
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

	for i := 0; i < len; i++ {
		bs.GenesSequence[i] = rand.Intn(2)
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

func (bs BinarySequence) Swap() Genotype {
	newSequence := make([]int, bs.Length)
	copy(newSequence, bs.GenesSequence)
	first, second := rand.Intn(bs.Length), rand.Intn(bs.Length)
	newSequence[first], newSequence[second] = newSequence[second], newSequence[first]
	return BinarySequence{newSequence, bs.Length}
}

func (bs BinarySequence) Crossover(other Genotype) (Genotype, error) {
	if reflect.TypeOf(other) != reflect.TypeOf(bs) {
		return nil, GenotypesCompareTypeError{TypeOfFirstGenotype: reflect.TypeOf(bs), TypeOfSecondGenotype: reflect.TypeOf(other)}
	}

	if bs.Length != other.(BinarySequence).Length {
		return nil, GenotypesCompareLengthError{FirstGenotypeLength: bs.Length, SecondGenotypeLength: other.(BinarySequence).Length}
	}

	index := rand.Intn(bs.Length)
	fmt.Println(index)
	resultSequence := bs.GenesSequence[:index]
	resultSequence = append(resultSequence, other.(BinarySequence).GenesSequence[index:]...)
	return BinarySequence{resultSequence, bs.Length}, nil
}

type IntSequence struct {
	GenesSequence []int
	Length        int
	MaxNumber     int
}

func GenerateIntSequence(len, maxNum int) (IntSequence, error) {
	if len <= 0 {
		return IntSequence{}, GenotypeCreateLengthError{Length: len}
	}

	is := IntSequence{make([]int, len), len, maxNum}

	for i := 0; i < len; i++ {
		is.GenesSequence[i] = rand.Intn(maxNum)
	}

	return is, nil
}

func GenerateIntSequenceWithUniqueValues(len int) (IntSequence, error) { // for now, let it be just simple permutation
	if len <= 0 {
		return IntSequence{}, GenotypeCreateLengthError{Length: len}
	}

	is := IntSequence{rand.Perm(len), len, len}

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

func (is IntSequence) Swap() Genotype {
	newSequence := make([]int, is.Length)
	copy(newSequence, is.GenesSequence)
	first, second := rand.Intn(is.Length), rand.Intn(is.Length)
	newSequence[first], newSequence[second] = newSequence[second], newSequence[first]
	return IntSequence{newSequence, is.Length, is.MaxNumber}
}

func (is IntSequence) Crossover(other Genotype) (Genotype, error) {
	if reflect.TypeOf(other) != reflect.TypeOf(is) {
		return nil, GenotypesCompareTypeError{TypeOfFirstGenotype: reflect.TypeOf(is), TypeOfSecondGenotype: reflect.TypeOf(other)}
	}

	if is.Length != other.(BinarySequence).Length {
		return nil, GenotypesCompareLengthError{FirstGenotypeLength: is.Length, SecondGenotypeLength: other.(IntSequence).Length}
	}

	index := rand.Intn(is.Length)
	fmt.Println(index)
	resultSequence := is.GenesSequence[:index]
	resultSequence = append(resultSequence, other.(BinarySequence).GenesSequence[index:]...)
	return BinarySequence{resultSequence, is.Length}, nil
}
