package rpc

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strings"
	"testing"
)

func TestGobRpcServer(t *testing.T) {
	closer, err := StartGobRpcServer(":8000")
	require.NoError(t, err, "StartGobRpcServer should not return an error")
	closer()
}

func TestJsonRpcServer(t *testing.T) {
	closer, err := StartJsonRpcServer(":8001")
	require.NoError(t, err, "StartJsonRpcServer should not return an error")
	closer()
}

func TestHttpJsonRpcServer(t *testing.T) {
	closer, err := StartHttpJsonRpcServer(":8002")
	require.NoError(t, err, "StartHttpJsonRpcServer should not return an error")
	closer()
}

func TestRawHttpJsonRpcServer(t *testing.T) {
	closer, err := StartRawHttpJsonRpcServer(":8003")
	require.NoError(t, err, "StartRawHttpJsonRpcServer should not return an error")
	closer()
}

func TestGobRpcRequest(t *testing.T) {
	addr := "127.0.0.1:8004"
	closer, err := StartGobRpcServer(addr)
	if err != nil {
		t.Fatal(err)
	}
	defer closer()

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Error(err)
	}
	cc := rpc.NewClient(conn)

	helloService := NewHelloServiceClient(cc)
	var reply string
	if err = helloService.Hello("from gob client", &reply); err != nil {
		t.Error(err)
	}
	fmt.Println(reply)
}

func TestJsonRpcRequest(t *testing.T) {
	addr := "127.0.0.1:8005"
	closer, err := StartJsonRpcServer(addr)
	if err != nil {
		t.Fatal(err)
	}
	defer closer()

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Error(err)
	}
	cc := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	helloService := NewHelloServiceClient(cc)
	var reply string
	if err = helloService.Hello("from jsonrpc client", &reply); err != nil {
		t.Error(err)
	}
	fmt.Println(reply)
}

func TestHttpJsonRpcRequest(t *testing.T) {
	addr := "127.0.0.1:8006"
	closer, err := StartHttpJsonRpcServer(addr)
	if err != nil {
		t.Fatal(err)
	}
	defer closer()

	body := strings.NewReader(`{"method":"path/to/pkg.HelloService.Hello","params":["from jsonrpc client"],"id":0}`)
	req, _ := http.NewRequest("POST", "http://"+addr+"/jsonrpc", body)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	for i := 0; i < 5; i++ {
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return
		}
		resp.Body.Close()
		body.Seek(0, io.SeekStart)
		fmt.Println(resp.Request.RemoteAddr, string(b))
	}
}

func TestRawHttpJsonRpcRequest(t *testing.T) {
	addr := "127.0.0.1:8007"
	closer, err := StartRawHttpJsonRpcServer(addr)
	if err != nil {
		t.Fatal(err)
	}
	defer closer()

	body := strings.NewReader(`{"method":"path/to/pkg.HelloService.Hello","params":["from jsonrpc client"],"id":0}`)
	req, _ := http.NewRequest("POST", "http://"+addr+"/jsonrpc", body)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	for i := 0; i < 5; i++ {
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return
		}
		resp.Body.Close()
		body.Seek(0, io.SeekStart)
		fmt.Println(resp.Request.RemoteAddr, string(b))
	}
}
