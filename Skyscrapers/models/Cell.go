package models

type Cell struct {
	Row    int
	Column int
	*CanRep
}

type Elimination struct {
	Row     int
	Col     int
	Num     int
	Comment string
}

func (e *Elimination) Equals(other *Elimination) bool {
	if e.Row != other.Row {
		return false
	}
	if e.Col != other.Col {
		return false
	}
	if e.Num != other.Num {
		return false
	}
	return true
}

func NewCell(row, col int) *Cell {
	return &Cell{

		Row:    row,
		Column: col,
		CanRep: NewCanRep(),
	}
}

func (c *Cell) DoesContain(number int) bool {

	return c.CanRep.DoesContain(number)
}

/*
data contains correct candidates
eliminate candidated from rep, if other does not contains them
returns list of eliminations
*/
func (c *Cell) ProcessAbsent(data *CanRep) []Elimination {
	result := make([]Elimination, 0)
	for num := range c.list {
		if !data.DoesContain(num) {
			el := Elimination{
				Row: c.Row,
				Col: c.Column,
				Num: num,
			}
			result = append(result, el)
		}
	}
	return result
}
