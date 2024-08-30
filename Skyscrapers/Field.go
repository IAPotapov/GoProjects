package main

type Field struct {
	Size  int
	Clues []int
	Cells []*Cell
	Lines []Line
}

func NewField(clues []int) *Field {

	var field Field = Field{
		Size:  len(clues) / 4,
		Clues: clues,
	}
	field.Cells = make([]*Cell, field.Size*field.Size)
	for i := 0; i < len(field.Cells); i++ {
		field.Cells[i] = NewCell(i/field.Size, i%field.Size)
		field.Cells[i].Solution = 0
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

	return &field
}

func (field *Field) ProcessSoloCandidates() {
	for i := 0; i < len(field.Lines); i++ {
		r := field.Lines[i].FindSoloCandidates()
		for j := 0; j < len(r); j++ {
			field.ProcessSolution(r[j].Row, r[j].Col, r[j].Number)
		}
	}

}

func (field *Field) ProcessCaseAlpha() (done bool) {
	done = false
	for i := 0; i < len(field.Lines); i++ {
		field.Lines[i].GetPermutations()
		field.Lines[i].RemoveIncorrectPermutations()
		b := field.Lines[i].ProcessCaseAlpha()
		if b {
			done = true
		}
	}
	return
}

func (field *Field) ProcessSolution(row, col, num int) {
	i := row*field.Size + col
	field.Cells[i].Solution = num

	line := field.Lines[row]
	line.RemoveCandidate(num)
	line = field.Lines[field.Size+col]
	line.RemoveCandidate(num)
}

func (field *Field) FindSolution() bool {
	for {
		ok := field.ProcessCaseAlpha()
		if !ok {
			break
		}
		field.ProcessSoloCandidates()
	}
	result := true

	for i := 0; i < len(field.Cells); i++ {
		if field.Cells[i].Solution == 0 {
			result = false
		}
	}
	return result
}
