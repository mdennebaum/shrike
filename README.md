This is a starter app for the [goshire](https://github.com/trendrr/goshire) web framework. You can use this as a boilerplate when starting a new goshire project. 

##Install

To get started make sure you have the following installed

```
go get github.com/kylelemons/go-gypsy/yaml
go get github.com/hoisie/mustache
go get code.google.com/p/go.net/websocket
go get github.com/trendrr/goshire
```

Checkout or download wildling to a directory in your gopath.

##Config

Edit the config/config.yaml file to match with your settings.

##Run

Once you've edited the config.yaml to your liking run the following. 

```
go run wildling.go -config=/path/to/config.yaml
```

Alternatively if your config folder is located in the same dir as the executable you can omit the -config flag

```
go run wildling.go
```

Then you should be able to point your browser at http://localhost:8000/ and http://localhost:8000/api to see examples of html controller output and api controller output. 

##More

For more info on how to use the goshire framework checkout https://github.com/trendrr/goshire. If there are any issues please submit an ticket. 

