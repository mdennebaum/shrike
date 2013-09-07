#Shrike

Simple memory based key value store cheshire service.

##Install

To get started make sure you have the following installed

```
github.com/trendrr/goshire
```

Checkout or download shrike to a directory in your gopath.

##Config

Edit the config/config.yaml file to match with your settings.

##Run

Once you've edited the config.yaml to your liking run the following.

```
go run main.go -config=/path/to/config.yaml
```

Alternatively if your config folder is located in the same dir as the executable you can omit the -config flag

```
go run main.go
```

Then you should be able to point your browser at http://localhost:8010/ping to make sure its working

##More

For more info on how to use the goshire framework checkout https://github.com/trendrr/goshire. If there are any issues please submit an ticket.

