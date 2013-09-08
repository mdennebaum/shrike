package main

import (
	"flag"
	"github.com/mdennebaum/shrike/controllers"
	"github.com/pmylund/go-cache"
	"github.com/trendrr/goshire/cheshire"
	"log"
	"time"
)

//init the app
func init() {
	//init to parse the command line args
	flag.StringVar(&config, "config", "config/config.yaml", "path to the app config.yaml")
}

//string to hold the config path
var config string

//app entry point
func main() {

	//parse the command line args
	flag.Parse()

	//setup a bootstrap
	bootstrap := cheshire.NewBootstrapFile(config)

	//create new cache with default timeout durration
	c := cache.New(5*time.Minute, 30*time.Second)

	//start the cache api
	controllers.StartApi(c)

	//tell everyone we started up
	log.Println("starting app with config: " + config)

	//starts listening on all configured interfaces
	bootstrap.Start()
}
