package main

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestPost_WriteHTML(t *testing.T) {
	p := &Post{
		Title:    "Test Post",
		Body:     "This is a test post.\n\nIt has multiple paragraphs, *italics*, **bold**, does_not_emphasize_mid-word, and includes `code-blocks`.",
		Created:  time.Now(),
		Modified: time.Now(),
		ID:       1,
	}
	buf := bytes.NewBuffer([]byte{})

	if err := p.WriteHTML(buf); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", buf)
}
