package util

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type Crypto interface {
	Encrypt(plaintext string) (string, error)
	Decrypt(ciphertext string) (string, error)
}

func NewAesCrypto(key string) Crypto {
	aesKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		panic(err)
	}
	return &aesCrypto{aesKey}
}

func NewPrefixCrypto(prefix string, crypto Crypto) Crypto {
	return &prefixCrypto{prefix, crypto}
}

type aesCrypto struct {
	aesKey []byte
}

func (a aesCrypto) Encrypt(plaintext string) (string, error) {
	encrypt, err := AesECBEncrypt([]byte(plaintext), a.aesKey, Pkcs5Padding)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encrypt), nil
}

func (a aesCrypto) Decrypt(ciphertext string) (string, error) {
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	decrypt, err := AesECBDecrypt(ciphertextBytes, a.aesKey, Pkcs5Padding)
	if err != nil {
		return "", err
	}
	return string(decrypt), nil
}

type prefixCrypto struct {
	prefix string
	Crypto
}

func (p prefixCrypto) Encrypt(plaintext string) (string, error) {
	encrypt, err := p.Crypto.Encrypt(plaintext)
	if err != nil {
		return "", err
	}
	return p.prefix + encrypt, nil
}

func (p prefixCrypto) Decrypt(ciphertext string) (string, error) {
	if !strings.HasPrefix(ciphertext, p.prefix) {
		return "", fmt.Errorf("ciphertext does not contain a prefix: %s", p.prefix)
	}
	return p.Crypto.Decrypt(ciphertext[len(p.prefix):])
}
