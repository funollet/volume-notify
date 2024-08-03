package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	if r, _ := run("echo hi"); r != "hi\n" {
		t.Error("unexpected output")
	}

	if _, err := run("echo hi"); err != nil {
		t.Error("unexpected exit value")
	}
}
