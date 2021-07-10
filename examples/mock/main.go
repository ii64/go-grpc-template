package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"mime"
	"net"
	"net/http"
	"os"

	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/ii64/go-grpc-template/examples/mock/insecure"
	"github.com/ii64/go-grpc-template/examples/mock/server"
	_ "github.com/ii64/go-grpc-template/examples/mock/statik"
	"github.com/ii64/go-grpc-template/gen"

	fs "github.com/rakyll/statik/fs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

var (
	grpcPort    int = 10000
	gatewayPort int = 11000
)
var log grpclog.LoggerV2

func init() {
	flag.IntVar(&grpcPort, "grpc-port", grpcPort, "gRPC server port")
	flag.IntVar(&gatewayPort, "gateway-port", gatewayPort, "gRPC-gateway server")
	log = grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)
}

func main() {
	flag.Parse()

	addr := fmt.Sprintf(":%d", grpcPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
		grpc.UnaryInterceptor(grpc_validator.UnaryServerInterceptor()),
		grpc.StreamInterceptor(grpc_validator.StreamServerInterceptor()),
	)
	gen.RegisterMyServiceServer(s, server.New())
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	// grpc-gateway
	dialAddr := fmt.Sprintf("passthrough://localhost/%s", addr)
	conn, err := grpc.DialContext(
		context.Background(),
		dialAddr,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(insecure.CertPool, "")),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Failed to dial grpc server:", err)
	}

	// http mux
	mux := http.NewServeMux()

	gwMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{}),
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)
	err = gen.RegisterMyServiceHandler(
		context.Background(),
		gwMux,
		conn,
	)
	// err = gen.RegisterMyServiceHandlerFromEndpoint(
	// 	context.Background(),
	// 	gwMux,
	// 	addr,
	// 	[]grpc.DialOption{
	// 		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(insecure.CertPool, "")),
	// 		grpc.WithBlock(),
	// 	},
	// )
	if err != nil {
		log.Fatalln("Failed to register gateway server [MyService]:", err)
	}
	mux.Handle("/", gwMux)

	err = serveOpenAPI(mux)
	if err != nil {
		log.Fatalln("Failed to serve OpenAPI UI")
	}

	gatewayAddr := fmt.Sprintf("localhost:%d", gatewayPort)
	log.Infof("Serving gRPC-Gateway on https://%s", gatewayAddr)
	log.Infof("Serving OpenAPI Documentation on https://%s%s", gatewayAddr, "/openapi-ui/")
	gwServer := http.Server{
		Addr: gatewayAddr,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{insecure.Cert},
		},
		Handler: mux,
	}
	log.Fatalln(gwServer.ListenAndServeTLS("", ""))
}

func serveOpenAPI(mux *http.ServeMux) error {
	mime.AddExtensionType(".svg", "image/svg+xml")
	statikFS, err := fs.New()
	if err != nil {
		return err
	}
	fileServer := http.FileServer(statikFS)
	prefix := "/openapi-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
	return nil
}
