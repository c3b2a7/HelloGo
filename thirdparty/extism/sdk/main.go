package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/extism/go-sdk"
	"log"
	"os"
)

//go:embed wasm/plugin.wasm
var wasm []byte

func main() {
	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmData{
				Data: wasm,
			},
		},
	}

	ctx := context.Background()
	config := extism.PluginConfig{
		EnableWasi: true,
	}
	plugin, err := extism.NewPlugin(ctx, manifest, config, []extism.HostFunction{})

	if err != nil {
		fmt.Printf("Failed to initialize plugin: %v\n", err)
		os.Exit(1)
	}
	_, i, err := plugin.CallWithContext(context.Background(), "greet", []byte("Jack"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response: %s", string(i))

	_, i, err = plugin.CallWithContext(context.Background(), "greet", []byte("Benjamin"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response: %s", string(i))
}
