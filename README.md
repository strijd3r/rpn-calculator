# RPN Calculator
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/strijd3r/rpn-calculator)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/strijd3r/rpn-calculator)
[![Github Actions](https://github.com/strijd3r/rpn-calculator/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/strijd3r/rpn-calculator/actions/workflows/goreleaser.yml)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/strijd3r/rpn-calculator)
![Languages](https://img.shields.io/github/languages/count/strijd3r/rpn-calculator)
![Top language](https://img.shields.io/github/languages/top/strijd3r/rpn-calculator)
---

## Introduction
In reverse Polish notation, the operators follow their operands; for instance, to add 3 and 4 together, one would write 3 4 + rather than 3 + 4. If there are multiple operations, operators are given immediately after their second operands; so the expression written 3 − 4 + 5 in conventional notation would be written 3 4 − 5 + in reverse Polish notation: 4 is first subtracted from 3, then 5 is added to it. An advantage of reverse Polish notation is that it removes the need for parentheses that are required by infix notation. While 3 − 4 × 5 can also be written 3 − (4 × 5), that means something quite different from (3 − 4) × 5. In reverse Polish notation, the former could be written 3 4 5 × −, which unambiguously means 3 (4 5 ×) − which reduces to 3 20 − (which can further be reduced to -17); the latter could be written 3 4 − 5 × (or 5 3 4 − ×, if keeping similar formatting), which unambiguously means (3 4 −) 5 ×.

## Table of contents
- [RPN Calculator](#rpn-calculator)
  - [!Top language](#)
  - [Introduction](#introduction)
  - [Table of contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
  - [Running, building and testing](#running-building-and-testing)
    - [Building a new binary](#building-a-new-binary)
    - [Running the calculator](#running-the-calculator)
    - [Installing dependencies](#installing-dependencies)
    - [Running the tests](#running-the-tests)
  - [Releasing](#releasing)
  - [Requirements](#requirements)
  - [Addendum](#addendum)
  
## Prerequisites
* [Golang](https://golang.org) > 0.15
* [Goreleaser](https://goreleaser.com/)
* [Pre-commit](https://pre-commit.com/)
* [Golang pre-commit hook](https://github.com/dnephin/pre-commit-golang/)

## Running, building and testing
This package contains a `Makefile` which has commands such as `build`, `deps`, `clean` and `test`.

### Building a new binary
Checkout the repository and run `make build`. The produced binary can be found in `dist/calculator`

### Running the calculator
The calculator can be run in multiple ways, when running it from a produced binary use the provided `rpn-calculator` binary. For development you can run `go run main.go` with additional arguments.

**Interactive**
Run `rpn-calculator -i` or `go run main.go -i` to run in interactive shell mode

**Command line input**
Run `rpn-calculator 1 2 "*" 4 \` or `go run main.go 1 2 "*" 4 \`  to calculate the given input, beware to quote shell arguments like `*`

### Installing dependencies
This project requires you to have the `ginkgo` binary installed prior to running tests. Ginkgo is a BDD-style testing library which simplifies
testing whilst allow to write languages in natural language.

### Running the tests
Tests can be run with Ginkgo by executing `make test`. This will run Ginkgo and output coverprofiles that can be parsed by any 3rd party system.

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

## Addendum
* The interactive console has been adapted with an improved interface allowing to easily calculate with speed.
* Use of `float64` suffices to the specs, to allow further precision the system has to be refactored to support `complex128` types.

