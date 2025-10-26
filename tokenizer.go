package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"unicode"
)

type Operator int
type Function int
type Constant int

const (
	ADDITION Operator = iota
	SUBTRACTION
	MULTIPLICATION
	DIVISION
	MODULUS
	NEGATION
	EXPONENTIATION
	LEFT_PAREN
	RIGHT_PAREN
	COMMA
)

const (
	ABS Function = iota
	ACOS
	ASIN
	ATAN
	COS
	COSH
	EXP
	LN
	LOG
	ROUND
	SIN
	SINH
	SQRT
	TAN
	TANH
	TRUNC
	NEG
)

const (
	E Constant = iota
	PI
)

type Token interface {
}

type ValueToken struct {
	value float64
}

type OperatorToken struct {
	operator Operator
}

type FunctionToken struct {
	function Function
}

type ConstantToken struct {
	constant Constant
}

type tokenizerState struct {
	runes []rune
	idx   int
}

func Tokenize(expression string) ([]Token, error) {
	state := tokenizerState{runes: []rune(expression), idx: 0}
	tokens := make([]Token, 0)
	for {
		token, err := nextToken(&state)
		if err != nil {
			return nil, err
		}
		if token == nil {
			break
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func nextToken(state *tokenizerState) (Token, error) {
	skipWhitespace(state)
	c := currChar(state)
	if c == 0 {
		return nil, nil
	}
	if c == '.' || c >= '0' && c <= '9' {
		number, err := scanNumber(state)
		if err != nil {
			return nil, err
		}
		return ValueToken{number}, nil
	} else if c == '_' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
		identifier, err := scanIdentifier(state)
		if err != nil {
			return nil, err
		}
		return functionOrConstantToken(state, identifier)
	}
	nextChar(state)
	if c == '+' {
		return OperatorToken{ADDITION}, nil
	} else if c == '-' {
		return OperatorToken{SUBTRACTION}, nil
	} else if c == '*' {
		return OperatorToken{MULTIPLICATION}, nil
	} else if c == '/' {
		return OperatorToken{DIVISION}, nil
	} else if c == '%' {
		return OperatorToken{MODULUS}, nil
	} else if c == '^' {
		return OperatorToken{EXPONENTIATION}, nil
	} else if c == '(' {
		return OperatorToken{LEFT_PAREN}, nil
	} else if c == ')' {
		return OperatorToken{RIGHT_PAREN}, nil
	} else if c == ',' {
		return OperatorToken{COMMA}, nil
	}
	return nil, fmt.Errorf("unexpected character %c", c)
}

func skipWhitespace(state *tokenizerState) {
	for {
		c := currChar(state)
		if !unicode.IsSpace(c) {
			break
		}
		nextChar(state)
	}
}

func scanNumber(state *tokenizerState) (float64, error) {
	number := 0.0
	divider := 0.1
	dotSeen := false
	for {
		c := currChar(state)
		if c == 0 {
			break
		}
		if c == '.' {
			dotSeen = true
		} else if c == 'e' || c == 'E' {
			c := nextChar(state)
			if c == 0 {
				return 0.0, errors.New("expected something after the exponent sign")
			}
			sign := 1.0
			if c == '-' {
				sign = -1.0
				c := nextChar(state)
				if c == 0 {
					return 0.0, errors.New("expected something after the exponent - sign")
				}
			} else if c == '+' {
				c := nextChar(state)
				if c == 0 {
					return 0.0, errors.New("expected something after the exponent + sign")
				}
			}
			exp, err := scanNumber(state)
			if err != nil {
				return 0.0, err
			}
			number = number * math.Pow(10.0, sign*exp)
			break
		} else if c >= '0' && c <= '9' {
			digit := float64(c - '0')
			if dotSeen {
				number += digit * divider
				divider /= 10.0
			} else {
				number = number*10.0 + digit
			}
		} else {
			break
		}
		nextChar(state)
	}
	return number, nil
}

func functionOrConstantToken(state *tokenizerState, identifier string) (Token, error) {
	switch strings.ToUpper(identifier) {
	case "ABS":
		return FunctionToken{ABS}, nil
	case "ACOS":
		return FunctionToken{ACOS}, nil
	case "ASIN":
		return FunctionToken{ASIN}, nil
	case "ATAN":
		return FunctionToken{ATAN}, nil
	case "COS":
		return FunctionToken{COS}, nil
	case "COSH":
		return FunctionToken{COSH}, nil
	case "EXP":
		return FunctionToken{EXP}, nil
	case "LN":
		return FunctionToken{LN}, nil
	case "LOG":
		return FunctionToken{LOG}, nil
	case "ROUND":
		return FunctionToken{ROUND}, nil
	case "SIN":
		return FunctionToken{SIN}, nil
	case "SINH":
		return FunctionToken{SINH}, nil
	case "SQRT":
		return FunctionToken{SQRT}, nil
	case "TAN":
		return FunctionToken{TAN}, nil
	case "TANH":
		return FunctionToken{TANH}, nil
	case "TRUNC":
		return FunctionToken{TRUNC}, nil
	case "NEG":
		return FunctionToken{NEG}, nil
	case "E":
		return ConstantToken{E}, nil
	case "PI":
		return ConstantToken{PI}, nil
	default:
		return nil, fmt.Errorf("unknown function or constant %s", identifier)
	}
}

func scanIdentifier(state *tokenizerState) (string, error) {
	identifier := ""
	c := currChar(state)
	for {
		if !(c != 0 && (c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z')) {
			break
		}
		identifier += string(c)
		c = nextChar(state)
	}
	return identifier, nil
}

func currChar(state *tokenizerState) rune {
	if state.idx >= len(state.runes) {
		return 0
	}
	return state.runes[state.idx]
}

func nextChar(state *tokenizerState) rune {
	state.idx++
	if state.idx >= len(state.runes) {
		return 0
	}
	return state.runes[state.idx]
}

func peekNextChar(state *tokenizerState) rune {
	if state.idx >= len(state.runes)-1 {
		return 0
	}
	return state.runes[state.idx+1]
}
