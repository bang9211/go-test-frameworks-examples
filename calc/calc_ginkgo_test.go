package calc_test

import (
	"go-test-framework-example/calc"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCalc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculator Suite")
}

var _ = Describe("Calculator", func() {
	Describe("calc.Add numbers", func() {
		Context("1 and 2", func() {
			It("should be 3", func() {
				Expect(calc.Add(1, 2)).To(Equal(3))
			})
		})
	})

	Describe("calc.Subtract numbers", func() {
		Context("3 from 5", func() {
			It("should be 2", func() {
				Expect(calc.Subtract(5, 3)).To(Equal(2))
			})
		})
	})

	Describe("Multiply numbers", func() {
		Context("5 with 6", func() {
			It("should be 30", func() {
				Expect(calc.Multiply(5, 6)).To(Equal(30))
			})
		})
	})

	Describe("calc.Divide numbers", func() {
		Context("10 by 2", func() {
			It("should be 30", func() {
				Expect(calc.Divide(10, 2)).To(Equal(float64(5)))
			})
		})
	})
})
