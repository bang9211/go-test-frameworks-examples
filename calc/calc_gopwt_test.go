package calc_test

import (
	"flag"
	"go-test-framework-example/calc"
	"os"
	"testing"

	. "github.com/ToQoz/gopwt"
	. "github.com/ToQoz/gopwt/assert"
)

func TestMain(m *testing.M) {
	flag.Parse()
	Empower()
	os.Exit(m.Run())
}

func TestAddWithGopwt(t *testing.T) {
	result := calc.Add(1, 2)
	OK(t, 3 == result)
}

func TestSubtractWithGopwt(t *testing.T) {
	result := calc.Subtract(5, 3)
	OK(t, 2 == result)
}

func TestMultiplyWithGopwt(t *testing.T) {
	result := calc.Multiply(5, 6)
	OK(t, 30 == result)
}

func TestDivideWithGopwt(t *testing.T) {
	result := calc.Divide(10, 2)
	OK(t, float64(5) == result)
}
