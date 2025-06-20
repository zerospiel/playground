module github.com/zerospiel/playground

go 1.25

godebug (
	default=go1.24
	fips140=on
)

require (
	github.com/envoyproxy/protoc-gen-validate v1.2.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.0
	golang.org/x/net v0.41.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250603155806-513f23925822
	google.golang.org/grpc v1.73.0
	google.golang.org/grpc/examples v0.0.0-20250619055035-0100d21c8f9b
	google.golang.org/protobuf v1.36.6
)

require (
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250603155806-513f23925822 // indirect
)
