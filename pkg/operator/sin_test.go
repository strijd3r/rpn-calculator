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
package operator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/strijd3r/rpn-calculator/pkg/operator"
	"github.com/strijd3r/rpn-calculator/pkg/stack"
)

var _ = Describe("SinOperator", func() {
	var o operator.SinOperator
	var s *stack.Stack
	var err error
	BeforeEach(func() {
		o = operator.NewSinOperator()
		s = stack.NewStack()
	})

	Context("Identifier() string", func() {
		It("should return sin", func() {
			Expect(o.Identifier()).To(Equal("sin"))
		})
	})

	Context("Operate(*stack.Stack) error", func() {
		When("The stack has less than 1 entry", func() {
			It("should return an error", func() {
				Expect(o.Operate(s)).To(HaveOccurred())
			})
		})

		When("The stack has more than 1 entry", func() {
			var f float64
			BeforeEach(func() {
				s.Push(1, 2, 3)
			})
			It("should replace the last entry with the calculated result", func() {
				Expect(o.Operate(s)).NotTo(HaveOccurred())
				Expect(s.Size()).To(Equal(3))
				f, err = s.Pop()
				Expect(f).To(Equal(0.1411200080598672))
				Expect(err).NotTo(HaveOccurred())

				f, err = s.Pop()
				Expect(f).To(Equal(2.0))
				Expect(err).NotTo(HaveOccurred())

				f, err = s.Pop()
				Expect(f).To(Equal(1.0))
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})
