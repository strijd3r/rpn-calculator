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
package operator

import "github.com/strijd3r/rpn-calculator/pkg/stack"

// ArithmeticOperator is a base struct which implements some
// helper functions to perform common operations on the stack.
// FYI: This is only here to demonstrate inheritance in Golang
// since the implementation is so simple, the Get method should
// just be a private method in this package
type ArithmeticOperator struct{}

// NewArithmeticOperator creates a new base implementation for the arithmetic
// operator
func NewArithmeticOperator() ArithmeticOperator {
	return ArithmeticOperator{}
}

// Get retrieves the last two values of the stack, or an error
// when the stack is too small.
func (o ArithmeticOperator) Get(s *stack.Stack) (a float64, b float64, err error) {
	if b, err = s.Pop(); err != nil {
		return
	}
	if a, err = s.Pop(); err != nil {
		s.Push(b)
		return
	}
	return
}
