package main

import (
	"errors"
	"strconv"
	"strings"
)

type Stack []any

func EnstackLine(stack Stack, line string) Stack {
	for _, token := range strings.Fields(line) {
		num, err := strconv.ParseFloat(token, 64)
		if err != nil {
			stack = append(stack, token)
		} else {
			stack = append(stack, num)
		}
	}
	return stack
}

func (stack *Stack) Peek() (any, error) {
	stackSize := len(*stack)
	if stackSize == 0 {
		return "", errors.New("Peek empty stack")
	}

	return (*stack)[stackSize-1], nil
}

func (stack *Stack) Pop() (any, error) {
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

	if value, ok := tos.(float64); ok {
		return value, nil
	}
	return 0, errors.New("Not a float")
}

func (stack *Stack) Push(token any) {
	*stack = append(*stack, token)
}
