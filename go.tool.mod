module github.com/zerospiel/playground

go 1.26

tool google.golang.org/grpc/cmd/protoc-gen-go-grpc

require (
	github.com/envoyproxy/protoc-gen-validate v1.2.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.1
	golang.org/x/net v0.42.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250804133106-a7a43d27e69b
	google.golang.org/grpc v1.74.2
	google.golang.org/grpc/examples v0.0.0-20250818051530-9ac0ec87ca2e
	google.golang.org/protobuf v1.36.6
)

require (
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250804133106-a7a43d27e69b // indirect
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.5.1 // indirect
)
