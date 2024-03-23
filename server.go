package main

import (
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
  "server/services"
)

func main()  {
  s := grpc.NewServer()

  listener, err := net.Listen("tcp", ":50051")
  if err != nil{
    log.Fatal(err)
  }

  services.RegisterGreetingServer(s, services.NewGreetingServer())

  fmt.Println("gRPC server listening on port 50051")

  err = s.Serve(listener)
  if err != nil{
    log.Fatal(err)
  }
}
