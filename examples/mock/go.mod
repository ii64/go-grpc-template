module github.com/ii64/go-grpc-template/examples/mock

go 1.16

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/ii64/go-grpc-template v0.0.0-00010101000000-000000000000
	github.com/nametake/protoc-gen-gohttp v1.5.0 // indirect
	github.com/rakyll/statik v0.1.7
	google.golang.org/genproto v0.0.0-20210708141623-e76da96a951f // indirect
	google.golang.org/grpc v1.39.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.39.0

replace github.com/ii64/go-grpc-template => ../../
