package server

import (
	"net"
	"os"

	"google.golang.org/grpc"
)

type GRPCRegistrar func(s grpc.ServiceRegistrar)

func runGRPCServer(regs []GRPCRegistrar) {
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		Logger.Error("grpc listen failed", "addr", grpcAddr, "err", err)
		os.Exit(1)
	}

	s := grpc.NewServer()
	for _, reg := range regs {
		reg(s)
	}

	if err := s.Serve(lis); err != nil {
		Logger.Error("grpc server failed", "err", err)
		os.Exit(1)
	}
}
