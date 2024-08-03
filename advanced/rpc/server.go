package rpc

import (
	"bufio"
	"context"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type ServerCloser func()

func StartGobRpcServer(addr string) (ServerCloser, error) {
	return startRpcServer(addr, func(server *rpc.Server, conn net.Conn) {
		server.ServeConn(conn)
	})
}

func StartJsonRpcServer(addr string) (ServerCloser, error) {
	return startRpcServer(addr, func(server *rpc.Server, conn net.Conn) {
		server.ServeCodec(jsonrpc.NewServerCodec(conn)) // jsonrpc
	})
}

func StartHttpJsonRpcServer(addr string) (ServerCloser, error) {
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

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	go func() {
		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		select {
		case <-ctx.Done():
			server.Close()
			return
		}
	}()

	// FIXME: using startRpcServer
	//startRpcServer(addr, func(server *rpc.Server, conn net.Conn) {
	//})

	return ServerCloser(cancelFunc), nil
}

func StartRawHttpJsonRpcServer(addr string) (ServerCloser, error) {
	return startRpcServer(addr, func(server *rpc.Server, conn net.Conn) {
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

func startRpcServer(address string, callback func(*rpc.Server, net.Conn)) (ServerCloser, error) {
	server := rpc.NewServer()
	if err := RegisterHelloService(server, &HelloService{}); err != nil {
		return nil, err
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		defer func() { _ = listener.Close() }()
		for {

			select {
			case <-ctx.Done():
				return
			default:
			}

			var conn net.Conn
			if conn, err = listener.Accept(); err != nil {
				log.Print("rpc.Serve: accept:", err.Error())
				panic(err)
			}
			go callback(server, conn)
		}
	}()

	return ServerCloser(cancelFunc), nil
}
