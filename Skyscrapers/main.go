package main

import "fmt"

type Candidate struct {
	Row    int
	Col    int
	Number int
}

func main() {

	field := NewField([]int{1, 3, 2, 2, 2, 2, 1, 3, 2, 2, 1, 3, 2, 3, 2, 1})

	// field.ProcessCaseAlpha()
	// list := field.FindSoloCandidates()
	// for _, v := range list {
	// 	field.ProcessSolution(v.Row, v.Col, v.Number)
	// }

	ok := field.FindSolution()
	if ok {
		for row := 0; row < field.Size; row++ {
			if row > 0 {
				fmt.Print("/n")
			}
			for col := 0; col < field.Size; col++ {
				fmt.Print(field.Cells[col+row*field.Size].Solution)
			}
		}
	} else {
		fmt.Print("Solution does not found")
	}

	fmt.Print("Done")
}
