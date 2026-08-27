module github.com/zerospiel/playground

go 1.28

godebug (
	default=go1.27
	fips140=on
)

require (
	github.com/envoyproxy/protoc-gen-validate v1.3.3
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.30.0
	golang.org/x/net v0.58.0
	google.golang.org/genproto/googleapis/api v0.0.0-20260825221802-da73d73af1c5
	google.golang.org/grpc v1.83.2
	google.golang.org/grpc/examples v0.0.0-20260827041148-664e87d82d84
	google.golang.org/protobuf v1.36.12
)

require (
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.41.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260825221802-da73d73af1c5 // indirect
)
