package main

import (
	gs "github.com/madari/goskirt"
	"io"
	"time"
)

type Post struct {
	Title string
	Body  string

	Created, Modified time.Time
	ID                int
}

var GS = gs.Goskirt{
	Extensions: gs.EXT_NO_INTRA_EMPHASIS | gs.EXT_TABLES |
		gs.EXT_FENCED_CODE | gs.EXT_STRIKETHROUGH |
		gs.EXT_SUPERSCRIPT | gs.EXT_LAX_SPACING,
}

// WriteHTML creates an HTML version of the Post from markdown via
// Goskirt and returns any errors.
func (p *Post) WriteHTML(w io.Writer) error {
	_, err := GS.WriteHTML(w, []byte(
		"# "+p.Title+" #\n"+p.Body))
	return err
}
