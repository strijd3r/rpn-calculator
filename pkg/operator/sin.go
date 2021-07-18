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
	"math"

	"go.awx.im/challenges/rpn-calculator/pkg/stack"
)

// SqrtOperator provides a simple interface to compute the
// sinoid value of the last value on the stack
type SinOperator struct{}

// NewSinOperator is a factory method to instantiate a new
// SinOperator.
func NewSinOperator() SinOperator {
	return SinOperator{}
}

// Identifier returns the operation a user can input whilst performing
// arithmetic operations on the stack.
func (o SinOperator) Identifier() string {
	return "sin"
}

// Operate gets the last entry of the stack and performs math.Sin on it.
func (o SinOperator) Operate(s *stack.Stack) (err error) {
	var a float64
	if a, err = s.Pop(); err != nil {
		return
	}
	s.Push(math.Sin(a))
	return
}
