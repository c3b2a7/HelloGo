package main

import (
	"encoding/json"
	"fmt"
	"github.com/c3b2a7/HelloGo/thirdparty/extism-pdk-go/registry"

	"github.com/extism/go-pdk"
)

type Code uint32

const (
	SuccessCode Code = iota
	ErrorCode
)

// Invocation holds the information required for invoking SDK functionality.
type Invocation struct {
	Method string            `json:"method"`
	Args   []json.RawMessage `json:"args"`
}

//go:wasmimport host getKey
func hostGetKey(offset uint64) uint64

//go:wasmimport host getData
func hostGetData(offset uint64) uint64

func getKey(args []json.RawMessage) string {
	mem := pdk.AllocateBytes(args[0])
	defer mem.Free()
	ptr := hostGetKey(mem.Offset())
	rmem := pdk.FindMemory(ptr)
	defer rmem.Free()
	return string(rmem.ReadBytes())
}

func getData(args []json.RawMessage) string {
	mem := pdk.AllocateBytes(args[0])
	defer mem.Free()
	ptr := hostGetData(mem.Offset())
	rmem := pdk.FindMemory(ptr)
	defer rmem.Free()
	return string(rmem.ReadBytes())
}

// WASI supports two modules: Reactors and Commands
// `_initialize` is a startup callback for reactors module in WASI application
// see: https://github.com/WebAssembly/WASI/blob/main/legacy/application-abi.md
//
// we register real methods that can be invoked here
//
//export _initialize
func initialize() {
	handleCryptoRequest := func(args []json.RawMessage, crypto func(key, data string) (string, error)) (any, error) {
		return crypto(getKey(args), getData(args))
	}

	registry.RegisterMethod("encrypt", func(args []json.RawMessage) (any, error) {
		return handleCryptoRequest(args, dispatcher.Encrypt)
	})
	registry.RegisterMethod("decrypt", func(args []json.RawMessage) (any, error) {
		return handleCryptoRequest(args, dispatcher.Decrypt)
	})
	registry.RegisterMethod("getKeys", func(_ []json.RawMessage) (any, error) {
		return dispatcher.GetKeys(), nil
	})
}

//export invoke
func invoke() Code {
	var invocation Invocation
	if err := pdk.InputJSON(&invocation); err != nil {
		return handleError(err)
	}

	method := registry.GetMethod(invocation.Method)
	if method == nil {
		return handleError(fmt.Errorf("method '%s' not registered", invocation.Method))
	}

	return must(method.Call(invocation.Args))
}

func must(v any, err error) Code {
	if err != nil {
		pdk.SetError(err)
		return ErrorCode
	}

	return handleError(pdk.OutputJSON(&v))
}

func handleError(err error) Code {
	if err != nil {
		pdk.SetError(err)
		return ErrorCode
	}
	return SuccessCode
}

func main() {
}
