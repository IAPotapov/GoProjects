package tests

// import (
// 	"Skyscrapers/models"
// 	"fmt"
// 	"strings"
// 	"testing"
// )

/*func DoTest(line *models.Line, expectList []int, t *testing.T) {
	line.GetFreeCandidates()
	result := line.freeCandidates
	sort.Ints(result)
	if len(result) != len(expectList) {
		t.Fatalf("Expected len: %v, Got len %v", len(expectList), len(result))
	}
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	encoder.Encode(result)
	result2 := writer.String()
	writer.Reset()
	encoder.Encode(expectList)
	result3 := writer.String()
	if result2 != result3 {
		t.Fatalf("Expected: %v, Got: %v", result3, result2)
	}
}

func TestGetFreeCandidates(t *testing.T) {
	x := &models.Line{
		Index:     0,
		IsRow:     true,
		LeftClue:  4,
		RightClue: 1,
		Cells: []*models.Cell{
			models.NewCell(0, 1),
			models.NewCell(0, 2),
			models.NewCell(0, 3),
			models.NewCell(0, 4),
		},
	}
	for i := 0; i < len(x.Cells); i++ {
		for j := 1; j <= 4; j++ {
			x.Cells[i].Add(j)
		}
	}
	DoTest(x, []int{1, 2, 3, 4}, t)

	x = &Line{
		Index:     0,
		IsRow:     true,
		LeftClue:  4,
		RightClue: 1,
		Cells: []Cell{
			{Solution: 0, Candidates: []int{1, 2, 3}},
			{Solution: 0, Candidates: []int{1, 2, 3}},
			{Solution: 0, Candidates: []int{1, 2, 3}},
			{Solution: 4, Candidates: []int{4}},
		},
	}
	DoTest(x, []int{1, 2, 3}, t)

	x = &Line{
		Index:     0,
		IsRow:     true,
		LeftClue:  4,
		RightClue: 1,
		Cells: []Cell{
			{Solution: 0, Candidates: []int{1, 2}},
			{Solution: 0, Candidates: []int{1, 2}},
			{Solution: 0, Candidates: []int{1, 2, 3}},
			{Solution: 4, Candidates: []int{4}},
		},
	}
	DoTest(x, []int{1, 2, 3}, t)
}*/

/*func TestGetPermutations(t *testing.T) {
	x := &Line{
		Index:     0,
		IsRow:     true,
		LeftClue:  4,
		RightClue: 1,
		Cells: []Cell{
			{Solution: 0, Candidates: []int{1, 2}},
			{Solution: 0, Candidates: []int{1, 2}},
			{Solution: 3, Candidates: []int{3}},
			{Solution: 4, Candidates: []int{4}},
		},
	}
	x.GetPermutations()

	if len(x.permutations) != 2 {
		t.Fatalf("Expected len 2, Got: %v", len(x.permutations))
	}
	if x.permutations[0][0] == x.permutations[1][0] {
		t.Fatalf("Wrong permutations: %v and %v", x.permutations[0], x.permutations[1])
	}
}*/

/*func GetPermutationString(line *models.Line, index int) string {
	var sb strings.Builder

	for i := 0; i < len(line.permutations[index]); i++ {
		s := fmt.Sprintf("%d", line.permutations[index][i])
		sb.WriteString(s)
	}

	return sb.String()
}

func TestRemoveIncorrectPermutations(t *testing.T) {
	x := &models.Line{
		Index:     0,
		IsRow:     true,
		LeftClue:  4,
		RightClue: 1,
		Cells: []*models.Cell{
			models.NewCell(0, 1),
			models.NewCell(0, 2),
			models.NewCell(0, 3),
			models.NewCell(0, 4),
		},
	}
	for i := 0; i < len(x.Cells); i++ {
		for j := 1; j <= 4; j++ {
			x.Cells[i].Add(j)
		}
	}
	x.GetPermutations()
	x.RemoveIncorrectPermutations()

	if len(x.permutations) != 1 {
		t.Fatalf("Expected len 1, Got: %v, %v", len(x.permutations), x.permutations)
	}
	s := GetPermutationString(x, 0)
	if s != "1234" {
		t.Fatalf("Expected: \"1234\", Got: %v", s)
	}
}

func TestRemoveIncorrectPermutations2(t *testing.T) {
	x := &models.Line{
		Index:     0,
		IsRow:     true,
		LeftClue:  1,
		RightClue: 2,
		Cells: []*models.Cell{
			models.NewCell(0, 1),
			models.NewCell(0, 2),
			models.NewCell(0, 3),
			models.NewCell(0, 4),
		},
	}
	for i := 0; i < len(x.Cells); i++ {
		for j := 1; j <= 4; j++ {
			x.Cells[i].Add(j)
		}
	}
	x.GetPermutations()
	x.RemoveIncorrectPermutations()

	if len(x.permutations) != 2 {
		t.Fatalf("Expected len 2, Got: %v, %v", len(x.permutations), x.permutations)
	}

	s1 := GetPermutationString(x, 0)
	s2 := GetPermutationString(x, 1)

	if (s1 != "4123" && s2 != "4123") || (s1 != "4213" && s2 != "4213") {
		t.Fatalf("Expected 4132 and 4231, Got: %v, %v", s1, s2)
	}
}*/

/*func TestRemoveCandidates(t *testing.T) {
	x := &Line{
		Index:     0,
		IsRow:     true,
		LeftClue:  4,
		RightClue: 1,
		Cells: []Cell{
			{Solution: 1, Candidates: []int{1}},
			{Solution: 0, Candidates: []int{1, 2, 3, 4}},
			{Solution: 0, Candidates: []int{1, 2, 3, 4}},
			{Solution: 0, Candidates: []int{1, 2, 3, 4}},
		},
	}

	x.RemoveCandidate(1)

	a := len(x.Cells[1].Candidates)
	if a != 3 {
		t.Fatalf("Expected len 3, Got: %v, %v", a, x.Cells[1].Candidates)
	}

	if x.Cells[1].Candidates[0] != 2 || x.Cells[1].Candidates[1] != 3 || x.Cells[1].Candidates[2] != 4 {
		t.Fatalf("Expected [2,3,4], Got: %v", x.Cells[1].Candidates)
	}
}*/
