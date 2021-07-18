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

// ClearOperator provides a simple interface to pop
// the last value of the stack.
type ClearOperator struct{}

// NewClearOperator creates a new ClearOperator instance
func NewClearOperator() ClearOperator {
	return ClearOperator{}
}

// Identifier returns the operation a user can input whilst performing
// operations on the stack.
func (o ClearOperator) Identifier() string {
	return "clear"
}

// Operate pops the last item from the stack.
func (o ClearOperator) Operate(s *stack.Stack) (err error) {
	s.Reset()
	return
}
