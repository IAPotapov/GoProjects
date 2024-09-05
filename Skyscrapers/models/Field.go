package models

type FieldStatus int

const (
	Unknown = iota
	Fail
	Solved
)

type Field struct {
	Size   int
	Clues  []int
	Cells  []*Cell
	Lines  []Line
	Status FieldStatus
}

func NewField(clues []int) *Field {

	var field Field = Field{
		Size:  len(clues) / 4,
		Clues: clues,
	}
	field.Cells = make([]*Cell, field.Size*field.Size)
	for i := 0; i < len(field.Cells); i++ {
		field.Cells[i] = NewCell(i/field.Size, i%field.Size)

		for j := 0; j < field.Size; j++ {
			field.Cells[i].Add(j + 1)
		}
	}
	field.Lines = make([]Line, field.Size*2)

	for y := 0; y < field.Size; y++ {
		y2 := y + field.Size
		field.Lines[y].Cells = make([]*Cell, field.Size)
		field.Lines[y].LeftClue = field.Clues[len(field.Clues)-1-y]
		field.Lines[y].RightClue = field.Clues[field.Size+y]
		field.Lines[y2].Cells = make([]*Cell, field.Size)
		field.Lines[y2].LeftClue = field.Clues[y]
		field.Lines[y2].RightClue = field.Clues[len(field.Clues)-1-field.Size-y]
		for x := 0; x < field.Size; x++ {
			i := x + y*field.Size
			i2 := y + x*field.Size
			field.Lines[y].Cells[x] = field.Cells[i]
			field.Lines[y2].Cells[x] = field.Cells[i2]
		}
	}
	field.Status = Unknown
	return &field
}

/*
Search eliminations in all lines
returns first found
otherwise returns empty list
*/
func (field *Field) FindAnyElimination() []Elimination {
	var result []Elimination
	for i := 0; i < len(field.Lines); i++ {
		result = field.Lines[i].ProcessCaseAlpha()
		if len(result) > 0 {
			return result
		}
		result = field.Lines[i].ProcessNakedCandidates()
		if len(result) > 0 {
			return result
		}
		result = field.Lines[i].ProcessHiddenCandidates()
		if len(result) > 0 {
			return result
		}
	}
	return make([]Elimination, 0)
}

func (field *Field) CommitEllimination(e *Elimination) {
	i := e.Row*field.Size + e.Col
	field.Cells[i].Remove(e.Num)
}
func FindCommon(list0, list1 []Elimination) []Elimination {
	result := make([]Elimination, 0)
	for i := 0; i < len(list0); i++ {
		for j := 0; j < len(list1); j++ {
			if list0[i].Equals(&list1[j]) {
				result = append(result, list0[i])
				break
			}
		}
	}
	return result
}

func (field *Field) FindEliminationsByPair(index int) []Elimination {
	save := field.Serialize()
	cands := field.Cells[index].GetAll()

	e0 := Elimination{
		Row:     field.Cells[index].Row,
		Col:     field.Cells[index].Column,
		Num:     cands[0],
		Comment: "Guess",
	}
	e1 := Elimination{
		Row:     field.Cells[index].Row,
		Col:     field.Cells[index].Column,
		Num:     cands[1],
		Comment: "Guess",
	}
	// option 0
	field.CommitEllimination(&e0)
	list0 := field.FindAllElliminations()
	if field.Status == Solved {
		list0 = append(list0, e0)
		return list0
	}
	field.Deserialize(save)

	// option 1
	field.CommitEllimination(&e1)
	list1 := field.FindAllElliminations()
	if field.Status == Solved {
		list1 = append(list1, e1)
		return list1
	}
	field.Deserialize(save)
	// compare list0 and list1
	return FindCommon(list0, list1)
}

func (field *Field) FindAllElliminations() []Elimination {
	result := make([]Elimination, 0)
	for {
		r := field.FindAnyElimination()
		if len(r) > 0 {
			result = append(result, r...)
			for i := 0; i < len(r); i++ {
				field.CommitEllimination(&r[i])
			}
		} else {
			break
		}
	}

	field.Status = Solved
	for i := 0; i < len(field.Cells); i++ {
		n := len(field.Cells[i].list)
		if n == 0 {
			field.Status = Fail
			break
		}
		if n > 1 {
			field.Status = Unknown
			break
		}
	}

	return result
}

func (field *Field) Guess() FieldStatus {
	field.FindAllElliminations()
	if field.Status == Solved {
		return Solved
	}
	if field.Status == Fail {
		return Fail
	}
	pilist := field.FindAllPairs()
	for _, i := range pilist {
		e := field.FindEliminationsByPair(pilist[i])
		if len(e) > 0 {
			for j := 0; j < len(e); j++ {
				field.CommitEllimination(&e[j])
			}
			return field.Guess()
		}
	}
	return Unknown
}

func (field *Field) Serialize() [][]int {
	result := make([][]int, len(field.Cells))
	for i := 0; i < len(field.Cells); i++ {
		result[i] = field.Cells[i].GetAll()
	}
	return result
}

func (field *Field) Deserialize(data [][]int) {
	for i := 0; i < len(field.Cells); i++ {
		r := data[i]
		field.Cells[i].SetAll(r)
	}
	field.Status = Unknown
}

// returns indexes of all cells, containing only two candidates
func (field *Field) FindAllPairs() []int {
	result := make([]int, 0)
	for i := 0; i < len(field.Cells); i++ {
		if len(field.Cells[i].list) == 2 {
			result = append(result, i)
		}
	}
	return result
}
