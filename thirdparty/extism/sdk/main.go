package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/extism/go-sdk"
	"log"
	"strings"
)

//go:embed wasm/plugin.wasm
var wasm []byte

// loadWASM returns the WASM core loaded into an extism.Plugin
func loadWASM(ctx context.Context, functions []extism.HostFunction) (*extism.Plugin, error) {
	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmData{
				Data: wasm,
			},
		},
	}

	config := extism.PluginConfig{
		EnableWasi: true,
	}
	plugin, err := extism.NewPlugin(ctx, manifest, config, functions)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize plugin: %v\n", err)
	}

	return plugin, nil
}

func main() {
	plugin, err := loadWASM(context.Background(), []extism.HostFunction{})
	if err != nil {
		log.Fatal(err)
	}

	keys, err := getKeys(plugin)
	if err != nil {
		log.Fatal(err)
	}

	plaintext := "hello extism"
	for _, key := range keys {
		ciphertext, err := encrypt(plugin, key, plaintext)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[encrypt][key: %s] input: %q, output: %q\n", key, plaintext, ciphertext)

		plaintext, err = decrypt(plugin, key, ciphertext)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[decrypt][key: %s] input: %q, output: %q\n", key, ciphertext, plaintext)
	}
}

func encrypt(plugin *extism.Plugin, key, data string) (string, error) {
	_, ret, err := plugin.Call("encrypt", getInput(key, data))
	return string(ret), err
}

func decrypt(plugin *extism.Plugin, key, data string) (string, error) {
	_, ret, err := plugin.Call("decrypt", getInput(key, data))
	return string(ret), err
}

func getKeys(plugin *extism.Plugin) ([]string, error) {
	_, ret, err := plugin.Call("getKeys", nil)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(ret), ","), nil
}

func getInput(key, data string) []byte {
	return []byte(strings.Join([]string{key, data}, ","))
}
