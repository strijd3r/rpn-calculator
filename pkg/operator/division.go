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

// DivisionOperator is an implementation of the operator.Operator
// interface which performs the division of the last 2 numbers
// on the stack.
type DivisionOperator struct {
	ArithmeticOperator
}

// NewDivisionOperator is a factory method to construct a new
// DisivionOperator instance.
func NewDivisionOperator() DivisionOperator {
	return DivisionOperator{NewArithmeticOperator()}
}

// Identifier returns the operation a user can input whilst performing
// arithmetic operations on the stack
func (o DivisionOperator) Identifier() string {
	return "/"
}

// Operate divides two numbers on the stack, it leverages the ArtithmeticOperator
// to easily define the last two digits and sets the return value.
func (o DivisionOperator) Operate(s *stack.Stack) (err error) {
	var a, b float64
	if a, b, err = o.Get(s); err == nil {
		s.Push(a / b)
	}
	return
}
