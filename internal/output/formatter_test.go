package output

import (
	"bytes"
	"strings"
	"testing"

	"github.com/mattn/go-runewidth"
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

func TestTableWriter_AlignsWideCharacters(t *testing.T) {
	var buf bytes.Buffer
	tw := NewTableWriter(&buf, []string{"NAME", "GENRE", "AREA"})
	tw.Row("Sushi Place", "寿司", "東京")
	tw.Row("ラーメン屋", "ラーメン", "大阪")
	tw.Flush()

	lines := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
	var want int
	for i, line := range lines {
		idx := strings.Index(line, []string{"AREA", "東京", "大阪"}[i])
		if idx < 0 {
			t.Fatalf("line %d missing last column: %q", i, line)
		}
		w := runewidth.StringWidth(line[:idx])
		if i == 0 {
			want = w
		} else if w != want {
			t.Fatalf("line %d: last column starts at display width %d, want %d\n%s", i, w, want, buf.String())
		}
	}
}
