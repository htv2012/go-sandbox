package main

import (
	"errors"
)

type Tree struct {
	value       int
	left, right *Tree
}

type Visit func(int)

func (t *Tree) Insert(value int) (*Tree, error) {
	switch {
	case t == nil:
		return &Tree{value: value}, nil
	case t.value == value:
		return nil, errors.New("Duplicate value")
	case value < t.value:
		newLeft, err := t.left.Insert(value)
		if err != nil {
			return nil, err
		}
		t.left = newLeft
		return t, nil
	default:
		newRight, err := t.right.Insert(value)
		if err != nil {
			return nil, err
		}
		t.right = newRight
		return t, nil
	}
}

func (t *Tree) Contain(value int) bool {
	switch {
	case t == nil:
		return false
	case value == t.value:
		return true
	case value < t.value:
		return t.left.Contain(value)
	default:
		return t.right.Contain(value)
	}
}

func (t *Tree) InOrderTraverse(visit Visit) {
	if t == nil {
		return
	}
	t.left.InOrderTraverse(visit)
	visit(t.value)
	t.right.InOrderTraverse(visit)
}

func (t *Tree) PreOrderTraverse(visit Visit) {
	if t == nil {
		return
	}
	visit(t.value)
	t.left.PreOrderTraverse(visit)
	t.right.PreOrderTraverse(visit)
}

func (t *Tree) PostOrderTraverse(visit Visit) {
	if t == nil {
		return
	}
	t.left.PostOrderTraverse(visit)
	t.right.PostOrderTraverse(visit)
	visit(t.value)
}
