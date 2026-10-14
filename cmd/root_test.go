package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestRuntimeErrorPrintedOnceWithoutUsage(t *testing.T) {
	t.Setenv("HOTPEPPER_API_KEY", "")
	var stdout, stderr bytes.Buffer
	if code := run([]string{"budget"}, &stdout, &stderr); code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	out := stdout.String() + stderr.String()
	if n := strings.Count(out, "HOTPEPPER_API_KEY environment variable is required"); n != 1 {
		t.Fatalf("error message printed %d times:\n%s", n, out)
	}
	if strings.Contains(out, "Usage:") {
		t.Fatalf("usage printed for runtime error:\n%s", out)
	}
}

func TestFlagErrorStillShowsUsage(t *testing.T) {
	budgetCmd.SilenceUsage = false
	var stdout, stderr bytes.Buffer
	if code := run([]string{"budget", "--no-such-flag"}, &stdout, &stderr); code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	out := stdout.String() + stderr.String()
	if n := strings.Count(out, "unknown flag: --no-such-flag"); n != 1 {
		t.Fatalf("error message printed %d times:\n%s", n, out)
	}
	if !strings.Contains(out, "Usage:") {
		t.Fatalf("expected usage for flag error:\n%s", out)
	}
}
