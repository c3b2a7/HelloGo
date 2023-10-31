package kit

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
	"net/url"
	"testing"
)

type request struct {
}

type response struct {
}

func TestKit(t *testing.T) {
	server := httptransport.NewServer(func(ctx context.Context, request interface{}) (interface{}, error) {
		return map[string]string{"a": "a", "b": "b"}, nil
	}, func(ctx context.Context, r *http.Request) (interface{}, error) {
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return nil, err
		}
		return req, nil
	}, httptransport.EncodeJSONResponse)

	u, _ := url.Parse("http://localhost:8080/kit")
	client := httptransport.NewClient("GET", u, httptransport.EncodeJSONRequest, func(ctx context.Context, h *http.Response) (response interface{}, err error) {
		var req request
		if err := json.NewDecoder(h.Body).Decode(&req); err != nil {
			return nil, err
		}
		return req, nil
	})
	client.Endpoint()

	mux := http.NewServeMux()
	mux.Handle("/kit", server)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", mux))
	}()
}
