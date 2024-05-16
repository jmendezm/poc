package main

import (
	"flag"
	"technical_test/config"
	"technical_test/server"
)

func main() {
	configFilePath := flag.String("config", "./config/config.yml", "configuration file path")
	flag.Parse()
	config.Init(*configFilePath)
	server.Start()
}
