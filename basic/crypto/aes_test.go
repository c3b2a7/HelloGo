package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"testing"
)

func TestCryptoDispatcher(t *testing.T) {
	// AES256 32byte 32*8=256bit
	// AES192 24byte 24*8=196bit
	// AES128 16byte 16*8=128bit

	// AES-256 使用 32 字节的密钥
	key := make([]byte, 32)

	// 生成随机密钥
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err)
	}

	fmt.Printf("Generated AES-256 Key: %s\n", base64.StdEncoding.EncodeToString(key))

	// 验证密钥长度
	if len(key) != 32 {
		fmt.Println("Key length is incorrect")
		return
	}
}
