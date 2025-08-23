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
)

func main() {
	StartGRPCServer()
}

func StartGRPCServer() {
	app := bootstrap.App()
	env := app.Env
	log := app.Log
	db := app.DB
	mailHistoryService := grpcmailhistory.NewMailHistoryService(db, env)
	mailProviderService := grpcmailprovider.NewMailProviderService(db, env)
	mailTmplService := grpcmailtmpl.NewMailTmplService(db, env)
	mailStatusService := grpcmailstatus.NewMailStatusService(db, env)
	typeMailService := grpctypemail.NewTypeMailService(db, env)
	statusHistoryService := grpcstatushistory.NewStatusHistoryService(db, env)
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
