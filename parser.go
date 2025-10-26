package main

import "errors"

type parserState struct {
	inTokens  []Token
	index     int
	token     Token
	outTokens []Token
}

func (s *parserState) eof() bool {
	return s.token == nil
}

func (s *parserState) next() {
	if s.index >= len(s.inTokens) {
		s.token = nil
	} else {
		s.token = s.inTokens[s.index]
		s.index++
	}
}

func (s *parserState) nextCheckEof() error {
	s.next()
	if s.eof() {
		return errors.New("unexpected end of input")
	}
	return nil
}

func (s *parserState) isOperatorMatch(ot Operator) bool {
	if s.eof() {
		return false
	}
	if operatorToken, ok := s.token.(OperatorToken); ok {
		return operatorToken.operator == ot
	}
	return false
}

func (s *parserState) addOutToken(token Token) {
	s.outTokens = append(s.outTokens, token)
}

func ConvertInfixToPostfix(inTokens []Token) ([]Token, error) {
	state := parserState{inTokens: inTokens, index: 0, token: nil, outTokens: make([]Token, 0)}
	err := state.nextCheckEof()
	if err != nil {
		return nil, err
	}
	err = parseExpression(&state)
	if err != nil {
		return nil, err
	}
	if !state.eof() {
		return nil, errors.New("unexpected text at the end")
	}
	return state.outTokens, nil
}

func parseFunctionExpression(state *parserState) error {
	functionToken, _ := state.token.(FunctionToken)
	state.next()
	if !state.isOperatorMatch(LEFT_PAREN) {
		return errors.New("missing ( after function name")
	}
	err := state.nextCheckEof()
	if err != nil {
		return err
	}
	for !state.isOperatorMatch(RIGHT_PAREN) {
		err := parseExpression(state)
		if err != nil {
			return err
		}
		if state.isOperatorMatch(COMMA) {
			err := state.nextCheckEof()
			if err != nil {
				return err
			}
			if state.isOperatorMatch(RIGHT_PAREN) {
				return errors.New("missing function argument after comma")
			}
		}
	}
	state.next()
	state.addOutToken(functionToken)
	return nil
}

func parsePrimaryExpression(state *parserState) error {
	if _, ok := state.token.(ValueToken); ok {
		state.addOutToken(state.token)
		state.next()
		return nil
	}
	if _, ok := state.token.(ConstantToken); ok {
		state.addOutToken(state.token)
		state.next()
		return nil
	}
	if _, ok := state.token.(FunctionToken); ok {
		err := parseFunctionExpression(state)
		if err != nil {
			return err
		}
		return nil
	}
	if state.isOperatorMatch(LEFT_PAREN) {
		err := state.nextCheckEof()
		if err != nil {
			return err
		}
		err = parseExpression(state)
		if err != nil {
			return err
		}
		if !state.isOperatorMatch(RIGHT_PAREN) {
			return errors.New("unmatched parenthesis")
		}
		state.next()
		return nil
	}
	return errors.New("unexpected operator")
}

func parseUnaryExpression(state *parserState) error {
	negate := false
	if state.isOperatorMatch(SUBTRACTION) {
		negate = true
		err := state.nextCheckEof()
		if err != nil {
			return err
		}
	} else if state.isOperatorMatch(ADDITION) {
		err := state.nextCheckEof()
		if err != nil {
			return err
		}
	}
	err := parsePrimaryExpression(state)
	if err != nil {
		return err
	}
	if negate {
		state.addOutToken(OperatorToken{NEGATION})
	}
	return nil
}

func parseExponentialExpression(state *parserState) error {
	err := parseUnaryExpression(state)
	if err != nil {
		return err
	}
	count := 0
	for state.isOperatorMatch(EXPONENTIATION) {
		err := state.nextCheckEof()
		if err != nil {
			return err
		}
		err = parseUnaryExpression(state)
		if err != nil {
			return err
		}
		count++
	}
	for q := 0; q < count; q++ {
		state.addOutToken(OperatorToken{EXPONENTIATION})
	}
	return nil
}

func parseMultiplicativeExpression(state *parserState) error {
	err := parseExponentialExpression(state)
	if err != nil {
		return err
	}
	for state.isOperatorMatch(MULTIPLICATION) || state.isOperatorMatch(DIVISION) || state.isOperatorMatch(MODULUS) {
		var operatorToken Token
		if state.isOperatorMatch(MULTIPLICATION) {
			operatorToken = OperatorToken{MULTIPLICATION}
		} else if state.isOperatorMatch(DIVISION) {
			operatorToken = OperatorToken{DIVISION}
		} else {
			operatorToken = OperatorToken{MODULUS}
		}
		err = state.nextCheckEof()
		if err != nil {
			return err
		}
		err = parseExponentialExpression(state)
		if err != nil {
			return err
		}
		state.addOutToken(operatorToken)
	}
	return nil
}

func parseAdditiveExpression(state *parserState) error {
	err := parseMultiplicativeExpression(state)
	if err != nil {
		return err
	}
	for state.isOperatorMatch(ADDITION) || state.isOperatorMatch(SUBTRACTION) {
		var operatorToken Token
		if state.isOperatorMatch(ADDITION) {
			operatorToken = OperatorToken{ADDITION}
		} else {
			operatorToken = OperatorToken{SUBTRACTION}
		}
		err = state.nextCheckEof()
		if err != nil {
			return err
		}
		err = parseMultiplicativeExpression(state)
		if err != nil {
			return err
		}
		state.addOutToken(operatorToken)
	}
	return nil
}

func parseExpression(state *parserState) error {
	return parseAdditiveExpression(state)
}
