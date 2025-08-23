package grpcservice

import (
	"context"
	"fmt"
	"net"

	"mail-service/bootstrap"
	"mail-service/domain/service/logger"
	proto_mail_tmpl "mail-service/proto/gen/mail_tmpl/v1"

	"buf.build/go/protovalidate"
	protovalidate_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	server *grpc.Server
	port   string
}

func NewGRPCServer(
	env *bootstrap.Env,
	log logger.Log,
	mailTmplService proto_mail_tmpl.MailTmplServiceServer,
) *GRPCServer {
	validator, err := protovalidate.New()
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			protovalidate_middleware.UnaryServerInterceptor(validator),
			LoggingInterceptor(log),
		),
	)

	proto_mail_tmpl.RegisterMailTmplServiceServer(server, mailTmplService)

	if !env.IsProduction() {
		reflection.Register(server)
	}

	return &GRPCServer{
		server: server,
		port:   env.PORT_GRPC,
	}
}

func (s *GRPCServer) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	fmt.Printf("gRPC server starting on port %s\n", s.port)

	go func() {
		if err := s.server.Serve(lis); err != nil {
			fmt.Printf("failed to serve: %v", err)
		}
	}()

	<-ctx.Done()

	fmt.Println("Shutting down gRPC server...")
	s.server.GracefulStop()

	return nil
}

func (s *GRPCServer) Stop() {
	s.server.GracefulStop()
}

func (s *GRPCServer) GetServer() *grpc.Server {
	return s.server
}
