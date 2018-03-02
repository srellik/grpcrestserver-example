package main

import (
	"context"

	"grpcrestserver-example/students_app/service_impl/studentsservice"

	"github.com/srellik/grpcrestserver"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	sspb "grpcrestserver-example/students_app/proto/students_service"
)

func main() {
	opts := grpcrestserver.Options{
		GRPCServerAddr: "localhost:8889",
		RESTServerAddr: ":8888",
		GRPCServiceHandlers: []grpcrestserver.GRPCHandlerCallback{
			func(s *grpc.Server) {
				sspb.RegisterStudentsServiceServer(s, studentsservice.New())
			},
		},
		RESTEndpointHandlers: []grpcrestserver.RESTHandlerCallback{
			func(ctx context.Context, gw *runtime.ServeMux, grpcServerAddr string, dialOpts []grpc.DialOption) error {
				return sspb.RegisterStudentsServiceHandlerFromEndpoint(ctx, gw, grpcServerAddr, dialOpts)
			},
		},
	}
	grpcrestserver.RunWith(opts)
}
