package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/c3b2a7/HelloGo/thirdparty/protobuf"
	"github.com/c3b2a7/HelloGo/thirdparty/protobuf/interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func StartGrpcServer() {
	listen, _ := net.Listen("tcp", "127.0.0.1:9000")

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.UnaryLoggingInterceptor),
		grpc.StreamInterceptor(interceptor.StreamLoggingInterceptor),
	)
	srv.RegisterService(&protobuf.GreetService_ServiceDesc, protobuf.NewGreetServiceServer())
	//protobuf.RegisterGreetServiceServer(srv, protobuf.NewGreetServiceServer())

	if err := srv.Serve(listen); err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}

func main() {
	go StartGrpcServer()

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

	client := protobuf.NewGreetServiceClient(clientConn)
	stream, err := client.HelloStream(context.Background())
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	defer stream.CloseSend()

	done := make(chan struct{})
	errCh := make(chan error, 1)
	go func() {
		for {
			response := new(protobuf.Response)
			if err = stream.RecvMsg(response); err != nil {
				if err == io.EOF {
					close(done)
				} else {
					errCh <- err
				}
				return
			}
			log.Printf("[Client] Recv response: %s\n", response)
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
			return
		}

		select {
		case <-done:
			return
		case err = <-errCh:
			log.Printf("err: %s", err)
			return
		default:
			request := &protobuf.Request{Id: id, Type: protobuf.Type_NORMAL, Data: []byte(cmd)}
			log.Printf("[Client] Send request %s\n", request)
			if err = stream.SendMsg(request); err != nil {
				if !errors.Is(err, io.EOF) {
					log.Printf("send msg err: %s", err)
				}
				return
			}
			id++
		}
	}
}
