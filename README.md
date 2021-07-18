# RPN Calculator
![Github Actions](https://github.com/github/docs/actions/workflows/goreleaser.yml/badge.svg)
---

Some of the best calculators in the world have an 'RPN' (reverse polish notation) mode. Please implement a command-line RPN calculator as your code challenge. However, We still feel it is worth to give you a big picture of the future requirments and hoping you can take them into consideration when design your solution. Ultimately, we would like to build an online tool which is powerful and highly flexible. The work has been divided into 3 stages and is your code challenge.

## Prerequisites
* Golang > 0.15
* (Goreleaser)[https://goreleaser.com/]
* (Pre-commit)[https://pre-commit.com/]
* (Golang pre-commit hook)[https://github.com/dnephin/pre-commit-golang/]

## Running, building and testing
This package contains a `Makefile` which has three commands `build`, `clean` and `test`.

## Releasing
To create a new release just create a new `tag` and push this to Github. The Github actions will prepare a new release.

## Requirements
* The calculator has a stack that can contain real numbers.
* The calculator waits for user input and expects to receive strings containing whitespace separated lists of numbers and operators.
* Numbers are pushed on to the stack. Operators operate on numbers that are on the stack. • Available operators are +, -, *, /, sqrt, undo, clear.
* Operators pop their parameters off the stack, and push their results back onto the stack.
* The `clear` operator removes all items from the stack.
* The `undo` operator undoes the previous operation. `undo undo` will undo the previo us two operations.
* sqrt performs a square root on the top item from the stack.
* The `+`, `-`, `*`, `/` operators perform addition, subtraction, multiplication and division respectively on the top two items from the stack.
* After processing an input string, the calculator displays the current contents of the stack as a space-separated list.
* Numbers should be stored on the stack to at least 15 decimal places of precision, but displayed to 10 decimal places (or less if it causes no loss of precision).
* All numbers should be formatted as plain decimal strings (ie. no engineering formatting).
* If an operator cannot find a sufficient number of parameters on the stack, a warning is displayed: operator <operator> (position: <pos>): insufficient parameters
* After displaying the warning, all further processing of the string terminates and the current state of the stack is displayed.