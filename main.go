package main

import (
	"flag"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "c", "./aliddns.yaml", "config file")
	flag.Parse()
	Init(configFile)
	aliddns()
}
