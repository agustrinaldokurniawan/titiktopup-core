package server

import (
	"context"
	"net/http"
	"os"
)

func RunServers(grpcRegs []GRPCRegistrar, httpRegs []HTTPRegistrar) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go runGRPCServer(grpcRegs)

	httpMux := buildHTTPMux(ctx, httpRegs)
	if err := http.ListenAndServe(httpAddr, httpMux); err != nil {
		Logger.Error("http server failed", "err", err)
		os.Exit(1)
	}
}
