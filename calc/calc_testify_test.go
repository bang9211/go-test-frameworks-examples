package calc_test

import (
	"go-test-framework-example/calc"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestAddWithTestify(t *testing.T) {
	result := calc.Add(1, 2)
	Equal(t, 3, result)
}

func TestSubtractWithTestify(t *testing.T) {
	result := calc.Subtract(5, 3)
	Equal(t, 2, result)
}

func TestMultiplyWithTestify(t *testing.T) {
	result := calc.Multiply(5, 6)
	Equal(t, 30, result)
}

func TestDivideWithTestify(t *testing.T) {
	result := calc.Divide(10, 2)
	Equal(t, float64(5), result)
}
