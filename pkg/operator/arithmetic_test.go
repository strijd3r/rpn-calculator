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

var _ = Describe("ArithmeticOperator", func() {
	var o operator.ArithmeticOperator
	var s *stack.Stack
	var err error
	BeforeEach(func() {
		o = operator.NewArithmeticOperator()
		s = stack.NewStack()
	})

	var a, b float64
	Context("Get(*stack.Stack) (float64, float64, error)", func() {
		When("The stack has less than 2 entries", func() {
			It("should return an error", func() {
				a, b, err = o.Get(s)
				Expect(err).To(HaveOccurred())
			})

			It("should push an eventual popped item from to the stack", func() {
				s.Push(1.0)
				a, b, err = o.Get(s)
				Expect(err).To(HaveOccurred())
				Expect(s.Size()).To(Equal(1))
				Expect(b).To(Equal(0.0))

				a, err = s.Pop()
				Expect(err).NotTo(HaveOccurred())
				Expect(a).To(Equal(1.0))
			})
		})

		When("The stack has more than 2 entries", func() {
			It("should return the last two entries", func() {
				s.Push(1, 2)

				a, b, err = o.Get(s)
				Expect(err).NotTo(HaveOccurred())
				Expect(a).To(Equal(1.0))
				Expect(b).To(Equal(2.0))
			})
		})
	})
})
