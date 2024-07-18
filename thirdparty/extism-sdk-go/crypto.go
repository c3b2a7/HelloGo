package crypto

import (
	"context"
	"github.com/c3b2a7/HelloGo/thirdparty/extism-sdk-go/internal"
)

type CryptoAPI interface {
	Encrypt(key, plaintext string) (string, error)
	Decrypt(key, ciphertext string) (string, error)
}

type cryptoAPI struct {
	internal.InnerSDK
}

func (c cryptoAPI) Encrypt(key, plaintext string) (s string, err error) {
	err = c.InnerSDK.Invoke(context.Background(), "encrypt", []any{map[string]interface{}{
		"key":  key,
		"data": plaintext,
	}}, &s)
	return

}

func (c cryptoAPI) Decrypt(key, ciphertext string) (s string, err error) {
	err = c.InnerSDK.Invoke(context.Background(), "decrypt", []any{map[string]interface{}{
		"key":  key,
		"data": ciphertext,
	}}, &s)
	return
}

func NewCryptoAPI(innerSDK internal.InnerSDK) CryptoAPI {
	return &cryptoAPI{innerSDK}
}
