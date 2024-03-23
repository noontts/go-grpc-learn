package services

import (
	context "context"
	"fmt"
)

type greetingServer struct {
  
}

func NewGreetingServer() GreetingServer {
  return greetingServer{}
}

func (greetingServer) mustEmbedUnimplementedGreetingServer(){}

func (greetingServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
  result := fmt.Sprintf("Hello %v", req.Name)
  res := HelloResponse{
    Result: result,
  }
  return &res, nil
}
