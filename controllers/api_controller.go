package controllers

import (
	"github.com/mdennebaum/shrike/cache"
	"github.com/trendrr/goshire/cheshire"
)

//shared cache object
var cache *cache.Cache

//register controllers
func StartApi(c *cache.Cache) {

	//init shared cache object
	cache = c

	//register the ping controller.
	cheshire.RegisterApi("/ping", "GET", Ping)
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
		if cacheVal, found := cache.GetString(key); found {

			//add the val to the response data
			response.Put("data", cacheVal)
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

	response.Put("data", "PONG")
	txn.Write(response)
}

// a demo Ping controller function
// func Delete(txn *cheshire.Txn) {
// 	response := cheshire.NewResponse(txn)
// 	response.Put("data", "PONG")
// 	txn.Write(response)
// }
