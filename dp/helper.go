package dp

import (
	"fmt"
	"strings"
)

type table struct {
	rows    string
	columns string
	data    [][]byte
}

func newTable(row, column string) *table {
	data := make([][]byte, len(row)+1)
	for i := range data {
		data[i] = make([]byte, len(column)+1)
	}

	return &table{
		row,
		column,
		data,
	}
}

func (t *table) String() string {
	var b strings.Builder

	rows := len(t.data)
	columns := len(t.data[0])

	for i := -1; i < rows; i++ {
		for j := -1; j < columns; j++ {
			switch i {
			// columns header
			case -1:
				if j <= 0 {
					b.WriteRune('\t')
				} else {
					b.WriteString(fmt.Sprintf("\t%c", t.columns[j-1]))
				}

			case 0:
				if j < 0 {
					b.WriteRune('\t')
				} else {
					b.WriteString(fmt.Sprintf("\t%v", t.data[i][j]))
				}

			default:
				if j < 0 {
					// rows header
					b.WriteString(fmt.Sprintf("\t%c", t.rows[i-1]))
				} else {
					b.WriteString(fmt.Sprintf("\t%v", t.data[i][j]))
				}
			}
		}
		b.WriteRune('\n')
	}

	return b.String()
}
