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

	"go.awx.im/challenges/rpn-calculator/pkg/operator"
	"go.awx.im/challenges/rpn-calculator/pkg/stack"
)

var (
	// ErrInvalidOperator
	ErrInvalidOperator = errors.New("invalid operator specific")
)

// Calculator is
type Calculator struct {
	s *stack.Stack
	o map[string]operator.Operator
}

// NewDefaultCalculator is
func NewDefaultCalculator() (*Calculator, error) {
	return NewCalculator(
		operator.NewAdditionOperator(),
		operator.NewMultiplicationOperator(),
		operator.NewSubtractionOperator(),
		operator.NewDivisionOperator(),
		operator.NewSqrtOperator(),
		operator.NewSinOperator(),
		operator.NewUndoOperator(),
	)
}

// NewCalculator
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

// Calculate is
func (c *Calculator) Calculate(input string) (err error) {
	var scan scanner.Scanner
	scan.Init(strings.NewReader(input))

	for tok := scan.Scan(); tok != scanner.EOF; tok = scan.Scan() {
		t := scan.TokenText()
		val, ok := c.o[t]
		if ok {
			if err = val.Operate(c.s); err != nil {
				return fmt.Errorf("operator %s (position: %d): insufficient parameters", t, scan.Position.Offset)
			}
		} else {
			if err = c.s.PushString(t); err != nil {
				return fmt.Errorf("operator %s (position: %d): invalid operator", t, scan.Position.Offset)
			}
		}
	}
	return
}
