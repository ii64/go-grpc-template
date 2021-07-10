package server

import (
	"context"
	"log"

	"github.com/ii64/go-grpc-template/gen"
)

type Backend struct {
	gen.MyServiceServer
}

func New() *Backend {
	return &Backend{}
}

func (b *Backend) DoEcho(ctx context.Context, in *gen.EchoRequest) (out *gen.EchoResponse, err error) {
	log.Printf("%+#v", in)
	out = &gen.EchoResponse{
		MType:   in.MType,
		Message: in.Message,
	}
	return
}
