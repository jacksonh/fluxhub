package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	configPath := ""

	flag.StringVar(&configPath, "config", "", "Path to the config file.")
	flag.Parse()

	if len(configPath) < 1 {
		log.Println("config argument must be specified.")
		os.Exit(-1)
	}

	yaml, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Println("Unable to read file specified by config flag.")
		panic(err)
	}

	config, err := parseConfig(yaml)
	if err != nil {
		panic(err)
	}
	runDocker(config.Jobs[0].Docker)
}
