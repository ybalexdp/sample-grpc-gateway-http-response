package main

import (
	"context"
	"flag"
	"net/http"

	"sample-grpc-gateway-http-response/service/pb"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func run() error {
	endpoint := flag.String("endpoint", "api:2007", "API Service")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	dualOpts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterSampleHandlerFromEndpoint(ctx, mux, *endpoint, dualOpts); err != nil {
		return err
	}
	return http.ListenAndServe(":2008", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
