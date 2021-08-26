package calc_test

import (
	"go-test-framework-example/calc"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestAddWithGocheck(c *C) {
	result := calc.Add(1, 2)
	c.Assert(result, Equals, 3)
}

func (s *MySuite) TestSubtractWithGocheck(c *C) {
	result := calc.Subtract(5, 3)
	c.Assert(result, Equals, 2)
}

func (s *MySuite) TestMultiplyWithGocheck(c *C) {
	result := calc.Multiply(5, 6)
	c.Assert(result, Equals, 30)
}

func (s *MySuite) TestDivideWithGocheck(c *C) {
	result := calc.Divide(10, 2)
	c.Assert(result, Equals, float64(5))
}
