package controllers

import (
	"github.com/pmylund/go-cache"
	"github.com/trendrr/goshire/cheshire"
	"time"
)

//shared cache object
var c *cache.Cache

//register controllers
func StartApi(ca *cache.Cache) {

	//init shared cache object
	c = ca

	//register controllers
	cheshire.RegisterApi("/ping", "GET", Ping)
	cheshire.RegisterApi("/get", "GET", Get)
	cheshire.RegisterApi("/set", "POST", Set)
	cheshire.RegisterApi("/delete", "DELETE", Delete)
}

// a demo Ping controller function
func Ping(txn *cheshire.Txn) {

	//create new response for this transaction
	response := cheshire.NewResponse(txn)

	//add a simple pong message to the data packet
	response.Put("data", "PONG")

	//write the response
	txn.Write(response)
}

//get a value from the cache service
func Get(txn *cheshire.Txn) {

	//create a new response for this transaction
	response := cheshire.NewResponse(txn)

	//get the key for the requested value
	if key, ok := txn.Params().GetString("key"); ok {

		//try to get the val from the cache
		if cacheVal, found := c.Get(key); found {

			//add the val to the response data
			response.Put("data", cacheVal.(string))
		} else {

			//value not found. set status and message
			response.SetStatus(404, "key not found")
		}

	} else {

		//no key param set. set error status and message
		response.SetStatus(400, "missing key param")
	}

	//write the response
	txn.Write(response)
}

//setup a value in the cache service
func Set(txn *cheshire.Txn) {

	//create a new response for this transaction
	response := cheshire.NewResponse(txn)

	//set timeout. pass -1 to cache indefinatly. defaults to (24 hours)
	timeout, err := time.ParseDuration(txn.Params().MustString("timeout", "0"))

	if err != nil {
		timeout = 0
	}

	//get the key param
	if key, kok := txn.Params().GetString("key"); kok {

		//get the val param
		if val, vok := txn.Params().GetString("val"); vok {

			//set the val for given key with given timeout in cache
			c.Set(key, val, timeout)
		} else {

			//set error status for missing val
			response.SetStatus(400, "missing val param")
		}
	} else {

		//set error status for missing key
		response.SetStatus(400, "missing key param")
	}

	//write the response
	txn.Write(response)
}

// a demo Ping controller function
func Delete(txn *cheshire.Txn) {

	//create a new response for this transaction
	response := cheshire.NewResponse(txn)

	//get the key for the requested value
	if key, ok := txn.Params().GetString("key"); ok {

		//delete the value for the passed key
		c.Delete(key)
	} else {

		//set error status for missing key
		response.SetStatus(400, "missing key param")
	}

	//write the response
	txn.Write(response)
}
