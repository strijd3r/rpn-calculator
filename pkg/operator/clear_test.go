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

var _ = Describe("ClearOperator", func() {
	var o operator.ClearOperator
	var s *stack.Stack
	BeforeEach(func() {
		o = operator.NewClearOperator()
		s = stack.NewStack()
	})

	Context("Identifier() string", func() {
		It("should return clear", func() {
			Expect(o.Identifier()).To(Equal("clear"))
		})
	})

	Context("Operate(*stack.Stack) error", func() {
		It("should not error when the stack is empty", func() {
			Expect(o.Operate(s)).NotTo(HaveOccurred())
			Expect(s.Size()).To(Equal(0))
		})

		It("should clear the stack when its not empty", func() {
			s.Push(1, 2, 3)
			Expect(s.Size()).To(Equal(3))
			Expect(o.Operate(s)).NotTo(HaveOccurred())
			Expect(s.Size()).To(Equal(0))
		})
	})
})
