package main

import "testing"

func TestTokenize(t *testing.T) {
	tokens, err := Tokenize("2.0e-1*SIN(PI)")
	if err != nil {
		t.Fatal(err)
	}
	numTokens := len(tokens)
	expectedNumTokens := 6
	if numTokens != expectedNumTokens {
		t.Errorf("got %d, want %d", numTokens, expectedNumTokens)
	}
	_, ok := tokens[0].(ValueToken)
	if !ok {
		t.Errorf("got %T, want ValueToken", tokens[0])
	}
	if tokens[0].(ValueToken).value != 0.2 {
		t.Errorf("got %f, want %f", tokens[0].(ValueToken).value, 0.2)
	}
	_, ok = tokens[1].(OperatorToken)
	if !ok {
		t.Errorf("got %T, want OperatorToken", tokens[1])
	}
	if tokens[1].(OperatorToken).operator != MULTIPLICATION {
		t.Errorf("got %d, want MULTIPLICATION", tokens[1].(OperatorToken).operator)
	}
	_, ok = tokens[2].(FunctionToken)
	if !ok {
		t.Errorf("got %T, want FunctionToken", tokens[2])
	}
	if tokens[2].(FunctionToken).function != SIN {
		t.Errorf("got %d, want SIN", tokens[2].(FunctionToken).function)
	}
	_, ok = tokens[3].(OperatorToken)
	if !ok {
		t.Errorf("got %T, want OperatorToken", tokens[3])
	}
	if tokens[3].(OperatorToken).operator != LEFT_PAREN {
		t.Errorf("got %d, want LEFT_PAREN", tokens[3].(OperatorToken).operator)
	}
	_, ok = tokens[4].(ConstantToken)
	if !ok {
		t.Errorf("got %T, want ConstantToken", tokens[4])
	}
	if tokens[4].(ConstantToken).constant != PI {
		t.Errorf("got %d, want PI", tokens[4].(ConstantToken).constant)
	}
	_, ok = tokens[5].(OperatorToken)
	if !ok {
		t.Errorf("got %T, want OperatorToken", tokens[5])
	}
	if tokens[5].(OperatorToken).operator != RIGHT_PAREN {
		t.Errorf("got %d, want RIGHT_PAREN", tokens[5].(OperatorToken).operator)
	}
}
