// blago is a Go blogging application centered around markdown.
package main

import (
	"flag"
	"github.com/golang/glog"
	"os"
)

var (
	Conf *Config
)

// command line flag
var (
	fConf = flag.String("c", "conf.json", "configuration file path")
)

func main() {
	flag.Parse()

	conffile, err := os.Open(*fConf)
	if err != nil {
		glog.Fatal(err)
	}
	Conf, err = ReadConfig(conffile)
	conffile.Close()
	if err != nil {
		glog.Fatal(err)
	}
}
