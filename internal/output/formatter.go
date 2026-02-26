package output

import (
	"encoding/json"
	"fmt"
	"io"
	"text/tabwriter"
)

type TableWriter struct {
	w *tabwriter.Writer
}

func NewTableWriter(out io.Writer, headers []string) *TableWriter {
	w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
	for i, h := range headers {
		if i > 0 {
			_, _ = fmt.Fprint(w, "\t")
		}
		_, _ = fmt.Fprint(w, h)
	}
	_, _ = fmt.Fprintln(w)
	return &TableWriter{w: w}
}

func (t *TableWriter) Row(values ...string) {
	for i, v := range values {
		if i > 0 {
			_, _ = fmt.Fprint(t.w, "\t")
		}
		_, _ = fmt.Fprint(t.w, v)
	}
	_, _ = fmt.Fprintln(t.w)
}

func (t *TableWriter) Flush() {
	_ = t.w.Flush()
}

func WriteJSON(out io.Writer, data interface{}) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}
