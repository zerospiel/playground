module github.com/zerospiel/playground

go 1.25

godebug (
	default=go1.24
	fips140=on
)

require (
	github.com/envoyproxy/protoc-gen-validate v1.2.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.1
	golang.org/x/net v0.35.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250207221924-e9438ea467c6
	google.golang.org/grpc v1.70.0
	google.golang.org/grpc/examples v0.0.0-20250211194034-0003b4fa356e
	google.golang.org/protobuf v1.36.5
)

require (
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250207221924-e9438ea467c6 // indirect
)
