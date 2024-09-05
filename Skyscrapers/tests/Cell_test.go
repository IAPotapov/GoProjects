package tests

import (
	"Skyscrapers/models"
	"testing"
)

func TestEquals(t *testing.T) {
	c := models.NewCell(0, 0)
	c.Add(1)
	c.Add(4)
	c.Add(3)
	list := []int{4, 3, 1}
	// b := c.Equals(list)
	// if !b {
	// 	t.Fatalf("Expected Equals to be true. Got false. %v and %v", c.CanRep.list, list)
	// }
	for _, v := range list {
		if !c.DoesContain(v) {
			t.Fatalf("Expected to contain %v, got false", v)
		}
	}
}
