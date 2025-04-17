test:
	go test ./...

test-third-party: test-extism-sdk-go test-gopacket test-kit test-protobuf

test-extism-sdk-go:
	cd thirdparty/extism-sdk-go && go test --count=1 ./...

test-gopacket:
	cd thirdparty/gopacket && go test --count=1 ./...

test-kit:
	cd thirdparty/kit && go test --count=1 ./...

test-protobuf:
	cd thirdparty/protobuf && go test --count=1 ./...

