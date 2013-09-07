package controllers

import (
	// "github.com/mdennebaum/shrike/cache"
	"github.com/trendrr/goshire/cheshire"
)

//register controllers
func init() {

	//register the ping controller.
	cheshire.RegisterApi("/ping", "GET", Ping)

}

// a demo Ping controller function
func Ping(txn *cheshire.Txn) {
	response := cheshire.NewResponse(txn)
	response.Put("data", "PONG")
	txn.Write(response)
}

// a demo Ping controller function
func Get(txn *cheshire.Txn) {
	response := cheshire.NewResponse(txn)
	response.Put("data", "PONG")
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
