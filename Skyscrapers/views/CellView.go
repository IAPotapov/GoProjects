package views

import (
	"Skyscrapers/models"
	"strings"
)

func GetCellTableData(cell *models.Cell) string {
	var sb strings.Builder
	var s string
	// if cell.Solution > 0 {
	// 	s = fmt.Sprintf("<td>%v</td>", cell.Solution)
	// } else {
	// 	s = GetCanRepTableData(cell.CanRep)
	// }
	s = GetCanRepTableData(cell.CanRep)
	sb.WriteString(s)
	return sb.String()
}
