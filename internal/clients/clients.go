package clients

import (
	"log"
	"titiktopup-core/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	TopupGRPC pb.TopupServiceClient
	Conn      *grpc.ClientConn
}

func InitClients() *Clients {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("❌ Gagal connect ke gRPC client: %v", err)
	}

	topupClient := pb.NewTopupServiceClient(conn)

	log.Println("🔌 Internal Clients Initialized")

	return &Clients{
		TopupGRPC: topupClient,
		Conn:      conn,
	}
}

func (c *Clients) CloseAll() {
	if c.Conn != nil {
		c.Conn.Close()
	}
}
