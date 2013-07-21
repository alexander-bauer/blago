// blago is a Go blogging application centered around markdown.
package main

import (
	"flag"
	"github.com/golang/glog"
	"os"
)

var (
	Version = "0.0"

	Conf *Config
)

// command line flag
var (
	fConf = flag.String("c", "conf.json", "configuration file path")
)

func main() {
	flag.Parse()

	glog.Infof("blago version %s\n", Version)

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
