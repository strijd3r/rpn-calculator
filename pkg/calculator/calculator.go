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
package calculator

import (
	"errors"
	"fmt"
	"strings"
	"text/scanner"

	"github.com/strijd3r/rpn-calculator/pkg/operator"
	"github.com/strijd3r/rpn-calculator/pkg/stack"
)

// define a special exist operator
const exitOperator = "exit"

// ErrInvalidOperator is a predefined error for invalid operators.
var ErrInvalidOperator = errors.New("invalid operator specific")

// ErrUserExited is a hacky way to exit the current operation
var ErrUserExited = errors.New("user exited the operation")

// Calculator is the main application which allows parsing input
// and adds item to the stack or performs operations on the stack.
type Calculator struct {
	s *stack.Stack
	o map[string]operator.Operator
}

// NewDefaultCalculator is a factory method to return a new Calculator
// instance with all default operators.
func NewDefaultCalculator() (*Calculator, error) {
	return NewCalculator(
		operator.NewAdditionOperator(),
		operator.NewDivisionOperator(),
		operator.NewMultiplicationOperator(),
		operator.NewSinOperator(),
		operator.NewSqrtOperator(),
		operator.NewSubtractionOperator(),
		operator.NewClearOperator(),
		operator.NewUndoOperator(),
	)
}

// NewCalculator is a factory method to create a new Calculator instance and
// adds the provided operators to its internal map. It detects duplicate
// operator types and invalid operators which do not define any operation.
func NewCalculator(operators ...operator.Operator) (c *Calculator, err error) {
	c = &Calculator{
		s: stack.NewStack(),
		o: map[string]operator.Operator{},
	}
	for _, operator := range operators {
		var operation = operator.Identifier()
		if operation == "" {
			err = ErrInvalidOperator
		} else if c.o[operation] != nil {
			err = fmt.Errorf("operator %s: duplicate operator detected", operation)
		} else {
			c.o[operation] = operator
		}
	}
	return
}

// Stack returns the current stack
func (c *Calculator) Stack() *stack.Stack {
	return c.s
}

// Calculate receives an input string, tokenizes it and checks whether the input is
// a) a valid operator or b) a valid float64 value.
func (c *Calculator) Calculate(input string) (err error) {
	var scan scanner.Scanner
	scan.Init(strings.NewReader(input))

	for tok := scan.Scan(); tok != scanner.EOF; tok = scan.Scan() {
		t := scan.TokenText()
		if t == exitOperator {
			return ErrUserExited
		}
		val, ok := c.o[t]
		if ok {
			if err = val.Operate(c.s); err != nil {
				return fmt.Errorf("operator %s (position: %d): insufficient parameters", t, scan.Position.Offset+1)
			}
		} else {
			if err = c.s.PushString(t); err != nil {
				return fmt.Errorf("operator %s (position: %d): invalid operator", t, scan.Position.Offset)
			}
		}
	}
	return
}
