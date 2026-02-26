package output

import (
	"bytes"
	"strings"
	"testing"
)

func TestTableWriter(t *testing.T) {
	var buf bytes.Buffer
	tw := NewTableWriter(&buf, []string{"NAME", "GENRE", "AREA"})
	tw.Row("Sushi Place", "寿司", "東京")
	tw.Row("Ramen Shop", "ラーメン", "大阪")
	tw.Flush()

	out := buf.String()
	if !strings.Contains(out, "NAME") {
		t.Fatal("expected header NAME")
	}
	if !strings.Contains(out, "Sushi Place") {
		t.Fatal("expected Sushi Place")
	}
}

func TestJSONOutput(t *testing.T) {
	var buf bytes.Buffer
	data := map[string]string{"name": "test"}
	if err := WriteJSON(&buf, data); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), `"name"`) {
		t.Fatal("expected JSON output")
	}
}
