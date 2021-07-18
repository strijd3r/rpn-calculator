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
)

var _ = Describe("Calculator", func() {
	var i *bytes.Buffer
	var o *bytes.Buffer
	var c *calculator.Calculator
	var err error

	BeforeEach(func() {
		i = &bytes.Buffer{}
		o = &bytes.Buffer{}
		c, err = calculator.NewDefaultCalculator()
		Expect(c).NotTo(BeNil())
		Expect(err).NotTo(HaveOccurred())
	})

	Context("RunCalculator(*Calculator, io.Writer, string) error", func() {
		It("should not error", func() {
			err = calculator.RunCalculator(c, o, "1 2 +")
			Expect(err).NotTo(HaveOccurred())

			err = calculator.RunCalculator(c, o, "1 2 exit")
			Expect(err).NotTo(HaveOccurred())
		})

		It("should print an error when the input is invalid", func() {
			err = calculator.RunCalculator(c, o, "1 2 invalid")
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("RunInteractiveCalculator(*Calculator, io.Writer, io.Reader) error", func() {
		It("should error when the input string is not terminated with a newline", func() {
			i.Write([]byte("1 2 exit"))
			err = calculator.RunInteractiveCalculator(c, o, i)
			Expect(err).To(HaveOccurred())
		})

		It("should not error", func() {
			i.Write([]byte("1 2 +\n"))
			i.Write([]byte("1 2 exit\n"))
			err = calculator.RunInteractiveCalculator(c, o, i)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
