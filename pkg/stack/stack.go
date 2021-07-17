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
package stack

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"sync"
)

// ErrStackEmpty
var ErrStackEmpty = errors.New("stack: is empty")

// Stack is
type Stack struct {
	mu sync.Mutex

	stack []float64
}

// NewStack is the factory method to instantiate a new Stack
func NewStack() *Stack {
	return &Stack{}
}

// Print is a helper method to print the current stack in
// human readable format to the provided io.Writer interface.
func (s *Stack) Print(output io.Writer) {
	for i, v := range s.stack {
		if i != 0 {
			fmt.Fprint(output, " ")
		}
		fmt.Fprint(output, v)
	}
}

// Size returns the length of the current stack
func (s *Stack) Size() int {
	return len(s.stack)
}

// Push pushes a float64 to the end of the stack
func (s *Stack) Push(f ...float64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.stack = append(s.stack, f...)
}

// Push pushes a string to the end of the stack
func (s *Stack) PushString(str ...string) (err error) {
	var values []float64
	for _, v := range str {
		var f float64
		if f, err = strconv.ParseFloat(v, 64); err != nil {
			return
		}
		values = append(values, f)
	}
	s.Push(values...)
	return
}

// Pop pops a value from the end of the stack and returns it.
// If the stack is empty it will raise an error
func (s *Stack) Pop() (f float64, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var i, l int

	l = s.Size()
	if l == 0 {
		return 0, ErrStackEmpty
	}

	i = l - 1
	f = s.stack[i]
	s.stack = s.stack[:i]
	return
}
