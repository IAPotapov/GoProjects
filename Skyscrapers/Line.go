package main

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
		if c.Solution != 0 {
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
			if line.Cells[j].Solution > 0 {
				data[j] = line.Cells[j].Solution
			} else {
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
Check coppect permutations for "candidate elimination"
Otherwise return empty list

Should be executed before call:
line.GetPermutations()
line.RemoveIncorrectPermutations()

if len(line.permutations) == 0 { ok = false; return; }
*/
func (line *Line) ProcessCaseAlpha() (done bool) {
	done = false

	var i int = 0
	for cindex := 0; cindex < len(line.Cells); cindex++ {
		if line.Cells[cindex].Solution > 0 {
			continue
		}
		// data is the list of candidates for cell from permutations
		data := NewCanRep()
		for pindex := 0; pindex < len(line.permutations); pindex++ {
			//data[pindex] = line.permutations[pindex][i]
			data.Add(line.permutations[pindex][i])
		}
		i++
		//e := line.Cells[cindex].GetEliminations(data)
		//result = append(result, e...)

		// if !line.Cells[cindex].Equals(data) {
		// 	ok = true
		// 	line.Cells[cindex].SetCandidates(data)
		// }
		if line.Cells[cindex].RemoveAbsent(data) {
			done = true
		}
	}
	return
}

// finds both naked and hidden candidates
func (line *Line) FindSoloCandidates() []Candidate {
	result := make([]Candidate, 0)
	line.GetFreeCandidates()
	for _, v := range line.freeCandidates {
		row := 0
		col := 0
		num := 0
		count := 0
		for _, cell := range line.Cells {
			if cell.DoesContain(v) {
				count++
				if count > 1 {
					break
				} else {
					num = v
					row = cell.Row
					col = cell.Column
				}
			}
		}
		if count == 1 {
			result = append(result, Candidate{Row: row, Col: col, Number: num})
		}
	}
	return result
}

func (line *Line) RemoveCandidate(num int) {
	// for _, cell := range line.Cells {
	// 	cell.RemoveCandidate(num)
	// }
	for i := 0; i < len(line.Cells); i++ {
		line.Cells[i].RemoveCandidate(num)
	}
}
