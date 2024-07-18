package crypto

import (
	"context"
	"github.com/c3b2a7/HelloGo/thirdparty/extism-sdk-go/internal"
)

type KeysAPI interface {
	GetKeys(ctx context.Context) ([]string, error)
}

type keysAPI struct {
	internal.InnerSDK
}

func (k keysAPI) GetKeys(ctx context.Context) (ret []string, err error) {
	err = k.InnerSDK.Invoke(ctx, "getKeys", nil, &ret)
	return
}

func NewKeysAPI(innerSDK internal.InnerSDK) KeysAPI {
	return &keysAPI{innerSDK}
}
