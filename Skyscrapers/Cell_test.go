package main

import (
	"testing"
)

func TestEquals(t *testing.T) {
	c := &Cell{
		CanRep: NewCanRep(),
	}
	c.Add(1)
	c.Add(4)
	c.Add(3)
	list := []int{4, 3, 1}
	b := c.Equals(list)
	if !b {
		t.Fatalf("Expected Equals to be true. Got false. %v and %v", c.CanRep.list, list)
	}
}
