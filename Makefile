test:
	go test ./...

test-third-party: test-extism-sdk-go test-gopacket test-kit test-protobuf

test-extism-sdk-go:
	cd thirdparty/extism-sdk-go && go test ./...

test-gopacket:
	cd thirdparty/gopacket && go test ./...

test-kit:
	cd thirdparty/kit && go test ./...

test-protobuf:
	cd thirdparty/protobuf && go test ./...

