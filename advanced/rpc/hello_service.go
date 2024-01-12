package rpc

import (
	"fmt"
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface interface {
	Hello(string, *string) error
}

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = fmt.Sprintf("hello %s", request)
	return nil
}

func RegisterHelloService(server *rpc.Server, svc HelloServiceInterface) error {
	return server.RegisterName(HelloServiceName, svc)
}

type client struct {
	c *rpc.Client
}

func (c client) Hello(request string, reply *string) error {
	return c.c.Call(HelloServiceName+".Hello", request, reply)
}

func NewHelloServiceClient(cc *rpc.Client) HelloServiceInterface {
	return &client{cc}
}
