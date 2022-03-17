module greeter-service

go 1.16

require (
	github.com/asim/go-micro/plugins/broker/rabbitmq/v4 v4.0.0-20220317022205-c6d352c83291
	github.com/asim/go-micro/plugins/registry/etcd/v4 v4.0.0-20220317022205-c6d352c83291
	github.com/asim/go-micro/plugins/transport/nats/v4 v4.0.0-20220317022205-c6d352c83291
	go-micro.dev/v4 v4.6.0
	google.golang.org/protobuf v1.26.0
)
