package controllers

import (
	"github.com/mdennebaum/shrike/cache"
	"github.com/trendrr/goshire/cheshire"
)

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
	response := cheshire.NewResponse(txn)
	response.Put("data", "PONG")
	txn.Write(response)
}

//get a value from the queue service
func Get(txn *cheshire.Txn) {

	//get the key for the requested value
	key, ok := txn.Params().GetString("key")

	var cacheVal string

	if ok {
		cacheVal, ok = cache.GetString(key)
	}

	response := cheshire.NewResponse(txn)
	response.Put("data", cacheVal)
	txn.Write(response)
}

// a demo Ping controller function
func Set(txn *cheshire.Txn) {
	response := cheshire.NewResponse(txn)
	response.Put("data", "PONG")
	txn.Write(response)
}

// a demo Ping controller function
func Delete(txn *cheshire.Txn) {
	response := cheshire.NewResponse(txn)
	response.Put("data", "PONG")
	txn.Write(response)
}
