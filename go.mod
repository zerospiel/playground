module github.com/zerospiel/playground

go 1.27

godebug (
	default=go1.26
	fips140=on
)

require (
	github.com/envoyproxy/protoc-gen-validate v1.3.3
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.29.0
	golang.org/x/net v0.55.0
	google.golang.org/genproto/googleapis/api v0.0.0-20260523011958-0a33c5d7ca68
	google.golang.org/grpc v1.81.1
	google.golang.org/grpc/examples v0.0.0-20260522210837-db35da8bc5e8
	google.golang.org/protobuf v1.36.11
)

require (
	golang.org/x/sys v0.45.0 // indirect
	golang.org/x/text v0.37.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260523011958-0a33c5d7ca68 // indirect
)
