package main

import (
	"titiktopup-core/internal/config"
	"titiktopup-core/internal/handler"
	"titiktopup-core/internal/repository"
	"titiktopup-core/internal/server"
	"titiktopup-core/pb"

	"google.golang.org/grpc"
)

func main() {
	db := config.InitDB()

	repo := repository.NewTransactionRepository(db)
	topupHandler := handler.NewTopupHandler(handler.TopupHandlerDeps{
		Repo:   repo,
		Logger: server.Logger,
	})
	userHandler := handler.NewUserHandler(handler.UserHandlerDeps{
		Logger: server.Logger,
	})

	grpcRegs := []server.GRPCRegistrar{
		func(s grpc.ServiceRegistrar) { pb.RegisterTopupServiceServer(s, topupHandler) },
		func(s grpc.ServiceRegistrar) { pb.RegisterUserServiceServer(s, userHandler) },
	}

	httpRegs := []server.HTTPRegistrar{
		pb.RegisterTopupServiceHandlerFromEndpoint,
		pb.RegisterUserServiceHandlerFromEndpoint,
	}

	server.RunServers(grpcRegs, httpRegs)
}
