package main

import (
	"math"
	"testing"
)

func TestStackCalculator(t *testing.T) {
	state := stackCalculatorState{make([]float64, 0)}
	state.Push(1)
	state.Push(2)
	assertPopEquals(&state, 2, t)
	assertPopEquals(&state, 1, t)
	if len(state.stack) != 0 {
		t.Errorf("stack length should be 0")
	}

	state.Push(1)
	err := state.negate()
	if err != nil {
		t.Error(err)
	}
	assertPopEquals(&state, -1, t)

	state.Push(2)
	state.Push(3)
	err = state.add()
	if err != nil {
		t.Error(err)
	}
	assertPopEquals(&state, 2+3, t)

	state.Push(2)
	state.Push(3)
	err = state.subtract()
	if err != nil {
		t.Error(err)
	}
	assertPopEquals(&state, 2-3, t)

	state.Push(2)
	state.Push(3)
	err = state.multiply()
	if err != nil {
		t.Error(err)
	}
	assertPopEquals(&state, 2*3, t)

	state.Push(16)
	state.Push(4)
	err = state.divide()
	if err != nil {
		t.Error(err)
	}
	assertPopEquals(&state, 16/4, t)

	state.Push(10)
	state.Push(3)
	err = state.modulus()
	if err != nil {
		t.Error(err)
	}
	assertPopEquals(&state, 10%3, t)

	state.Push(10)
	state.Push(3)
	err = state.exponentiate()
	if err != nil {
		t.Error(err)
	}
	assertPopEquals(&state, math.Pow(10, 3), t)
}

func assertPopEquals(state *stackCalculatorState, expected float64, t *testing.T) {
	value, err := state.Pop()
	if err != nil {
		t.Error(err)
	}
	if value != expected {
		t.Errorf("Expected %v, got %v", expected, value)
	}
}
