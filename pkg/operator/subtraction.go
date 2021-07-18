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

import "go.awx.im/challenges/rpn-calculator/pkg/stack"

// SubtractionOperator provides a simple interface to compute two
// numbers in the stack.
type SubtractionOperator struct {
	ArithmeticOperator
}

// NewSubtractionOperator is a factory method to instantiate a new
// SubtractionOperator seeded with a default ArithmeticOpertor.
func NewSubtractionOperator() SubtractionOperator {
	return SubtractionOperator{NewArithmeticOperator()}
}

// Identifier returns the operation a user can input whilst performing
// arithmetic operations on the stack
func (o SubtractionOperator) Identifier() string {
	return "-"
}

// Operate gets the last two entries from the stack and then performs
// the subtraction of both entries.
func (o SubtractionOperator) Operate(s *stack.Stack) (err error) {
	var a, b float64
	if a, b, err = o.Get(s); err == nil {
		s.Push(a - b)
	}
	return
}
