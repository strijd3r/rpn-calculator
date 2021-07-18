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
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/strijd3r/rpn-calculator/pkg/calculator"
)

// Internal pointer to the calculator
var calc *calculator.Calculator

// whether the calculator should run in interactive mode
var interactive bool

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
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		calc, err = calculator.NewDefaultCalculator()
		return
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if interactive {
			return calculator.RunInteractiveCalculator(calc, os.Stdout, os.Stdin)
		}
		return calculator.RunCalculator(calc, os.Stdout, strings.Join(args, " "))
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
