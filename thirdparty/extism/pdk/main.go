package main

import (
	"errors"
	"github.com/c3b2a7/HelloGo/thirdparty/extism-pdk-go/util"
	"github.com/extism/go-pdk"
	"strings"
)

var cryptoMap = map[string]util.Crypto{
	"bidev": util.NewAesCrypto("CAMG8wU1SwV9PoAJUsPsYV8TdiN+5LDxK7v2qwkATY8="),
	"biprd": util.NewAesCrypto("f2h3tGKAUzzbOEBZFt5n2bihFfOs4KwfX41As7tawTE="),
	"bdb":   util.NewPrefixCrypto("P001_2_", util.NewAesCrypto("V+MZQ0PyISJyL0Webmow5WB58yR6L6SmhbCl3KcGVuY=")),
	"db":    util.NewPrefixCrypto("P001_1_", util.NewAesCrypto("UnpKaEkgieup86ixZkJStQ==")),
}

//export getKeys
func getKeys() int32 {
	var keys []string
	for key := range cryptoMap {
		keys = append(keys, key)
	}
	pdk.OutputString(strings.Join(keys, ","))
	return 0
}

//export decrypt
func decrypt() int32 {
	input := string(pdk.Input())
	kp := strings.SplitN(input, ",", 2)
	if len(kp) != 2 {
		pdk.SetError(errors.New("invalid input"))
		return 1
	}
	key, ciphertext := kp[0], kp[1]
	s, err := cryptoMap[key].Decrypt(ciphertext)
	if err != nil {
		pdk.SetError(err)
		return 1
	}
	pdk.OutputString(s)
	return 0
}

//export encrypt
func encrypt() int32 {
	input := string(pdk.Input())
	kp := strings.SplitN(input, ",", 2)
	if len(kp) != 2 {
		pdk.SetError(errors.New("invalid input"))
		return 1
	}
	key, plaintext := kp[0], kp[1]
	s, err := cryptoMap[key].Encrypt(plaintext)
	if err != nil {
		pdk.SetError(err)
		return 1
	}
	pdk.OutputString(s)
	return 0
}

func main() {
}
