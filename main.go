package main

import (
	"log"
	"net"

	app "sample-grpc-gateway-http-response/service/application"

	"sample-grpc-gateway-http-response/service/pb"

	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":2007")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()

	s := &app.SampleService{}
	pb.RegisterSampleServer(server, s)
	server.Serve(listenPort)
}
