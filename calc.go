package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	rpn := false
	expression := ""
	for _, arg := range os.Args[1:] {
		if arg == "-r" || arg == "--rpn" {
			rpn = true
		} else if arg == "-h" || arg == "--help" {
			help()
			os.Exit(0)
		} else {
			expression += " "
			expression += arg
		}
	}
	if len(expression) == 0 {
		expression = readStdin()
	}
	result, err := calculate(expression, rpn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.15G\n", result)
}

func calculate(expression string, rpn bool) (float64, error) {
	tokens, err := Tokenize(expression)
	if err != nil {
		return 0, err
	}
	if !rpn {
		tokens, err = ConvertInfixToPostfix(tokens)
		if err != nil {
			return 0, err
		}
	}
	value, err := Calculate(tokens)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func readStdin() string {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return string(stdin)
}

func help() {
	print("\n" +
		"calc -- a simple command-line calculator\n" +
		"\n" +
		"usage: calc [options] expression\n" +
		"\n" +
		"Options:\n" +
		"\n" +
		"  -h, --help  show this help\n" +
		"  -r, --rpn   use \"Reverse Polish Notation\" (postfix)\n" +
		"\n" +
		"Operators: - * / % ^\n" +
		"Functions: abs, acos, asin, atan, cos, cosh, exp, ln, log, neg,\n" +
		"           round, sin, sinh, sqrt, tan, tanh, trunc\n" +
		"Constants: e, pi\n" +
		"\n" +
		"For default infix expressions, function arguments must be given\n" +
		"in parenthesis. For RPN, parenthesis are illegal.\n" +
		"\n" +
		"Examples:\n" +
		"  calc \"sin(3.1415926)\"\n" +
		"  calc \"(5 3) * 7\"\n" +
		"  calc \"2^3\"\n" +
		"  calc -r \"pi sin\"\n" +
		"  calc -r \"5 3 7 *\"\n" +
		"  calc -r \"2 3 ^\"\n" +
		"  for Unix sh: A=`calc \"3+1\"`; B=`calc \"$A*4\"`\n")
}
