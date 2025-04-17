//go:build !optimize_wasm

package wasm

import (
	_ "embed"
	"fmt"
)

//go:embed plugin.wasm
var Data []byte

func init() {
	fmt.Println("wasm: embed plugin")
}
