package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListIsLooped(t *testing.T) {
	l := NewList()

	for i := 0; i < 15; i++ {
		l.Add(i)
	}

	fmt.Println(l)
	l.firstNode.next.next = l.firstNode

	assert.True(t, IsListLooped(l), "List should be looped")
}

func TestListIsNotLooped(t *testing.T) {
	l := NewList()

	for i := 0; i < 30; i++ {
		l.Add(i)
	}

	assert.False(t, IsListLooped(l), "List shouldn't be looped")
}

func TestEmptyListIsNotLooped(t *testing.T) {
	l := NewList()

	assert.False(t, IsListLooped(l), "Empty list shouldn't be looped")
}

func TestListWithOneElementIsNotLooped(t *testing.T) {
	l := NewList()
	l.Add(42)

	assert.False(t, IsListLooped(l), "List with one element shouldn't be looped")
}
