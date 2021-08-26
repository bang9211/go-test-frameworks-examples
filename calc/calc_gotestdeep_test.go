package calc_test

import (
	"go-test-framework-example/calc"
	"testing"

	. "github.com/maxatome/go-testdeep"
)

func TestAddWithGoTestDeep(t *testing.T) {
	result := calc.Add(1, 2)
	CmpNotZero(t, result)
	CmpDeeply(t, &result, Ptr(3))
	CmpDeeply(t, result, Code(func(r int) (bool, string) {
		if r == 3 {
			return true, ""
		}
		return false, "Result should be 3"
	}))
}

func TestSubtractWithGoTestDeep(t *testing.T) {
	result := calc.Subtract(5, 3)
	CmpNotZero(t, result)
	CmpDeeply(t, &result, Ptr(2))
	CmpDeeply(t, result, Code(func(r int) (bool, string) {
		if r == 2 {
			return true, ""
		}
		return false, "Result should be 2"
	}))
}

func TestMultiplyWithGoTestDeep(t *testing.T) {
	result := calc.Multiply(5, 6)
	CmpNotZero(t, result)
	CmpDeeply(t, &result, Ptr(30))
	CmpDeeply(t, result, Code(func(r int) (bool, string) {
		if r == 30 {
			return true, ""
		}
		return false, "Result should be 30"
	}))
}

func TestDivideWithGoTestDeep(t *testing.T) {
	result := calc.Divide(10, 2)
	CmpNotZero(t, result)
	CmpDeeply(t, &result, Ptr(float64(5)))
	CmpDeeply(t, result, Code(func(r float64) (bool, string) {
		if r == float64(5) {
			return true, ""
		}
		return false, "Result should be 5"
	}))
}
