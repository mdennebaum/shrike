/*
Shrike is a simple memory based key value store goshire (https://github.com/trendrr/goshire) service.
Shrike uses the STREST protocol (https://github.com/trendrr/strest-server/wiki/STREST-Protocol-Spec)
and by default is configured to enable http, json and binary listeners.
*/
package main

import (
	"flag"
	"github.com/mdennebaum/shrike/controllers"
	"github.com/pmylund/go-cache"
	"github.com/trendrr/goshire/cheshire"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//init the app
func init() {
	//init to parse the command line args
	flag.StringVar(&config, "config", "config.yaml", "path to the app config.yaml")
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

	//load from data file and setup backup
	c = initData(c, bootstrap.Conf)

	//start the cache api
	controllers.StartApi(c)

	//tell everyone we started up
	log.Println("starting app with config: " + config)

	//starts listening on all configured interfaces
	bootstrap.Start()
}

//init cache dump on exit and reload on start
func initData(cache *cache.Cache, conf *cheshire.ServerConfig) *cache.Cache {

	//load cache dump data
	if data_dir, ok := conf.GetString("data_dir"); ok {

		//backup file name
		fileName := data_dir + "/shrike.db"

		//load cache backup data
		cache.LoadFile(fileName)

		//create a signal channel
		c := make(chan os.Signal, 1)

		//register for notifications on interrupt and sigterm
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		//watch for the exit sigs and dump cache to file
		go func() {
			for sig := range c {

				//print notice to stdout
				log.Println(sig)
				log.Println("saving to file")

				//dump to file
				cache.SaveFile(fileName)

				//exit process
				os.Exit(1)
			}
		}()
	}

	//return the loaded cache
	return cache
}
