package main

import (
	"flag"
	"github.com/mdennebaum/shrike/controllers"
	"github.com/trendrr/goshire/cheshire"
	"log"
)

func init() {
	flag.StringVar(&config, "config", "config/config.yaml", "path to the app config.yaml")
}

var config string

func main() {

	//parse the command line args
	flag.Parse()

	bootstrap := cheshire.NewBootstrapFile(config)

	//make sure the linker includes our controllers and runs inits
	controllers.Load()

	//tell everyone we started up
	log.Println("starting app with config: " + config)

	//starts listening on all configured interfaces
	bootstrap.Start()
}
