package main

type Cell struct {
	Solution int
	Row      int
	Column   int
	*CanRep
}

func NewCell(row, col int) *Cell {
	return &Cell{
		Solution: 0,
		Row:      row,
		Column:   col,
		CanRep:   NewCanRep(),
	}
}

/*
	func (c *Cell) GetEliminations(correctCandidates []int) []Candidate {
		result := make([]Candidate, 0)
		for i := 0; i < len(c.Candidates); i++ {
			found := false
			for j := 0; j < len(correctCandidates); j++ {
				if c.Candidates[i] == correctCandidates[j] {
					found = true
					break
				}
			}
			if !found {
				e := Candidate{Row: c.Row, Col: c.Column, Number: c.Candidates[i]}
				result = append(result, e)
			}
		}
		return result
	}
*/

/*func (c *Cell) Equals(other *CanRep) bool {
	if len(c.list) != len(other.list) {
		return false
	}
	for item := range other.list {
		if !c.DoesContain(item) {
			return false
		}
	}
	return true
}

func (c *Cell) SetCandidates(other *CanRep) {
	c.CanRep = other
}*/

func (c *Cell) DoesContain(number int) bool {
	if c.Solution > 0 {
		return c.Solution == number
	}
	return c.CanRep.DoesContain(number)
}

func (c *Cell) RemoveCandidate(number int) {
	if c.Solution == 0 {
		c.Remove(number)
	}
}
