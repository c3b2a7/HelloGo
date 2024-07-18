
## Build WASI target

```bash
mkdir -p ../extism-sdk-go/internal/wasm && tinygo build -no-debug -o ../extism-sdk-go/internal/wasm/plugin.wasm -target=wasi .
```

## Optimization

```bash
wasm-opt -O4 --fast-math -o ../extism-sdk-go/internal/wasm/plugin-opt.wasm ../extism-sdk-go/internal/wasm/plugin.wasm
```