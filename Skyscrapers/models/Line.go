package models

import "fmt"

type Line struct {
	Index               int
	IsRow               bool
	LeftClue, RightClue int
	Cells               []*Cell
	freeCandidates      []int
	permutations        [][]int
}

func (line *Line) GetFreeCandidates() {
	result := make(map[int]bool)
	for _, c := range line.Cells {
		if c.ContainsSolution() {
			continue
		}
		for k := range c.CanRep.list {
			result[k] = true
		}
	}
	list := make([]int, 0)
	for x := range result {
		list = append(list, x)
	}
	line.freeCandidates = list
}

func (line *Line) AddPermutation() {
	p := make([]int, len(line.freeCandidates))
	copy(p, line.freeCandidates)
	line.permutations = append(line.permutations, p)
}

func (line *Line) Swap(index1, index2 int) {
	temp := line.freeCandidates[index1]
	line.freeCandidates[index1] = line.freeCandidates[index2]
	line.freeCandidates[index2] = temp
}

// Heap's algorithm
func (line *Line) Heaps(k int) {
	if k == 1 {
		line.AddPermutation()
	} else {
		line.Heaps(k - 1)
		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				line.Swap(i, k-1)
			} else {
				line.Swap(0, k-1)
			}
			line.Heaps(k - 1)
		}
	}
}

// Get all permutations for a line
func (line *Line) GetPermutations() {
	line.GetFreeCandidates()
	line.permutations = make([][]int, 0)
	len := len(line.freeCandidates)
	if len > 0 {
		line.Heaps(len)
	}
}

func (line *Line) PermutationIsCorrect(data []int) bool {
	var value, count int
	// check left clue
	if line.LeftClue > 0 {
		value = 0
		count = 0
		for i := 0; i < len(data); i++ {
			if data[i] > value {
				count++
				value = data[i]
			}
		}
		if count != line.LeftClue {
			return false
		}
	}
	// check right clue
	if line.RightClue > 0 {
		value = 0
		count = 0
		for i := len(data) - 1; i >= 0; i-- {
			if data[i] > value {
				count++
				value = data[i]
			}
		}
		if count != line.RightClue {
			return false
		}
	}
	return true
}

func (line *Line) RemoveIncorrectPermutations() {
	data := make([]int, len(line.Cells))
	for pindex := 0; pindex < len(line.permutations); {
		var i int = 0
		for j := 0; j < len(line.Cells); j++ {
			if line.Cells[j].ContainsSolution() {
				data[j] = line.Cells[j].GetFirst()
			} else {
				if i >= len(line.permutations[pindex]) {
					continue
				}
				data[j] = line.permutations[pindex][i]
				i++
			}
		}
		if line.PermutationIsCorrect(data) {
			pindex++
		} else {
			p1 := line.permutations[:pindex]
			p2 := line.permutations[pindex+1:]
			line.permutations = p1
			line.permutations = append(line.permutations, p2...)
		}
	}
}

/*
Case Alpha:
Get all permutations for a line
Remove incorrect permutatios (Permutation check)
Check correct permutations for candidate elimination
Return list of eliminations
*/
func (line *Line) ProcessCaseAlpha() []Elimination {
	result := make([]Elimination, 0)
	if line.LeftClue == 0 && line.RightClue == 0 {
		return result
	}
	line.GetPermutations()
	line.RemoveIncorrectPermutations()
	if len(line.permutations) == 0 {
		return result
	}
	var i int = 0
	for cindex := 0; cindex < len(line.Cells); cindex++ {
		if line.Cells[cindex].ContainsSolution() {
			continue
		}
		// data is the list of candidates for cell from permutations
		data := NewCanRep()
		for pindex := 0; pindex < len(line.permutations); pindex++ {
			data.Add(line.permutations[pindex][i])
		}
		i++

		r := line.Cells[cindex].ProcessAbsent(data)
		result = append(result, r...)
	}
	return result
}

/*
Hidden Candidates
Check all candidates in line
count cells, containing it
if count is 1, it is hidden candidate
eliminate other candidates in the cell
*/
func (line *Line) ProcessHiddenCandidates() []Elimination {
	result := make([]Elimination, 0)
	line.GetFreeCandidates()
	for _, v := range line.freeCandidates {
		var celref *Cell = nil
		count := 0
		for _, cell := range line.Cells {
			if cell.DoesContain(v) {
				if count == 0 {
					celref = cell
					count = 1
				} else {
					count++
					break
				}
			}
		}
		if count == 1 {

			for num := range celref.list {
				if num != v {
					el := Elimination{
						Row:     celref.Row,
						Col:     celref.Column,
						Num:     num,
						Comment: "ProcessHiddenCandidates",
					}
					result = append(result, el)
				}
			}
		}
	}
	return result
}

/*
Naked Candidate
Finds all cells, containing only one candidate
Eliminate it from all other cells
Returns list of eliminations
*/
func (line *Line) ProcessNakedCandidates() []Elimination {
	result := make([]Elimination, 0)
	for _, cell := range line.Cells {

		if len(cell.CanRep.list) == 1 {
			num := cell.GetFirst()
			for _, other := range line.Cells {
				if cell.Row == other.Row && cell.Column == other.Column {
					continue
				}
				if !other.DoesContain(num) {
					continue
				}
				el := Elimination{
					Row:     other.Row,
					Col:     other.Column,
					Num:     num,
					Comment: fmt.Sprintf("Naked candidate in row %v col %v", cell.Row, cell.Column),
				}
				result = append(result, el)
			}
		}
	}
	return result
}

// func (line *Line) RemoveCandidate(num int) {
// 	// for _, cell := range line.Cells {
// 	// 	cell.RemoveCandidate(num)
// 	// }
// 	for i := 0; i < len(line.Cells); i++ {
// 		line.Cells[i].RemoveCandidate(num)
// 	}
// }
