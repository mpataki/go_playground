module github.com/mpataki/go_playground/service

go 1.23.4

require (
	github.com/bufbuild/connect-go v1.10.0
	github.com/mpataki/go_playground/proto v0.0.0
)

require (
	golang.org/x/net v0.34.0 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
)

replace github.com/mpataki/go_playground/proto => ../proto
