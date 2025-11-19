package main

import (
    "fmt"
    "net/http"
    "golang.org/x/net/http2"
    "golang.org/x/net/http2/h2c"
	"connectrpc.com/grpcreflect"

    "github.com/mpataki/go_playground/service/servers"
    "github.com/mpataki/go_playground/proto/gen/go/mpataki/go_playground/proto/greeting/v1/greetingv1connect"
)

func main() {
    mux := http.NewServeMux()

    // Initialize the greeting server
    greetingServer := servers.NewGreetingServer()

    // Register the greeting service
    path, handler := greetingv1connect.NewGreetingServiceHandler(greetingServer)
    mux.Handle(path, handler)

	reflector := grpcreflect.NewStaticReflector(
		greetingv1connect.GreetingServiceName,
	)

	mux.Handle(grpcreflect.NewHandlerV1(reflector))

    // Start the server with h2c support
    port := "8080"
    fmt.Printf("gRPC server listening on port %s...\n", port)
    if err := http.ListenAndServe(
        ":"+port,
        h2c.NewHandler(mux, &http2.Server{}),
    ); err != nil {
        fmt.Printf("Failed to start server: %v\n", err)
    }
}

