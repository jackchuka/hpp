package output

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/mattn/go-runewidth"
)

type TableWriter struct {
	out  io.Writer
	rows [][]string
}

func NewTableWriter(out io.Writer, headers []string) *TableWriter {
	return &TableWriter{out: out, rows: [][]string{headers}}
}

func (t *TableWriter) Row(values ...string) {
	t.rows = append(t.rows, values)
}

func (t *TableWriter) Flush() {
	var widths []int
	for _, row := range t.rows {
		for i, v := range row {
			if i >= len(widths) {
				widths = append(widths, 0)
			}
			widths[i] = max(widths[i], runewidth.StringWidth(v))
		}
	}
	for _, row := range t.rows {
		var b strings.Builder
		for i, v := range row {
			if i == len(row)-1 {
				b.WriteString(v)
				break
			}
			b.WriteString(runewidth.FillRight(v, widths[i]+2))
		}
		_, _ = fmt.Fprintln(t.out, b.String())
	}
	t.rows = nil
}

func WriteJSON(out io.Writer, data interface{}) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}
