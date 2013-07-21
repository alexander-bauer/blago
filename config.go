package main

import (
	"encoding/json"
	"io"
)

type Config struct {
	Database struct {
		// DriverName is the basename of a supported database driver,
		// such as sqlite3 or mysql.
		DriverName string

		// Resource is the string that will be passed to the driver in
		// order to connect to the database. For sqlite3, this would
		// be a path to a file.
		Resource string
	}
}

// ReadConfig decodes a JSON-encoded configuration file from the given
// io.Reader.
func ReadConfig(r io.Reader) (*Config, error) {
	c := &Config{}
	err := json.NewDecoder(r).Decode(c)
	return c, err
}
