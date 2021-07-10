module github.com/ii64/go-grpc-template

go 1.16

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.39.0

require (
	github.com/gogo/protobuf v1.3.0
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0
	github.com/mwitkow/go-proto-validators v0.3.2
	golang.org/x/net v0.0.0-20210503060351-7fd8e65b6420 // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	google.golang.org/api v0.30.0
	google.golang.org/genproto v0.0.0-20210624195500-8bfb893ecb84
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
)
