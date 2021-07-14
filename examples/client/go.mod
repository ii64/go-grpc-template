module github.com/ii64/go-grpc-template/examples/client

go 1.16

require (
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/ii64/go-grpc-template v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.39.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.39.0

replace github.com/ii64/go-grpc-template => ../../
