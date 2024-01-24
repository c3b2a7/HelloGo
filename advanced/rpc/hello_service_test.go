package rpc

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strings"
	"testing"
)

func StartGoRpcServer(address string, callback func(*rpc.Server, net.Conn)) error {
	server := rpc.NewServer()
	if err := RegisterHelloService(server, &HelloService{}); err != nil {
		return err
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	for {
		var conn net.Conn
		if conn, err = listener.Accept(); err != nil {
			log.Print("rpc.Serve: accept:", err.Error())
			return err
		}
		go callback(server, conn)
	}
}

func TestGobRpcServer(t *testing.T) {
	err := StartGoRpcServer(":8000", func(server *rpc.Server, conn net.Conn) {
		server.ServeConn(conn)
	})
	if err != nil {
		t.Error(err)
	}
}

func TestJsonRpcServer(t *testing.T) {
	err := StartGoRpcServer(":8000", func(server *rpc.Server, conn net.Conn) {
		server.ServeCodec(jsonrpc.NewServerCodec(conn)) // jsonrpc
	})
	if err != nil {
		t.Error(err)
	}
}

func TestHttpJsonRpcServer(t *testing.T) {
	rpc.RegisterName(HelloServiceName, &HelloService{})
	mux := http.NewServeMux()
	mux.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			w,
			r.Body,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":8000", mux)
}

func TestRawHttpJsonRpcServer(t *testing.T) {
	StartGoRpcServer(":8000", func(server *rpc.Server, conn net.Conn) {
		defer conn.Close()
		for {
			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil || request.Method != "POST" {
				resp := &http.Response{
					Request:    request,
					Proto:      request.Proto,
					ProtoMajor: request.ProtoMajor,
					ProtoMinor: request.ProtoMinor,
					StatusCode: http.StatusBadRequest,
				}
				resp.Write(conn)
				return
			}
			if request.URL.RequestURI() != "/jsonrpc" {
				resp := &http.Response{
					Request:    request,
					Proto:      request.Proto,
					ProtoMajor: request.ProtoMajor,
					ProtoMinor: request.ProtoMinor,
					StatusCode: http.StatusNotFound,
				}
				resp.Write(conn)
				continue
			}
			server.ServeRequest(jsonrpc.NewServerCodec(conn))
		}
	})
}

func TestGobRpcRequest(t *testing.T) {
	conn, err := net.Dial("tcp", ":8000")
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
	conn, err := net.Dial("tcp", ":8000")
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
	body := strings.NewReader(`{"method":"path/to/pkg.HelloService.Hello","params":["from jsonrpc client"],"id":0}`)
	req, _ := http.NewRequest("POST", "http://127.0.0.1:8000/jsonrpc", body)
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
