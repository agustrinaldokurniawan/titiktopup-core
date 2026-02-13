package server

import (
	"context"
	"net/http"
	"os"
	"titiktopup-core/constant"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HTTPRegistrar func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

func buildHTTPMux(ctx context.Context, regs []HTTPRegistrar) *http.ServeMux {
	gwMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	for _, reg := range regs {
		if err := reg(ctx, gwMux, grpcEndpoint, opts); err != nil {
			Logger.Error("gateway registration failed", "err", err)
			os.Exit(1)
		}
	}

	mainMux := http.NewServeMux()
	
	// Add Prometheus metrics endpoint
	mainMux.Handle("/metrics", PrometheusMiddleware(MetricsHandler()))
	
	// Add API routes with Prometheus middleware
	apiHandler := PrometheusMiddleware(http.StripPrefix(constant.DefaultApiPrefixPath, gwMux))
	mainMux.Handle(constant.DefaultApiPrefixPath+"/", apiHandler)
	
	// Swagger endpoints (no metrics tracking)
	mainMux.HandleFunc("/swagger.json", serveSwaggerJSON)
	mainMux.HandleFunc("/docs", serveSwaggerUI)
	
	return mainMux
}
