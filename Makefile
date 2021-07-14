# all:make install build

generate:
# Generate gogo, gRPC-Gateway, swagger, go-validators output.
#
# -I declares import folders, in order of importance
# This is how proto resolves the protofile imports.
# It will check for the protofile relative to each of these
# folders and use the first one it finds.
#
# --gogo_out generates GoGo Protobuf output with gRPC plugin enabled.
# --grpc-gateway_out generates gRPC-Gateway output.
# --swagger_out generates an OpenAPI 2.0 specification for our gRPC-Gateway endpoints.
# --govalidators_out generates Go validation files for our messages types, if specified.
#
# The lines starting with Mgoogle/... are proto import replacements,
# which cause the generated file to import the specified packages
# instead of the go_package's declared by the imported protof files.
#
# ./proto is the output directory.
#
# proto/example.proto is the location of the protofile we use.
	protoc \
		-I proto \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway/v2/ \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway/ \
		-I proto/googleapis/googleapis/ \
		-I vendor/ \
		--go_out=paths=source_relative:./gen/ \
		--gohttp_out=paths=source_relative:./gen/ \
		--go-grpc_out=paths=source_relative:./gen/ \
		--grpc-gateway_out=allow_patch_feature=false,paths=source_relative:./gen/ \
		--swagger_out=third_party/OpenAPI/ \
		--govalidators_out=gogoimport=false,paths=source_relative:./gen \
		--gohttpclient_out=paths=source_relative:./gen/ \
		./proto/service.proto

#	Workaround for https://github.com/grpc-ecosystem/grpc-gateway/issues/229
	sed -i.bak "s/empty.Empty/types.Empty/g" gen/service.pb.gw.go && rm gen/service.pb.gw.go.bak || true

#	 Generate static assets for OpenAPI UI
	statik -m -f -dest examples/mock/ -src third_party/OpenAPI/

install:
	go get \
		google.golang.org/grpc \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
		github.com/mwitkow/go-proto-validators/protoc-gen-govalidators \
		github.com/ii64/protoc-gen-gohttpclient \
		github.com/ii64/protoc-gen-gohttpclient/lib \
		github.com/rakyll/statik

build:
	go build -v -x example


mock:
	go build -v -x example/mock