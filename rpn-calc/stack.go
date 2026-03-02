package main

import (
	"errors"
	"strconv"
	"strings"
)

type Stack []string

func EnstackLine(stack Stack, line string) Stack {
	for _, token := range strings.Fields(line) {
		stack = append(stack, token)
	}
	return stack
}

func (stack *Stack) Peek() (string, error) {
	stackSize := len(*stack)
	if stackSize == 0 {
		return "", errors.New("Peek empty stack")
	}

	return (*stack)[stackSize-1], nil
}

func (stack *Stack) Pop() (string, error) {
	stackSize := len(*stack)
	if stackSize == 0 {
		return "", errors.New("Pop empty stack")
	}

	last := stackSize - 1
	tos := (*stack)[last]
	*stack = (*stack)[:last]

	return tos, nil
}

func (stack *Stack) PopFloat64() (float64, error) {
	tos, err := stack.Pop()
	if err != nil {
		return 0, err
	}
	num, err := strconv.ParseFloat(tos, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (stack *Stack) Push(token string) {
	*stack = append(*stack, token)
}
