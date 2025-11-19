package servers

import (
    "context"
    "fmt"
    "time"
    "connectrpc.com/connect"
    greetingv1 "github.com/mpataki/go_playground/proto/gen/go/mpataki/go_playground/proto/greeting/v1"
    "github.com/mpataki/go_playground/proto/gen/go/mpataki/go_playground/proto/greeting/v1/greetingv1connect"
)

type GreetingServer struct {
	greetingv1connect.UnimplementedGreetingServiceHandler
}

// NewGreetingServer creates a new GreetingServer instance
func NewGreetingServer() *GreetingServer {
	return &GreetingServer{}
}

func (s *GreetingServer) SayHello(ctx context.Context, req *connect.Request[greetingv1.HelloRequest]) (*connect.Response[greetingv1.HelloResponse], error) {
    name := req.Msg.GetName()
    if name == "" {
        name = "World"
    }
    
    response := &greetingv1.HelloResponse{
        Message: fmt.Sprintf("Hello, %s!", name),
        Timestamp: time.Now().Unix(),
    }
    
    return connect.NewResponse(response), nil
}
