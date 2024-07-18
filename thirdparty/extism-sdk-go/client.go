package crypto

import (
	"context"
	"github.com/c3b2a7/HelloGo/thirdparty/extism-sdk-go/internal"
	"runtime"
)

type Client struct {
	runtimeInfo internal.SDKRuntimeInfo
	CryptoAPI
	KeysAPI
}

type ClientOption func(client *Client)

func Must(client *Client, err error) *Client {
	if err != nil {
		panic(err)
	}
	return client
}

// NewClient returns a 1Password Go SDK client using the provided ClientOption list.
func NewClient(ctx context.Context, opts ...ClientOption) (*Client, error) {
	core := internal.GetSharedCore()
	return createClient(ctx, core, opts...)
}

func createClient(ctx context.Context, core internal.Core, opts ...ClientOption) (*Client, error) {
	client := Client{
		runtimeInfo: internal.NewDefaultSDKRuntimeInfo(),
	}

	for _, opt := range opts {
		opt(&client)
	}

	if err := core.Init(ctx); err != nil {
		return nil, err
	}

	innerSDK := internal.InnerSDK{Core: core}
	initAPIs(&client, innerSDK)
	runtime.SetFinalizer(&client, func(f *Client) {
		internal.ReleaseCore()
	})

	return &client, nil
}

func initAPIs(client *Client, innerSDK internal.InnerSDK) {
	client.CryptoAPI = NewCryptoAPI(innerSDK)
	client.KeysAPI = NewKeysAPI(innerSDK)
}

func WithOSVersion(osVersion string) ClientOption {
	return func(client *Client) {
		client.runtimeInfo.SystemOSVersion = osVersion
	}
}
