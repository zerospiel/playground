module github.com/zerospiel/playground

go 1.28

godebug (
	default=go1.27
	fips140=on
)

require (
	github.com/envoyproxy/protoc-gen-validate v1.3.3
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.30.0
	golang.org/x/net v0.57.0
	google.golang.org/genproto/googleapis/api v0.0.0-20260803160001-6ac0973c030d
	google.golang.org/grpc v1.83.0
	google.golang.org/grpc/examples v0.0.0-20260806073429-03255a9237b6
	google.golang.org/protobuf v1.36.11
)

require (
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.40.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260803160001-6ac0973c030d // indirect
)
