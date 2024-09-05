package views

import (
	"Skyscrapers/models"
	//"fmt"
	"strings"
)

func GetLineRow(line *models.Line) string {
	n := len(line.Cells)
	var sb strings.Builder
	sb.WriteString("<tr>")
	for i := 0; i < n; i++ {
		sb.WriteString(GetCellTableData(line.Cells[i]))
	}
	sb.WriteString("</tr>")

	return sb.String()
}
