package views

import (
	"Skyscrapers/models"
	"fmt"
	"strings"
)

func GetCanRepTableData(rep *models.CanRep) string {
	var n int = 2
	var i int = 1

	var sb strings.Builder

	sb.WriteString("<td><table class=\"table2\" align=\"center\">")
	for row := 1; row <= n; row++ {
		sb.WriteString("<tr>")
		for col := 1; col <= n; col++ {
			if rep.DoesContain(i) {
				sb.WriteString(fmt.Sprintf("<td>%v</td>", i))
			} else {
				sb.WriteString("<td>&nbsp;</td>")
			}
			i++
		}
		sb.WriteString("</tr>")
	}
	sb.WriteString("</table></td>")

	return sb.String()
}
