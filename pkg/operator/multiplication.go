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

import (
	"go.awx.im/challenges/rpn-calculator/pkg/stack"
)

// MultiplicationOperator provides a simple interface to compute the
// multiple of the last two values on the stack.
type MultiplicationOperator struct {
	ArithmeticOperator
}

// NewMultiplicationOperator is a factory method to instantiate a new
// MultiplicationOperator with a default ArtithmeticOperator.
func NewMultiplicationOperator() MultiplicationOperator {
	return MultiplicationOperator{NewArithmeticOperator()}
}

// Identifier returns the operation a user can input whilst performing
// arithmetic operations on the stack.
func (o MultiplicationOperator) Identifier() string {
	return "*"
}

// Operate gets the last two entries of the stack and performs multiplication
// on it.
func (o MultiplicationOperator) Operate(s *stack.Stack) (err error) {
	var a, b float64
	if a, b, err = o.Get(s); err == nil {
		s.Push(a * b)
	}
	return
}
