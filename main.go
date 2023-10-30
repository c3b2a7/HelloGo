package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/c3b2a7/HelloGo/pkg/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func isPrime(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	if number > 1 {
		return true
	}
	return false
}

func findPrimeGame() {
	fmt.Println("Prime numbers less than 20:")
	for number := 0; number < 20; number++ {
		if isPrime(number) {
			fmt.Printf("%d ", number)
		}
	}
}

func guessNumberGame() {
	val := 0
	for {
		fmt.Print("Enter number:")
		fmt.Scanf("%d", &val)
		switch {
		case val < 0:
			panic("You entered a negative number!")
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		default:
			fmt.Println("You entered:", val)
		}
	}
}

func main() {
	//findPrimeGame()
	//guessNumberGame()
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
			log.Printf("Recv response: %s\n", response)
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

		request := &protobuf.Request{Id: id, Type: protobuf.Type_NORMAL, Data: []byte(cmd)}
		log.Printf("Send request %s\n", request)

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
