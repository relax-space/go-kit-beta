package main

import "github.com/alecthomas/kingpin"

var Version string

/*
$ go build -ldflags "-X main.Version=v2.0"
$ ./version -v
v2.0

**/

func main() {
	if len(Version) == 0 {
		Version = "v1.0"
	}
	kingpin.CommandLine.Version(Version).VersionFlag.Short('v')
	kingpin.Parse()
}
