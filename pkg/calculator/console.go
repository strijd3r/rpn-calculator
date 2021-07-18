package calculator

import (
	"bufio"
	"fmt"
	"io"

	"github.com/fatih/color"
)

// define a set of colors used in console output
var (
	blue  = color.New(color.FgHiBlue).SprintfFunc()
	white = color.New(color.FgHiWhite).SprintfFunc()
	grey  = color.New(color.FgWhite).SprintfFunc()
	green = color.New(color.FgGreen).SprintfFunc()
	red   = color.New(color.FgRed).SprintfFunc()
)

// RunCalculator runs the calculator with the provided input string and
// outputs colorized information to the provided io.Writer.
func RunCalculator(c *Calculator, w io.Writer, i string) error {
	fmt.Fprintln(w, blue(">"), grey(i))
	if err := c.Calculate(i); err != nil && err != ErrUserExited {
		fmt.Fprintln(w, red("x"), white(err.Error()))
	} else {
		fmt.Fprintln(w, green("="), white(c.Stack().String()))
	}
	return nil
}

// RunInteractiveCalculator runs the calculator with the provided io.Reader until
// it hits an exit operation and outputs colorized information to the provided io.Writer.
func RunInteractiveCalculator(c *Calculator, w io.Writer, r io.Reader) error {
	reader := bufio.NewReader(r)
	for {
		o := c.Stack().String()
		if o != "" {
			o = grey(fmt.Sprintf("%s ", o))
		}
		fmt.Fprint(w, blue("> "), o)
		input, err := reader.ReadString('\n')
		if err == nil {
			if err := c.Calculate(input); err != nil {
				fmt.Fprintln(w, red("x"), red(err.Error()))
				if err == ErrUserExited {
					break
				}
			}
		} else {
			return err
		}
	}
	return nil
}
