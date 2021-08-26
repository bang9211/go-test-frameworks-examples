package calc_test

import (
	"go-test-framework-example/calc"
	"testing"

	. "github.com/franela/goblin"
)

func TestCalculator(t *testing.T) {
	g := Goblin(t)
	g.Describe("Calculator", func() {
		g.It("should add two numbers ", func() {
			g.Assert(calc.Add(1, 2)).Equal(3)
		})

		g.It("should subtract two numbers", func() {
			g.Assert(calc.Subtract(5, 3)).Equal(2)
		})

		g.It("should multiply two numbers", func() {
			g.Assert(calc.Multiply(5, 6)).Equal(30)
		})

		g.It("should divide two numbers", func() {
			g.Assert(calc.Divide(10, 2)).Equal(float64(5))
		})
	})
}
