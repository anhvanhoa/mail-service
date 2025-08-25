package grpcservice

import (
	"context"
	"fmt"
	"net"

	"mail-service/bootstrap"
	"mail-service/domain/service/logger"

	proto_mail_history "github.com/anhvanhoa/sf-proto/gen/mail_history/v1"
	proto_mail_provider "github.com/anhvanhoa/sf-proto/gen/mail_provider/v1"
	proto_mail_status "github.com/anhvanhoa/sf-proto/gen/mail_status/v1"
	proto_mail_tmpl "github.com/anhvanhoa/sf-proto/gen/mail_tmpl/v1"
	proto_status_history "github.com/anhvanhoa/sf-proto/gen/status_history/v1"
	proto_type_mail "github.com/anhvanhoa/sf-proto/gen/type_mail/v1"

	"buf.build/go/protovalidate"
	protovalidate_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	server *grpc.Server
	port   string
	log    logger.Log
}

func NewGRPCServer(
	env *bootstrap.Env,
	log logger.Log,
	mailHistoryService proto_mail_history.MailHistoryServiceServer,
	mailProviderService proto_mail_provider.MailProviderServiceServer,
	mailTmplService proto_mail_tmpl.MailTmplServiceServer,
	mailStatusService proto_mail_status.MailStatusServiceServer,
	typeMailService proto_type_mail.TypeMailServiceServer,
	statusHistoryService proto_status_history.StatusHistoryServiceServer,
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
	proto_mail_status.RegisterMailStatusServiceServer(server, mailStatusService)
	proto_type_mail.RegisterTypeMailServiceServer(server, typeMailService)
	proto_status_history.RegisterStatusHistoryServiceServer(server, statusHistoryService)
	proto_mail_provider.RegisterMailProviderServiceServer(server, mailProviderService)
	proto_mail_history.RegisterMailHistoryServiceServer(server, mailHistoryService)

	if !env.IsProduction() {
		log.Info("Registering reflection enabled")
		reflection.Register(server)
	}

	return &GRPCServer{
		server: server,
		port:   env.PORT_GRPC,
		log:    log,
	}
}

func (s *GRPCServer) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		s.log.Error(fmt.Sprintf("failed to listen: %v", err))
		return err
	}

	s.log.Info(fmt.Sprintf("gRPC server starting on port %s", s.port))

	go func() {
		if err := s.server.Serve(lis); err != nil {
			s.log.Error(fmt.Sprintf("failed to serve: %v", err))
		}
	}()

	<-ctx.Done()

	s.log.Info("Shutting down gRPC server...")
	s.server.GracefulStop()

	return nil
}

func (s *GRPCServer) Stop() {
	s.server.GracefulStop()
}

func (s *GRPCServer) GetServer() *grpc.Server {
	return s.server
}
