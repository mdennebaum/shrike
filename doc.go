/*
Shrike

Simple memory based key value store goshire service. Shrike uses the STREST protocal and by default
is configured to enable http, json and binary listeners.

Persistance:

Shrike is primarily an in memory store but by setting a data_dir value in your config file, Shrike
will dump all data on interupt or sigterm to a backup file which it will load back in next time you
start it.

Install:

go get github.com/pmylund/go-cache
go get github.com/trendrr/goshire
go get github.com/mdennebaum/shrike


Checkout or download shrike to a directory in your gopath.


Config:

Edit the config.yaml file to match with your settings.


Run:

Once you've edited the config.yaml to your liking run the following.

go run main.go -config=/path/to/config.yaml

Alternatively if your config folder is located in the same dir as the executable you can omit the
-config flag

go run main.go

Then you should be able to point your browser at http://localhost:8010/ping to make sure its working


Client:
Goshire comes with a bundled client that supports the binary and json protocols.


More:

For more info on how to use the goshire framework checkout https://github.com/trendrr/goshire.
If there are any issues please submit an ticket.
*/
package main
