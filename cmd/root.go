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
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"go.awx.im/challenges/rpn-calculator/pkg/calculator"
)

// Internal pointer to the calculator
var calc *calculator.Calculator

// whether the calculator should run in interactive mode
var interactive bool

var blue = color.New(color.FgHiBlue).SprintfFunc()
var white = color.New(color.FgHiWhite).SprintfFunc()
var grey = color.New(color.FgWhite).SprintfFunc()
var green = color.New(color.FgGreen).SprintfFunc()
var red = color.New(color.FgRed).SprintfFunc()

func printOutput(c *calculator.Calculator, i string, printInput bool) {
	if err := c.Calculate(i); err != nil {
		fmt.Println(red("x"), white(err.Error()))
	} else {
		fmt.Println(green("="), white(c.Stack().String()))
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "rpn-calculator",
	Long: `Simple Reverse Polish Notation calculator`,
	// Instantiate a new calculator instance, since it is a pointer
	// we check for nil presence (which of course should not happen
	// in any case, but defensive programming here)
	// Allow the addition of any operator to the calculator. The operator
	// should implement the Operator interface which allows the calculator
	// to iterate over a the stack and apply any operator methods to its
	// aggregated result
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		calc, err = calculator.NewDefaultCalculator()
		return
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if interactive {
			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Print(blue("> "), grey(calc.Stack().String()), " ")
				input, err := reader.ReadString('\n')
				if err == nil {
					if err := calc.Calculate(input); err != nil {
						fmt.Println(red("x"), red(err.Error()))
					}
				} else {
					fmt.Println(err)
				}
			}
		} else {
			input := strings.Join(args, " ")
			fmt.Println(blue(">"), grey(input))
			printOutput(calc, input, true)
		}
		return
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	pf := rootCmd.PersistentFlags()
	pf.BoolVarP(&interactive, "interactive", "i", false, "Interactive mode")
}
