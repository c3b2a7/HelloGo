//go:build optimize_wasm

package wasm

import _ "embed"

//go:embed plugin-opt.wasm
var Data []byte
