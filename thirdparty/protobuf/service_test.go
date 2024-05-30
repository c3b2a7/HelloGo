package protobuf

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

func TestGrpcServer(t *testing.T) {
	listen, _ := net.Listen("tcp", "127.0.0.1:9000")
	server := grpc.NewServer()

	RegisterGreetServiceServer(server, NewGreetServiceServer())

	go func() {
		err := server.Serve(listen)
		if err != nil {
			fmt.Printf("failed to serve: %v", err)
			return
		}
	}()
}

func TestUnaryCall(t *testing.T) {
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	clientConn, err := grpc.DialContext(ctx, "127.0.0.1:9000", opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer clientConn.Close()

	request := &Request{Id: 1, Type: Type_NORMAL, Data: []byte("在吗？")}
	log.Printf("send request %s\n", request)

	client := NewGreetServiceClient(clientConn)
	ctx, cancelFunc = context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	response, err := client.Hello(ctx, request)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("recv response: %s", response.Data)
}

func TestStreamCall(t *testing.T) {
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	clientConn, err := grpc.DialContext(ctx, "127.0.0.1:9000", opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer clientConn.Close()

	client := NewGreetServiceClient(clientConn)
	stream, err := client.HelloStream(context.Background())
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	done := make(chan struct{})
	errCh := make(chan error)
	go func() {
		for {
			response, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					close(done)
				} else {
					errCh <- err
				}
				return
			}
			log.Printf("recv response: %s\n", response)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	var id int64 = 0
	for {
		cmd, _ := reader.ReadString('\n') // 读到换行
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}
		if strings.ToLower(cmd) == "quit" {
			break
		}

		request := &Request{Id: id, Type: Type_NORMAL, Data: []byte(cmd)}
		log.Printf("send request %s\n", request)

		err := stream.Send(request)
		if err != nil {
			errCh <- err
			return
		}
		id++
	}
	stream.CloseSend()

	// 读取响应
	for {
		select {
		case <-done:
			return
		case err := <-errCh:
			log.Printf("err: %s", err)
		}
	}
}
