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
	return g.processReq(ctx, request)
}

func (g *greetServiceServerImpl) HelloStream(stream GreetService_HelloStreamServer) error {
	reqCh, errCh := recvReq(stream)

	for {
		select {
		case err, ok := <-errCh:
			if !ok || err == nil {
				return nil // err channel closed or err is nil
			}
			return g.processErr(err)
		case req, ok := <-reqCh:
			if !ok {
				return nil // req channel closed
			}
			log.Printf("[Server] Recv request %v\n", req)
			resp, err := g.processReq(stream.Context(), req)
			if err != nil {
				return err
			}
			log.Printf("[Server] Send Response %v", resp)
			if err = stream.Send(resp); err != nil {
				return g.processErr(err)
			}
		}
	}
}

func (g *greetServiceServerImpl) processErr(err error) error {
	if err == io.EOF { // errors.Is(err, io.EOF)
		return nil
	}
	return err
}

func (g *greetServiceServerImpl) processReq(_ context.Context, req *Request) (*Response, error) {
	resp := Response{}
	resp.Id = req.Id
	switch req.Type {
	case Type_PING:
		resp.Type = Type_PONG
		resp.Data = []byte("pong!")
	case Type_PONG:
		resp.Type = Type_NORMAL
		if ret, err := magic(nil); err != nil {
			return nil, fmt.Errorf("processReq req[%v] method: magic, err: %w", req, err)
		} else {
			resp.Data = []byte(ret)
		}
	case Type_NORMAL:
		resp.Type = Type_NORMAL
		if ret, err := magic(req.Data); err != nil {
			return nil, fmt.Errorf("processReq req[%v] method: magic, err: %w", req, err)
		} else {
			resp.Data = []byte(ret)
		}
	}
	return &resp, nil
}

func recvReq(stream GreetService_HelloStreamServer) (<-chan *Request, <-chan error) {
	// 封装读取请求和错误处理的逻辑
	// 调用方只需处理返回的 chan
	reqCh := make(chan *Request)
	errCh := make(chan error, 1)

	go func() {
		defer close(reqCh)
		defer close(errCh)

		for {
			req, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					errCh <- err
				}
				return
			}

			select {
			case <-stream.Context().Done():
				errCh <- stream.Context().Err()
				return
			case reqCh <- req:
			}
		}
	}()

	return reqCh, errCh
}

func magic(b []byte) (ret string, err error) {
	if b == nil {
		return "Hello, Grpc!", nil
	}
	ret = string(b)
	ret = strings.ReplaceAll(ret, "吗", "")
	ret = strings.ReplaceAll(ret, "吧", "")
	ret = strings.ReplaceAll(ret, "你", "我")
	ret = strings.ReplaceAll(ret, "？", "!")
	ret = strings.ReplaceAll(ret, "?", "!")
	return ret, nil
}
