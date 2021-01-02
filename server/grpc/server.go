package grpc

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"github.com/go-kratos/kratos/v2/server"
)

var _ server.Server = new(Server)

// Server is a gRPC server wrapper.
type Server struct {
	*grpc.Server

	network string
	addr    string
}

// NewServer creates a gRPC server by options.
func NewServer(network, addr string, opts ...grpc.ServerOption) *Server {
	return &Server{
		network: network,
		addr:    addr,
		Server:  grpc.NewServer(opts...),
	}
}

// Start start the gRPC server.
func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen(s.network, s.addr)
	if err != nil {
		return err
	}
	return s.Serve(lis)
}

// Stop stop the gRPC server.
func (s *Server) Stop(ctx context.Context) error {
	s.GracefulStop()
	return nil
}
