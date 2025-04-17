package internal

import (
	"context"
	"encoding/json"

	extism "github.com/extism/go-sdk"
)

func getHostFunctions() []extism.HostFunction {
	return []extism.HostFunction{getKeyHostFunction(), getDataHostFunctions()}
}

func getKeyHostFunction() extism.HostFunction {
	getKey := extism.NewHostFunctionWithStack("getKey", func(ctx context.Context, p *extism.CurrentPlugin, stack []uint64) {
		bytes, err := p.ReadBytes(stack[0])
		if err != nil {
			panic(err)
		}
		var dataMap []any
		if err = json.Unmarshal(bytes, &dataMap); err != nil {
			panic(err)
		}
		key, ok := dataMap[0].(string) // key in arg[0]
		if !ok {
			key = ""
		}
		stack[0], _ = p.WriteString(key)
	},
		[]extism.ValueType{extism.ValueTypeI64},
		[]extism.ValueType{extism.ValueTypeI64},
	)
	getKey.SetNamespace("host")

	return getKey
}

func getDataHostFunctions() extism.HostFunction {
	getData := extism.NewHostFunctionWithStack("getData", func(ctx context.Context, p *extism.CurrentPlugin, stack []uint64) {
		bytes, err := p.ReadBytes(stack[0])
		if err != nil {
			panic(err)
		}
		var dataMap []any
		if err = json.Unmarshal(bytes, &dataMap); err != nil {
			panic(err)
		}
		data, ok := dataMap[1].(string) // data in arg[1]
		if !ok {
			data = ""
		}
		stack[0], _ = p.WriteString(data)
	},
		[]extism.ValueType{extism.ValueTypeI64},
		[]extism.ValueType{extism.ValueTypeI64},
	)
	getData.SetNamespace("host")

	return getData
}
