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
package stack_test

import (
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/strijd3r/rpn-calculator/pkg/stack"
)

var _ = Describe("Stack", func() {
	var s *stack.Stack
	BeforeEach(func() {
		s = stack.NewStack()
	})

	Context("Push(...float64)", func() {
		When("there is only one value push", func() {
			It("has a size of 1", func() {
				s.Push(0)
				Expect(s.Size()).To(Equal(1))
			})
		})

		When("there are multiple values pushed", func() {
			It("appends all values to the stack", func() {
				s.Push(0, 1, 2, 3)
				Expect(s.Size()).To(Equal(4))
			})
		})
	})

	Context("Pop() (float64, error)", func() {
		When("the stack is empty", func() {
			It("will raise an error", func() {
				f, err := s.Pop()
				Expect(f).To(Equal(0.0))
				Expect(err).To(HaveOccurred())
			})
		})

		When("the stack is not empty", func() {
			It("should return the last value added to the stack", func() {
				var v = []float64{0, 1, 2, 3}
				var f float64
				var err error

				s.Push(v...)
				for i := len(v) - 1; i >= 0; i-- {
					f, err = s.Pop()
					Expect(f).To(Equal(v[i]))
					Expect(err).NotTo(HaveOccurred())
				}
			})
		})
	})

	Context("PushString(...string) error", func() {
		When("there is only one value push", func() {
			It("has a size of 1", func() {
				err := s.PushString("0")
				Expect(s.Size()).To(Equal(1))
				Expect(err).NotTo(HaveOccurred())
			})
		})

		When("a value could not be cast as string", func() {
			It("should return an error", func() {
				err := s.PushString("0", "invalid")
				Expect(s.Size()).To(Equal(0))
				Expect(err).To(HaveOccurred())
			})
		})

		When("there are multiple values pushed", func() {
			It("appends all values to the stack", func() {
				var v = []string{"0", "1", "2", "3"}
				var r = []float64{0, 1, 2, 3}
				var f float64
				var err error

				err = s.PushString(v...)
				Expect(err).NotTo(HaveOccurred())

				for i := len(r) - 1; i >= 0; i-- {
					f, err = s.Pop()
					Expect(f).To(Equal(r[i]))
					Expect(err).NotTo(HaveOccurred())
				}
			})
		})
	})

	Context("Size() int", func() {
		When("the stack has values", func() {
			It("should print the current stack", func() {
				s.Push(0, 1, 2, 3)
				Expect(s.Size()).To(Equal(4))
			})
		})

		When("the stack is empty", func() {
			It("should return 0", func() {
				Expect(s.Size()).To(Equal(0))
			})
		})
	})

	Context("Reset()", func() {
		It("should clear the stack when its not empty", func() {
			Expect(s.Size()).To(Equal(0))
			s.Push(1, 2, 3)
			Expect(s.Size()).To(Equal(3))
			s.Reset()
			Expect(s.Size()).To(Equal(0))
		})
	})

	// @todo: write test for rounding decimal places
	Context("Print(io.Writer)", func() {
		var b *bytes.Buffer
		BeforeEach(func() {
			b = new(bytes.Buffer)
		})

		When("the stack has values", func() {
			It("should print the current stack", func() {
				s.Push(0, 1, 2, 3)

				s.Print(b)
				Expect(b.String()).To(Equal("0 1 2 3"))
			})
		})

		When("the stack is empty", func() {
			It("should return an empty string", func() {
				s.Print(b)
				Expect(b.String()).To(Equal(""))
			})
		})
	})

	Context("String() string", func() {
		It("should return the computed stack represented as a string", func() {
			Expect(s.String()).To(Equal(""))
		})
	})
})
