package server

import "os"

const (
	defaultGrpcAddr      = ":50051"
	defaultHttpAddr      = ":8080"
	defaultGrpcEndpoint  = "localhost:50051"
	defaultSwaggerJSON   = "gen/openapiv2/api_v1.swagger.json"
	defaultApiPrefixPath = "/api/v1"
)

var (
	grpcAddr      = getEnv("GRPC_ADDR", defaultGrpcAddr)
	httpAddr      = getEnv("HTTP_ADDR", defaultHttpAddr)
	grpcEndpoint  = getEnv("GRPC_ENDPOINT", defaultGrpcEndpoint)
	swaggerJSON   = getEnv("SWAGGER_JSON_PATH", defaultSwaggerJSON)
	apiPrefixPath = getEnv("API_PREFIX_PATH", defaultApiPrefixPath)
)

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
