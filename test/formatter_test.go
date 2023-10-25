package paf

import (
	"testing"
	paf "github.com/44/paf/src"
)

func TestFormatGrep(t *testing.T) {
	formatted := paf.FormatGrep("file.txt:10: text")
	if formatted != "file.txt 10  text" {
		t.Fatalf("Expected 'file.txt 10  text', got '%s'", formatted)
	}
}
