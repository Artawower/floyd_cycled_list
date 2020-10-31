package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsCycling(t *testing.T) {
	l := &List{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.firstNode.next.next = l.firstNode.next

	assert.True(t, isCycledList(l), "List should be cycled")
}

func TestIsNotCycling(t *testing.T) {
	l := &List{}
	for i := 0; i < 20; i++ {
		l.Add(i)
	}
	// assert.False(t, isCycledList(l), "List should not be cycled")
	assert.False(t, isCycledList(l), "List should not be cycled")
}
