package fieldelement

import (
	"fmt"
)

type FieldElement struct {
	num   int
	prime int
}

func (f *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%d(%d)", f.num, f.prime)
}

func NewFieldElement(num int, prime int) (*FieldElement, error) {
	if num >= prime || num < 0 {
		return nil, fmt.Errorf("num %d not in field range 0 to %d", num, prime-1)
	}
	return &FieldElement{num, prime}, nil
}

func (f *FieldElement) Equals(other *FieldElement) bool {
	if other == nil {
		return false
	}
	return f.num == other.num && f.prime == other.prime
}

// Exercise 1: Write the corresponding method __ne__, which checks if two FieldElement objects are not equal to each other.
func (f *FieldElement) NotEquals(other *FieldElement) bool {
	return !f.Equals(other)
}

func (f *FieldElement) Add(other *FieldElement) (*FieldElement, error) {
	if other == nil {
		return nil, fmt.Errorf("cannot add nil FieldElement")
	}

	if f.prime != other.prime {
		return nil, fmt.Errorf("cannot add two numbers in different Fields")
	}
	num := (f.num + other.num) % f.prime
	return NewFieldElement(num, f.prime)
}

func (f *FieldElement) Sub(other *FieldElement) (*FieldElement, error) {
	if other == nil {
		return nil, fmt.Errorf("cannot subtract nil FieldElement")
	}

	if f.prime != other.prime {
		return nil, fmt.Errorf("cannot subtract two numbers in different Fields")
	}
	num := (f.num - other.num + f.prime) % f.prime
	return NewFieldElement(num, f.prime)
}
