package grpcservice

import (
	"mail-service/bootstrap"

	grpc_service "github.com/anhvanhoa/service-core/bootstrap/grpc"
	"github.com/anhvanhoa/service-core/domain/cache"
	"github.com/anhvanhoa/service-core/domain/log"
	"github.com/anhvanhoa/service-core/domain/token"
	"github.com/anhvanhoa/service-core/domain/user_context"
	proto_mail_history "github.com/anhvanhoa/sf-proto/gen/mail_history/v1"
	proto_mail_provider "github.com/anhvanhoa/sf-proto/gen/mail_provider/v1"
	proto_mail_status "github.com/anhvanhoa/sf-proto/gen/mail_status/v1"
	proto_mail_tmpl "github.com/anhvanhoa/sf-proto/gen/mail_tmpl/v1"
	proto_status_history "github.com/anhvanhoa/sf-proto/gen/status_history/v1"
	proto_type_mail "github.com/anhvanhoa/sf-proto/gen/type_mail/v1"

	"google.golang.org/grpc"
)

func NewGRPCServer(
	env *bootstrap.Env,
	log *log.LogGRPCImpl,
	cacher cache.CacheI,
	mailHistoryService proto_mail_history.MailHistoryServiceServer,
	mailProviderService proto_mail_provider.MailProviderServiceServer,
	mailTmplService proto_mail_tmpl.MailTmplServiceServer,
	mailStatusService proto_mail_status.MailStatusServiceServer,
	typeMailService proto_type_mail.TypeMailServiceServer,
	statusHistoryService proto_status_history.StatusHistoryServiceServer,
) *grpc_service.GRPCServer {
	config := &grpc_service.GRPCServerConfig{
		PortGRPC:     env.PortGrpc,
		IsProduction: env.IsProduction(),
		NameService:  env.NameService,
	}
	middleware := grpc_service.NewMiddleware(
		token.NewToken(env.AccessSecret),
	)
	return grpc_service.NewGRPCServer(
		config,
		log,
		func(server *grpc.Server) {
			proto_mail_tmpl.RegisterMailTmplServiceServer(server, mailTmplService)
			proto_mail_status.RegisterMailStatusServiceServer(server, mailStatusService)
			proto_type_mail.RegisterTypeMailServiceServer(server, typeMailService)
			proto_status_history.RegisterStatusHistoryServiceServer(server, statusHistoryService)
			proto_mail_provider.RegisterMailProviderServiceServer(server, mailProviderService)
			proto_mail_history.RegisterMailHistoryServiceServer(server, mailHistoryService)
		},
		middleware.AuthorizationInterceptor(
			env.SecretService,
			func(action string, resource string) bool {
				hasPermission, err := cacher.Get(resource + "." + action)
				if err != nil {
					return false
				}
				return hasPermission != nil && string(hasPermission) == "true"
			},
			func(id string) *user_context.UserContext {
				userData, err := cacher.Get(id)
				if err != nil || userData == nil {
					return nil
				}
				uCtx := user_context.NewUserContext()
				uCtx.FromBytes(userData)
				return uCtx
			},
		),
	)
}
