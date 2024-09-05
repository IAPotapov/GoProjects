package views

import (
	"Skyscrapers/models"
	"fmt"
	"strings"
)

func GetTable(field *models.Field) string {

	var style string = `<style type="text/css">
	.table{
	border-collapse: collapse;
	border-spacing: 0;
	table-layout: fixed;
	}
	.table tr{
	   height: 4em;
	}
	.table td{
	   text-align: center;
	   width: 4em;
	   font-size:large;
	}
	.table1{
	border-collapse: collapse;
	border-spacing: 0;
	}
	.table1 tr{
	   height: 4em;
	}
	.table1 td {
	   border: 1px solid;
	   text-align: center;
	   width: 4em;
	   font-size:large;
	}
	.table2{
	background-color: beige;
	border-collapse: collapse;
	border-spacing: 0;
	}
	.table2 tr{
	   height: 1.5em;
	}
	.table2 td{
	border: 1px solid;
	text-align: center;
	font-size: small;
	width: 1.5em;
	}
	.btn {
		display: inline-block;
		font-weight: 400;
		line-height: 1.5;
		color: #212529;
		text-align: center;
		text-decoration: none;
		vertical-align: middle;
		cursor: pointer;
		-webkit-user-select: none;
		-moz-user-select: none;
		user-select: none;
		background-color: transparent;
		border: 1px solid transparent;
		padding: .375rem .75rem;
		font-size: 1rem;
		border-radius: .25rem;
		transition: color .15s ease-in-out, background-color .15s ease-in-out, border-color .15s ease-in-out, box-shadow .15s ease-in-out
	}
	.btn-primary {
		color: #fff;
		background-color: #0d6efd;
		border-color: #0d6efd
	}
 </style>`
	var sb strings.Builder

	sb.WriteString(style)

	sb.WriteString("<table class=\"table\">")
	// first row
	sb.WriteString("<tr>")
	// left corner
	//sb.WriteString("<td style=\"background-color: lightcyan;\">&nbsp;</td>")
	sb.WriteString("<td>&nbsp;</td>")
	// top clue table in td
	sb.WriteString("<td style=\"width: 16em;\">")
	sb.WriteString("<table class=\"table\">")
	sb.WriteString("<tr>")
	for i := 0; i < field.Size; i++ {
		sb.WriteString(fmt.Sprintf("<td>%v</td>", field.Clues[i]))
	}
	sb.WriteString("</tr>")
	sb.WriteString("</table>")
	sb.WriteString("</td>")
	// right corner
	//sb.WriteString("<td style=\"background-color: lightcyan;\">&nbsp;</td>")
	sb.WriteString("<td>&nbsp;</td>")
	sb.WriteString("</tr>")

	// middle row
	sb.WriteString("<tr style=\"height: 16em;\">")

	// left clue table in td
	sb.WriteString("<td>")
	sb.WriteString("<table class=\"table\">")
	for i := 0; i < field.Size; i++ {
		j := len(field.Clues) - 1 - i
		sb.WriteString(fmt.Sprintf("<tr><td>%v</td></tr>", field.Clues[j]))
	}
	sb.WriteString("</table>")
	sb.WriteString("</td>")

	// lines table in td
	sb.WriteString("<td>")
	sb.WriteString("<table class=\"table1\">")
	for i := 0; i < field.Size; i++ {
		sb.WriteString(GetLineRow(&field.Lines[i]))
	}
	sb.WriteString("</table>")
	sb.WriteString("</td>")

	// right clue table in td
	sb.WriteString("<td>")
	sb.WriteString("<table class=\"table\">")
	for i := 0; i < field.Size; i++ {
		j := field.Size + i
		sb.WriteString(fmt.Sprintf("<tr><td>%v</td></tr>", field.Clues[j]))
	}
	sb.WriteString("</table>")
	sb.WriteString("</td>")
	sb.WriteString("</tr>")

	// bottom row (same as top)
	sb.WriteString("<tr>")
	// left corner
	//sb.WriteString("<td style=\"background-color: lightcyan;\">&nbsp;</td>")
	sb.WriteString("<td>&nbsp;</td>")
	// bottom clue table in td
	sb.WriteString("<td>")
	sb.WriteString("<table class=\"table\">")
	sb.WriteString("<tr>")
	for i := 0; i < field.Size; i++ {
		j := len(field.Clues) - 1 - field.Size - i
		sb.WriteString(fmt.Sprintf("<td>%v</td>", field.Clues[j]))
	}
	sb.WriteString("</tr>")
	sb.WriteString("</table>")
	sb.WriteString("</td>")
	// right corner
	//sb.WriteString("<td style=\"background-color: lightcyan;\">&nbsp;</td>")
	sb.WriteString("<td>&nbsp;</td>")
	sb.WriteString("</tr>")
	sb.WriteString("</table>")

	//sb.WriteString("<a class=\"btn btn-primary\" href=\"/next\">NEXT</a>")

	return sb.String()
}
