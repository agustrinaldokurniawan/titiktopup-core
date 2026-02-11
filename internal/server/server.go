package server

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCRegistrar func(s grpc.ServiceRegistrar)
type HTTPRegistrar func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

func RunServers(grpcRegs []GRPCRegistrar, httpRegs []HTTPRegistrar) {
	// gRPC server
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		for _, reg := range grpcRegs {
			reg(s)
		}

		log.Println("🚀 gRPC Server: :50051")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// gRPC-Gateway setup
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	for _, reg := range httpRegs {
		if err := reg(ctx, mux, "localhost:50051", opts); err != nil {
			log.Fatalf("failed to register gateway: %v", err)
		}
	}

	// Standard Go Mux
	mainMux := http.NewServeMux()

	// 1. API Prefix
	mainMux.Handle("/api/v1/", http.StripPrefix("/api/v1", mux))

	// 2. Swagger JSON
	mainMux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		http.ServeFile(w, r, "gen/openapiv2/api_v1.swagger.json")
	})

	// 3. Swagger UI
	mainMux.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		html := `
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="utf-8" />
            <title>TitikTopup API Docs</title>
            <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" />
        </head>
        <body>
            <div id="swagger-ui"></div>
            <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
            <script>
                window.onload = () => {
                    window.ui = SwaggerUIBundle({
                        url: '/swagger.json',
                        dom_id: '#swagger-ui',
                    });
                };
            </script>
        </body>
        </html>`
		w.Write([]byte(html))
	})

	log.Println("🌐 REST API Server: :8080")
	log.Fatal(http.ListenAndServe(":8080", mainMux))
}
