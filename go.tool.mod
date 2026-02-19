module github.com/zerospiel/playground

go 1.27

tool google.golang.org/grpc/cmd/protoc-gen-go-grpc

require (
	github.com/envoyproxy/protoc-gen-validate v1.3.3
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.28.0
	golang.org/x/net v0.50.0
	google.golang.org/genproto/googleapis/api v0.0.0-20260217215200-42d3e9bedb6d
	google.golang.org/grpc v1.79.1
	google.golang.org/grpc/examples v0.0.0-20260218063904-c1a9239e408d
	google.golang.org/protobuf v1.36.11
)

require (
	golang.org/x/sys v0.41.0 // indirect
	golang.org/x/text v0.34.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260217215200-42d3e9bedb6d // indirect
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.5.1 // indirect
)
