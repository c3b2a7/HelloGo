//go:build !optimize_wasm

package wasm

import _ "embed"

//go:embed plugin.wasm
var Data []byte
