package crypto

import (
	"context"
	"testing"

	"github.com/c3b2a7/HelloGo/thirdparty/extism-sdk-go/internal"
)

var testcases = []struct {
	key        string
	plaintext  string
	ciphertext string
}{
	{"p001", "hello extism sdk", "P001_FiV7qf2MqgS6WvDeJwQcLvX2akrs0H3KW2tV4+rJfpo="},
	{"p002", "hello extism sdk", "P002_MAdZBjV0uhxpj8taAipv3da9t6T2wr6qUO5yes1Ld2k="},
	{"t1", "hello extism sdk", "nCBwzwK1BM9DQRMXRJ1wYFy+5uevcoPSvL/RIOIghOQ="},
	{"t2", "hello extism sdk", "groNdGUtkWQsPfZx/MezqSyPCr7OpODmd6hzITHTWUU="},
}

func TestNewClient(t *testing.T) {
	client := Must(NewClient(context.Background(), WithOSVersion("14.5 (23F79)")))
	if client == nil {
		t.Error("client is nil")
	}
}

func TestNewCryptoAPI(t *testing.T) {
	api := NewCryptoAPI(internal.InnerSDK{Core: internal.GetSharedCore()})
	if api == nil {
		t.Error("CryptoAPI is nil")
	}
}

func TestNewKeysAPI(t *testing.T) {
	api := NewKeysAPI(internal.InnerSDK{Core: internal.GetSharedCore()})
	if api == nil {
		t.Error("KeysAPI is nil")
	}
}

func TestCryptoAPI_Decrypt(t *testing.T) {
	client := Must(NewClient(context.Background()))
	for _, testcase := range testcases {
		t.Run(testcase.key, func(t *testing.T) {
			got, err := client.CryptoAPI.Decrypt(testcase.key, testcase.ciphertext)
			if err != nil {
				t.Fatal(err)
			}
			if got != testcase.plaintext {
				t.Errorf("Decrypt() got = %v, want %v", got, testcase.plaintext)
			}
		})
	}
}

func TestCryptoAPI_Encrypt(t *testing.T) {
	client := Must(NewClient(context.Background()))
	for _, testcase := range testcases {
		t.Run(testcase.key, func(t *testing.T) {
			got, err := client.CryptoAPI.Encrypt(testcase.key, testcase.plaintext)
			if err != nil {
				t.Fatal(err)
			}
			if got != testcase.ciphertext {
				t.Errorf("Encrypt() got = %v, want %v", got, testcase.ciphertext)
			}
		})
	}
}

func TestCryptoAPI_EncryptErr(t *testing.T) {
	client := Must(NewClient(context.Background()))
	_, err := client.CryptoAPI.Encrypt("unknown key", "hello extism sdk")
	if err == nil {
		t.Fatal("expected an error")
	}
}

func TestKeysAPI_GetKeys(t *testing.T) {
	client := Must(NewClient(context.Background()))
	keys, err := client.KeysAPI.GetKeys(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	var wantKeys []string
	for _, testcase := range testcases {
		wantKeys = append(wantKeys, testcase.key)
	}

	if !slicesEqualUnordered(keys, wantKeys) {
		t.Errorf("GetKeys() got = %v, want %v", keys, wantKeys)
	}
}

func slicesEqualUnordered(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	countMap := make(map[string]int)
	for _, val := range a {
		countMap[val]++
	}
	for _, val := range b {
		countMap[val]--
		if countMap[val] < 0 {
			return false
		}
	}
	for _, count := range countMap {
		if count != 0 {
			return false
		}
	}
	return true
}

func BenchmarkCryptoAPI_Encrypt(b *testing.B) {
	client := Must(NewClient(context.Background()))
	for _, testcase := range testcases {
		b.Run(testcase.key, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ciphertext, err := client.CryptoAPI.Encrypt(testcase.key, testcase.plaintext)
				if err != nil {
					b.Error(err)
				}
				if ciphertext != testcase.ciphertext {
					b.Errorf("Encrypt() got = %v, want %v", ciphertext, testcase.ciphertext)
				}
			}
		})
	}
}

func BenchmarkCryptoAPI_Decrypt(b *testing.B) {
	client := Must(NewClient(context.Background()))
	for _, testcase := range testcases {
		b.Run(testcase.key, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				plaintext, err := client.CryptoAPI.Decrypt(testcase.key, testcase.ciphertext)
				if err != nil {
					b.Error(err)
				}
				if plaintext != testcase.plaintext {
					b.Errorf("Encrypt() got = %v, want %v", plaintext, testcase.plaintext)
				}
			}
		})
	}
}

func BenchmarkKeysAPI_GetKeys(b *testing.B) {
	client := Must(NewClient(context.Background()))
	for i := 0; i < b.N; i++ {
		_, err := client.KeysAPI.GetKeys(context.Background())
		if err != nil {
			b.Error(err)
		}
	}
}
