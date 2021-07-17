/*
Copyright © 2021 Unknown <applicant@airwallex.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package calculator_test

import (
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"go.awx.im/challenges/rpn-calculator/pkg/calculator"
	"go.awx.im/challenges/rpn-calculator/pkg/operator"
	"go.awx.im/challenges/rpn-calculator/pkg/stack"
)

type InvalidOperator struct{}

func (o *InvalidOperator) Identifier() string {
	return ""
}

func (o *InvalidOperator) Operate(s *stack.Stack) error {
	return nil
}

var _ = Describe("Calculator", func() {
	var tests map[string]string
	var c *calculator.Calculator
	var err error

	BeforeEach(func() {
		c, err = calculator.NewDefaultCalculator()
		tests = map[string]string{}
		Expect(c).NotTo(BeNil())
		Expect(err).NotTo(HaveOccurred())
	})

	Context("Stack() *stack.Stack", func() {
		It("should return the stack of the calculator", func() {
			Expect(c.Stack()).NotTo(BeNil())
		})
	})

	Context("NewCalculator()", func() {
		When("an operator is added", func() {
			It("should not return an error", func() {
				c, err = calculator.NewCalculator(operator.NewAdditionOperator())
				Expect(err).NotTo(HaveOccurred())
				Expect(c).NotTo(BeNil())
			})
		})

		When("and invalid operator is passed", func() {
			It("should return an error", func() {
				c, err = calculator.NewCalculator(&InvalidOperator{})
				Expect(err).To(HaveOccurred())
			})
		})

		When("a duplicate operator is passed", func() {
			It("should return an error", func() {
				o := operator.NewAdditionOperator()
				c, err = calculator.NewCalculator(o, o)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Context("Calculate(string) error", func() {
		When("there are no errors", func() {
			tests = map[string]string{
				"":                      "",
				"1 2 3 * *":             "6",
				"2 2 3 * *":             "12",
				"3 2 3 * *":             "18",
				"12 3 * 12 24 + * sqrt": "36",
			}
			for k, v := range tests {
				It("should calculate the stack based on an input string", func() {
					Expect(c.Calculate(k)).NotTo(HaveOccurred())

					b := &bytes.Buffer{}
					c.Stack().Print(b)
					Expect(b.String()).To(Equal(v))
				})
			}
		})

		When("there are errors", func() {
			It("should calculate and raise an error with a message", func() {
				tests = map[string]string{
					"invalid":                 "operator invalid (position: 0): invalid operator",
					"12 3 * 12 24 + * sqrt *": "operator * (position: 22): insufficient parameters",
				}
				for k, v := range tests {
					e := c.Calculate(k)
					Expect(e).To(HaveOccurred())
					Expect(e.Error()).To(Equal(v))
				}
			})
		})
	})
})
