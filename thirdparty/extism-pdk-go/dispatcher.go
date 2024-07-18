package main

import (
	"errors"
	"github.com/c3b2a7/HelloGo/thirdparty/extism-pdk-go/util"
)

var dispatcher = CryptoDispatcher{
	"t1":   util.NewAesCrypto("HZ2WnzChtI6oEAQDNJ0JcqAxx6R2BW41kbD/4Yeond8="),
	"t2":   util.NewAesCrypto("6b/3GEwAThajVH3CoyIRwT9WQNB3kXJB7tWmSyEduiw="),
	"p001": util.NewPrefixCrypto("P001_", util.NewAesCrypto("ypHziBAsGKU9+vDyDYl3cA==")),
	"p002": util.NewPrefixCrypto("P002_", util.NewAesCrypto("JlZw22U2DNBc+gHVdlGONX7Z4TWzLfhR")),
}

type CryptoDispatcher map[string]util.Crypto

func (c CryptoDispatcher) Encrypt(key, plaintext string) (string, error) {
	if cc, ok := c[key]; ok {
		return cc.Encrypt(plaintext)
	}
	return "", errors.New("unknown crypto key: " + key)
}

func (c CryptoDispatcher) Decrypt(key, ciphertext string) (string, error) {
	if cc, ok := c[key]; ok {
		return cc.Decrypt(ciphertext)
	}
	return "", errors.New("unknown crypto key: " + key)
}

func (c CryptoDispatcher) GetKeys() []string {
	var keys []string
	for key := range c {
		keys = append(keys, key)
	}
	return keys
}
