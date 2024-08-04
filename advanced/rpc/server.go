package rpc

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"
	"time"
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
			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
				continue
			}

			if request.Method != "POST" {
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
			server.ServeRequest(newRawHttpCodec(conn, request))
		}
	})
}

type rawHttpCodec struct {
	rpc.ServerCodec

	request *http.Request
	conn    net.Conn

	statusCode int
	header     http.Header
	body       *bytes.Buffer
}

func newRawHttpCodec(conn net.Conn, request *http.Request) *rawHttpCodec {
	codec := &rawHttpCodec{
		request:    request,
		conn:       conn,
		statusCode: http.StatusOK,
		header:     make(http.Header),
		body:       new(bytes.Buffer),
	}
	codec.ServerCodec = jsonrpc.NewServerCodec(struct {
		io.Writer
		io.ReadCloser
	}{
		codec.body,
		request.Body,
	})
	return codec
}

// WriteResponse overrides to write response headers
func (rw *rawHttpCodec) WriteResponse(resp *rpc.Response, replay any) error {
	// write body
	err := rw.ServerCodec.WriteResponse(resp, replay)
	if err != nil {
		return err
	}

	// write headers
	if rw.header.Get("Date") == "" {
		rw.header.Set("Date", time.Now().UTC().Format(http.TimeFormat))
	}

	contentType := http.DetectContentType(rw.body.Bytes())
	rw.header.Set("Content-Type", contentType)
	rw.header.Set("Content-Length", strconv.Itoa(rw.body.Len()))

	return rw.Flush()
}

// WriteHeader sends an HTTP response header with the provided status code
func (rw *rawHttpCodec) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
}

// Flush sends the accumulated response to the client
func (rw *rawHttpCodec) Flush() error {
	writer := bufio.NewWriter(rw.conn)

	// Write status line
	statusLine := fmt.Sprintf("HTTP/%d.%d %d %s\r\n", rw.request.ProtoMajor, rw.request.ProtoMinor,
		rw.statusCode, http.StatusText(rw.statusCode))
	writer.WriteString(statusLine)

	// Write headers
	for key, values := range rw.header {
		for _, value := range values {
			writer.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
		}
	}
	writer.WriteString("\r\n")

	// Write body
	_, err := writer.Write(rw.body.Bytes())
	if err != nil {
		return err
	}

	return writer.Flush()
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
