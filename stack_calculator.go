package main

import (
	"errors"
	"math"
)

type stackCalculatorState struct {
	stack []float64
}

func (s *stackCalculatorState) Push(n float64) {
	s.stack = append(s.stack, n)
}

func (s *stackCalculatorState) Pop() (float64, error) {
	if len(s.stack) == 0 {
		return 0, errors.New("stack underflow")
	}
	lastIndex := len(s.stack) - 1
	n := s.stack[lastIndex]
	s.stack = s.stack[:lastIndex]
	return n, nil
}

func (s *stackCalculatorState) PopLast() (float64, error) {
	n, err := s.Pop()
	if err != nil {
		return 0, err
	}
	if len(s.stack) > 0 {
		return 0, errors.New("stack is not empty when it should be")
	}
	return n, nil
}

func (s *stackCalculatorState) negate() error {
	operand1, err := s.Pop()
	if err != nil {
		return err
	}
	s.Push(-operand1)
	return nil
}

func (s *stackCalculatorState) add() error {
	operand2, err := s.Pop()
	if err != nil {
		return err
	}
	operand1, err := s.Pop()
	if err != nil {
		return err
	}
	s.Push(operand1 + operand2)
	return nil
}

func (s *stackCalculatorState) subtract() error {
	operand2, err := s.Pop()
	if err != nil {
		return err
	}
	operand1, err := s.Pop()
	if err != nil {
		return err
	}
	s.Push(operand1 - operand2)
	return nil
}

func (s *stackCalculatorState) multiply() error {
	operand2, err := s.Pop()
	if err != nil {
		return err
	}
	operand1, err := s.Pop()
	if err != nil {
		return err
	}
	s.Push(operand1 * operand2)
	return nil
}

func (s *stackCalculatorState) divide() error {
	operand2, err := s.Pop()
	if err != nil {
		return err
	}
	operand1, err := s.Pop()
	if err != nil {
		return err
	}
	s.Push(operand1 / operand2)
	return nil
}

func (s *stackCalculatorState) modulus() error {
	operand2, err := s.Pop()
	if err != nil {
		return err
	}
	operand1, err := s.Pop()
	if err != nil {
		return err
	}
	s.Push(math.Mod(operand1, operand2))
	return nil
}

func (s *stackCalculatorState) exponentiate() error {
	operand2, err := s.Pop()
	if err != nil {
		return err
	}
	operand1, err := s.Pop()
	if err != nil {
		return err
	}
	s.Push(math.Pow(operand1, operand2))
	return nil
}

func Calculate(tokens []Token) (float64, error) {
	state := stackCalculatorState{make([]float64, 0)}
	for _, token := range tokens {
		if valueToken, ok := token.(ValueToken); ok {
			state.Push(valueToken.value)
		} else if operatorToken, ok := token.(OperatorToken); ok {
			var err error
			switch operatorToken.operator {
			case ADDITION:
				err = state.add()
			case SUBTRACTION:
				err = state.subtract()
			case MULTIPLICATION:
				err = state.multiply()
			case DIVISION:
				err = state.divide()
			case MODULUS:
				err = state.modulus()
			case NEGATION:
				err = state.negate()
			case EXPONENTIATION:
				err = state.exponentiate()
			default:
				err = errors.New("unhandled operator")
			}
			if err != nil {
				return 0.0, err
			}
		} else if functionToken, ok := token.(FunctionToken); ok {
			var err error
			var value float64
			switch functionToken.function {
			case ABS:
				value, err = state.Pop()
				if err == nil {
					state.Push(math.Abs(value))
				}
			case ACOS:
				value, err = state.Pop()
				if err == nil {
					state.Push(math.Acos(value))
				}
			case ASIN:
				value, err = state.Pop()
				if err == nil {
					state.Push(math.Asin(value))
				}
			case ATAN:
				value, err = state.Pop()
				if err == nil {
					state.Push(math.Atan(value))
				}
			case COS:
				value, err = state.Pop()
				if err == nil {
					state.Push(math.Cos(value))
				}
			case COSH:
				value, err = state.Pop()
				if err == nil {
					state.Push(math.Cosh(value))
				}
			case EXP:
				value, err = state.Pop()
				if err == nil {
					state.Push(math.Exp(value))
				}
			case LN:
				value, err = state.Pop()
				if err == nil {
					state.Push(math.Log(value))
				}
			case LOG:
				value, err = state.Pop()
				if err == nil {
					state.Push(math.Log10(value))
				}
			case ROUND:
				value, err = state.Pop()
				if err == nil {
					state.Push(math.Round(value))
				}
			case SIN:
				value, err = state.Pop()
				if err == nil {
					state.Push(math.Sin(value))
				}
			case SINH:
				value, err = state.Pop()
				if err == nil {
					state.Push(math.Sinh(value))
				}
			case SQRT:
				value, err = state.Pop()
				if err == nil {
					state.Push(math.Sqrt(value))
				}
			case TAN:
				value, err = state.Pop()
				if err == nil {
					state.Push(math.Tan(value))
				}
			case TANH:
				value, err = state.Pop()
				if err == nil {
					state.Push(math.Tanh(value))
				}
			case TRUNC:
				value, err = state.Pop()
				if err == nil {
					state.Push(math.Trunc(value))
				}
			case NEG:
				value, err = state.Pop()
				if err == nil {
					state.Push(-value)
				}
			default:
				err = errors.New("unhandled function")
			}
			if err != nil {
				return 0.0, err
			}
		} else if constantToken, ok := token.(ConstantToken); ok {
			switch constantToken.constant {
			case E:
				state.Push(math.E)
			case PI:
				state.Push(math.Pi)
			default:
				return 0.0, errors.New("unhandled constant")
			}
		}
	}
	return state.PopLast()
}
