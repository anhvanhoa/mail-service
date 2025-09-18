package main

import (
	"context"

	"mail-service/bootstrap"
	grpcservice "mail-service/infrastructure/grpc_service"
	grpcmailhistory "mail-service/infrastructure/grpc_service/mail_history"
	grpcmailprovider "mail-service/infrastructure/grpc_service/mail_provider"
	grpcmailstatus "mail-service/infrastructure/grpc_service/mail_status"
	grpcmailtmpl "mail-service/infrastructure/grpc_service/mail_tmpl"
	grpcstatushistory "mail-service/infrastructure/grpc_service/status_history"
	grpctypemail "mail-service/infrastructure/grpc_service/type_mail"

	"github.com/anhvanhoa/service-core/domain/discovery"
)

func main() {
	StartGRPCServer()
}

func StartGRPCServer() {
	app := bootstrap.App()
	env := app.Env
	log := app.Log
	db := app.DB

	discoveryConfig := &discovery.DiscoveryConfig{
		ServiceName:   env.NameService,
		ServicePort:   env.PortGrpc,
		ServiceHost:   env.HostGrpc,
		IntervalCheck: env.IntervalCheck,
		TimeoutCheck:  env.TimeoutCheck,
	}

	discovery, err := discovery.NewDiscovery(discoveryConfig)
	if err != nil {
		log.Fatal("Failed to create discovery: " + err.Error())
	}
	discovery.Register()

	mailHistoryService := grpcmailhistory.NewMailHistoryService(db)
	mailProviderService := grpcmailprovider.NewMailProviderService(db)
	mailTmplService := grpcmailtmpl.NewMailTmplService(db)
	mailStatusService := grpcmailstatus.NewMailStatusService(db)
	typeMailService := grpctypemail.NewTypeMailService(db)
	statusHistoryService := grpcstatushistory.NewStatusHistoryService(db)
	grpcSrv := grpcservice.NewGRPCServer(
		env,
		log,
		mailHistoryService,
		mailProviderService,
		mailTmplService,
		mailStatusService,
		typeMailService,
		statusHistoryService,
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}
