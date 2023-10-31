package protobuf

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"
)

type greetServiceServerImpl struct {
	UnimplementedGreetServiceServer
}

func NewGreetServiceServer() GreetServiceServer {
	return &greetServiceServerImpl{}
}

func (g *greetServiceServerImpl) Hello(ctx context.Context, request *Request) (*Response, error) {
	return g.processReq(request)
}

func (g *greetServiceServerImpl) HelloStream(server GreetService_HelloStreamServer) error {
	reqCh := make(chan *Request)
	errCh := make(chan error)
	go recvReq(server, reqCh, errCh)

	for {
		select {
		case err := <-errCh:
			return g.processErr(err)
		case req := <-reqCh:
			log.Printf("Recv request %s\n", req)
			resp, err := g.processReq(req)
			if err != nil {
				return fmt.Errorf("processReq req[%s] err: %s", req, err)
			}
			log.Printf("Send Response %s\n", resp)
			server.Send(resp)
		}
	}
}

func (g *greetServiceServerImpl) processErr(err error) error {
	if err != io.EOF {
		return fmt.Errorf("recv req err: %s", err)
	}
	return nil
}

func (g *greetServiceServerImpl) processReq(req *Request) (*Response, error) {
	resp := Response{}
	resp.Id = req.Id
	switch req.Type {
	case Type_PING:
		resp.Type = Type_PONG
		resp.Data = []byte("pong!")
	case Type_PONG:
		resp.Type = Type_NORMAL
		resp.Data = []byte(magic(nil))
	case Type_NORMAL:
		resp.Type = Type_NORMAL
		resp.Data = []byte(magic(req.Data))
	}
	return &resp, nil
}

func recvReq(server GreetService_HelloStreamServer, reqCh chan<- *Request, errCh chan<- error) {
	for {
		req, err := server.Recv()
		if err != nil {
			errCh <- err
			return
		}
		select {
		case <-server.Context().Done():
			return
		case reqCh <- req:
		}
	}
}

func magic(b []byte) (ret string) {
	if b == nil {
		return "Hello, Grpc!"
	}
	ret = string(b)
	ret = strings.ReplaceAll(ret, "吗", "")
	ret = strings.ReplaceAll(ret, "吧", "")
	ret = strings.ReplaceAll(ret, "你", "我")
	ret = strings.ReplaceAll(ret, "？", "!")
	ret = strings.ReplaceAll(ret, "?", "!")
	//ret, err := openai.AiMagic(ret)
	//if err != nil {
	//	return "err: " + err.Error()
	//}
	return ret
}
