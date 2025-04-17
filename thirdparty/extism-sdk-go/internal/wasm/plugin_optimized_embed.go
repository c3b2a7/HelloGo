//go:build optimize_wasm

package wasm

import (
	_ "embed"
	"fmt"
)

//go:embed plugin-opt.wasm
var Data []byte

func init() {
	fmt.Println("wasm: embed optimized plugin")
}
