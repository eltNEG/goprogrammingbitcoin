package ch01_test

import (
	"testing"

	fieldelement "github.com/eltNEG/goprogrammingbitcoin/ch01/src"
)

func TestFiniteField(t *testing.T) {
	testCases := []struct {
		num    int
		prime  int
		expect string
		hasErr bool
	}{
		{1, 2, "FieldElement_1(2)", false},
		{2, 2, "FieldElement_2(2)", true},
		{2, 8, "FieldElement_2(8)", false},
		{-1, 5, "", true},
		{5, 5, "", true},
		{6, 5, "", true},
	}

	for _, tc := range testCases {
		f, err := fieldelement.NewFieldElement(tc.num, tc.prime)
		if err != nil {
			if !tc.hasErr {
				t.Errorf("Unexpected error: %s", err)
			}
			continue
		}
		if f.String() != tc.expect {
			t.Errorf("Expected %s, got %s", tc.expect, f.String())
		}
	}
}

func CreateFieldElement(num int, prime int) *fieldelement.FieldElement {
	f, _ := fieldelement.NewFieldElement(num, prime)
	return f
}

func TestFieldElementEquals(t *testing.T) {
	testCases := []struct {
		f1     *fieldelement.FieldElement
		f2     *fieldelement.FieldElement
		expect bool
	}{
		{CreateFieldElement(1, 2), CreateFieldElement(1, 2), true},
		{CreateFieldElement(1, 2), CreateFieldElement(2, 2), false},
		{CreateFieldElement(1, 2), CreateFieldElement(1, 3), false},
		{CreateFieldElement(1, 2), nil, false},
	}

	for _, tc := range testCases {
		if tc.f1.Equals(tc.f2) != tc.expect {
			t.Errorf("Expected %t, got %t", tc.expect, tc.f1.Equals(tc.f2))
		}
	}

	// test not equals
	for _, tc := range testCases {
		if tc.f1.NotEquals(tc.f2) == tc.expect {
			t.Errorf("Expected %t, got %t", !tc.expect, tc.f1.NotEquals(tc.f2))
		}
	}
}

func TestFieldElementAdd(t *testing.T) {
	testCases := []struct {
		f1     *fieldelement.FieldElement
		f2     *fieldelement.FieldElement
		expect string
		hasErr bool
	}{
		{CreateFieldElement(1, 2), CreateFieldElement(1, 2), "FieldElement_0(2)", false},
		{CreateFieldElement(2, 3), CreateFieldElement(2, 3), "FieldElement_1(3)", false},
		{CreateFieldElement(1, 2), CreateFieldElement(3, 2), "", true},
		{CreateFieldElement(1, 3), CreateFieldElement(2, 2), "", true},
	}

	for _, tc := range testCases {
		f, err := tc.f1.Add(tc.f2)
		if err != nil {
			if !tc.hasErr {
				t.Errorf("Unexpected error: %s", err)
			}
			continue
		}
		if f.String() != tc.expect {
			t.Errorf("Expected %s, got %s", tc.expect, f.String())
		}
	}
}

/**
Exercise 2: Solve these problems in F57 (assume all +’s here are +f and –’s here are –f):
F(59)
• 44 + 33
• 9 – 29
• 17 + 42 + 49
• 52 – 30 – 38
*/

func TestExercise2(t *testing.T) {
	testcases := []struct {
		prime   int
		numbers []int
		expect  string
	}{
		{13, []int{7, 12}, "FieldElement_6(13)"},
		{19, []int{9, 10}, "FieldElement_0(19)"},
		{19, []int{11, -9}, "FieldElement_2(19)"},
		{19, []int{6, -13}, "FieldElement_12(19)"},
		{59, []int{44, 33}, "FieldElement_18(59)"},
		{59, []int{9, -29}, "FieldElement_39(59)"},
		{59, []int{17, 42, 49}, "FieldElement_49(59)"},
		{59, []int{52, -30, -38}, "FieldElement_43(59)"},
	}

	for _, tc := range testcases {
		var f *fieldelement.FieldElement
		var err error
		for _, n := range tc.numbers {
			if f == nil {
				f, _ = fieldelement.NewFieldElement(n, tc.prime)
				continue
			}
			if n < 0 {
				f, err = f.Sub(CreateFieldElement(-1*n, tc.prime))
			} else {
				f, err = f.Add(CreateFieldElement(n, tc.prime))
			}

			if err != nil {
				t.Errorf("Unexpected error: %s", err)
				break
			}
		}
		if f.String() != tc.expect {
			t.Errorf("Expected %s, got %s", tc.expect, f.String())
		}
	}
}
