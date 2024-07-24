package internal

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/c3b2a7/HelloGo/thirdparty/extism-sdk-go/internal/wasm"

	extism "github.com/extism/go-sdk"
)

const (
	invokeFuncName = "invoke"
	initFuncName   = "init"
)

var core *SharedCore
var mutex sync.Mutex

// GetSharedCore initializes the shared core once and returns the already existing one on subsequent calls.
func GetSharedCore() *SharedCore {
	if core == nil {
		mutex.Lock()
		if core == nil {
			plugin, err := loadWASM(context.Background(), getHostFunctions())
			if err != nil {
				panic(err)
			}
			core = &SharedCore{plugin: plugin}
		}
		mutex.Unlock()
	}
	return core
}

func ReleaseCore() {
	core.plugin.Close()
	core = nil
}

// SharedCore implements Core in such a way that all created client instances share the same core resources.
type SharedCore struct {
	// lock is used to synchronize access to the shared WASM core which is single threaded
	lock sync.Mutex
	// plugin is the Extism plugin which represents the WASM core loaded into memory
	plugin *extism.Plugin
}

func (c *SharedCore) Init(ctx context.Context) error {
	if !c.plugin.FunctionExists(initFuncName) {
		return nil
	}
	_, err := c.callWithCtx(ctx, initFuncName, nil)
	if err != nil {
		return err
	}
	return nil
}

// Invoke calls specified logic from core
func (c *SharedCore) Invoke(ctx context.Context, invocation Invocation) ([]byte, error) {
	input, err := json.Marshal(invocation)
	if err != nil {
		return nil, err
	}
	res, err := c.callWithCtx(ctx, invokeFuncName, input)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *SharedCore) callWithCtx(ctx context.Context, functionName string, serializedParameters []byte) ([]byte, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	_, response, err := c.plugin.CallWithContext(ctx, functionName, serializedParameters)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// loadWASM returns the WASM core loaded into an extism.Plugin
func loadWASM(ctx context.Context, functions []extism.HostFunction) (*extism.Plugin, error) {
	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmData{
				Data: wasm.Data,
			},
		},
	}

	config := extism.PluginConfig{
		EnableWasi: true,
	}
	if _, ok := os.LookupEnv("EXTISM_PLUGIN_DEBUG"); ok {
		config.LogLevel = extism.LogLevelTrace
	}
	plugin, err := extism.NewPlugin(ctx, manifest, config, functions)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize plugin: %v", err)
	}

	return plugin, nil
}
