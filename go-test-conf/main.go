package main

import (
	"./configuration"
	"flag"
	"fmt"
	"log"
)

func main() {
	fmt.Println("test")

	config := flag.String("config", "configuration/config.gcfg", "Configuration file")
	flag.Parse()
	conf, err := configuration.NewConfiguration(*config)

	if nil != err {
		log.Fatal("Failed to load configuration", err)
	}

	fmt.Println(conf)
}
