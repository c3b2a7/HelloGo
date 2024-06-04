```bash
tinygo build -o ../sdk/wasm/plugin.wasm -scheduler=none -target=wasi -gc=custom -tags='custommalloc nottinygc_finalizer' main.go
```
