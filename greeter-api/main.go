package main

import (
	_ "github.com/asim/go-micro/plugins/broker/rabbitmq/v4"
	_ "github.com/asim/go-micro/plugins/registry/etcd/v4"
	_ "github.com/asim/go-micro/plugins/transport/nats/v4"
	"greeter-api/client"
	"greeter-api/router"
	"log"
)

func main() {
	// Remember to call the Init() function to initialize the go-micro client service
	client.Init()

	// Start Gin Router at port 3000
	r := router.NewRouter()
	if err := r.Run("0.0.0.0:3000"); err != nil {
		log.Print(err.Error())
	}
}

//curl "http://localhost:3000/greeter?name=test"
