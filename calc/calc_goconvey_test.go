package calc_test

import (
	"go-test-framework-example/calc"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddWithGoConvey(t *testing.T) {
	Convey("Adding two numbers", t, func() {
		x := 1
		y := 2

		Convey("should produce the expected result", func() {
			So(calc.Add(x, y), ShouldEqual, 3)
		})
	})
}

func TestSubtractWithGoConvey(t *testing.T) {
	Convey("Subtracting two numbers", t, func() {
		x := 5
		y := 3

		Convey("should produce the expected result", func() {
			So(calc.Subtract(x, y), ShouldEqual, 2)
		})
	})
}

func TestMultiplyWithGoConvey(t *testing.T) {
	Convey("Multiplying two numbers", t, func() {
		x := 5
		y := 6

		Convey("should produce the expected result", func() {
			So(calc.Multiply(x, y), ShouldEqual, 30)
		})
	})
}

func TestDivideWithGoConvey(t *testing.T) {
	Convey("Dividing two numbers", t, func() {
		x := 10
		y := 2

		Convey("should produce the expected result", func() {
			So(calc.Divide(x, y), ShouldEqual, float64(5))
		})
	})
}
