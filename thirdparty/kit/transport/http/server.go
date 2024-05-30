package http

import (
	"context"
	"github.com/c3b2a7/HelloGo/thirdparty/kit/endpoint"
	"net/http"
)

type RequestDecoder func(ctx context.Context, request *http.Request) interface{}

type ResponseEncoder func(ctx context.Context, rw http.ResponseWriter, response interface{}) error

type severHandler struct {
	endpoint        endpoint.Endpoint
	requestDecoder  RequestDecoder
	responseEncoder ResponseEncoder
}

func (s severHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	response, err := s.endpoint(context.Background(), s.requestDecoder(context.Background(), request))
	if err != nil {
		return
	}
	err = s.responseEncoder(context.Background(), writer, response)
	if err != nil {
		return
	}

}

func NewServer(endpoint endpoint.Endpoint, decoder RequestDecoder, encoder ResponseEncoder) http.Handler {
	return &severHandler{endpoint, decoder, encoder}
}
