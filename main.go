package main

import (
	"flag"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "c", "./config/config.yaml", "config file")
	flag.Parse()
	Init(configFile)
	aliddns()
}
