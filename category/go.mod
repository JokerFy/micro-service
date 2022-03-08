module micro-service/category

go 1.16

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.5.2
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/config/source/consul/v2 v2.9.1 // indirect
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1 // indirect
	github.com/micro/micro/v3 v3.9.0
	github.com/prometheus/common v0.6.0
	google.golang.org/protobuf v1.27.1
)
