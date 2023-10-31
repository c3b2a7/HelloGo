gen-proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative **/**/*.proto

test: test-linux-amd64 test-linux-arm64 test-freebsd-amd64 test-freebsd-arm64 test-macos-amd64 test-macos-arm64 test-win64 test-win32

test-linux-amd64:
	GOARCH=amd64 GOOS=linux go test ./...

test-linux-arm64:
	GOARCH=arm64 GOOS=linux go test ./...

test-freebsd-amd64:
	GOARCH=amd64 GOOS=freebsd go test ./...

test-freebsd-arm64:
	GOARCH=arm64 GOOS=freebsd go test ./...

test-macos-amd64:
	GOARCH=amd64 GOOS=darwin go test ./...

test-macos-arm64:
	GOARCH=arm64 GOOS=darwin go test ./...

test-win64:
	GOARCH=amd64 GOOS=windows go test ./...

test-win32:
	GOARCH=386 GOOS=windows go test ./...
