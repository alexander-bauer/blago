package main

import (
	"encoding/json"
	"io"
)

type Config struct {
	// Author is the full or screen name of the author of this blog.
	Author string

	// DataDir is the directory which is to contain JSON-encoded blog
	// posts.
	DataDir string
}

// ReadConfig decodes a JSON-encoded configuration file from the given
// io.Reader.
func ReadConfig(r io.Reader) (*Config, error) {
	c := &Config{}
	err := json.NewDecoder(r).Decode(c)
	return c, err
}
