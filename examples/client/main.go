package main

import (
	"context"
	"io/ioutil"
	"os"
	"time"

	"github.com/ii64/go-grpc-template/examples/insecure"
	"github.com/ii64/go-grpc-template/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

var (
	grpcServer = "localhost:10000"
)

var log grpclog.LoggerV2

func init() {
	log = grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)
}

func main() {
	conn, err := grpc.Dial(
		grpcServer,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(insecure.CertPool, "")),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Failed to dial grpc server:", err)
	}
	defer conn.Close()

	client := gen.NewMyServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.DoEcho(ctx, &gen.EchoRequest{
		MType:   gen.MessageType_TEXT,
		Message: "hellox",
	})
	if err != nil {
		log.Fatal("RPC call got error:", err)
	}
	log.Infof("Response: %+#v", r)
}
