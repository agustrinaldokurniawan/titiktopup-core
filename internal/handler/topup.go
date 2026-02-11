package handler

import (
	"context"
	"titiktopup-core/internal/domain"
	"titiktopup-core/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)

type TopupHandler struct {
	pb.UnimplementedTopupServiceServer
	repo domain.TransactionRepository
}

func NewTopupHandler(r domain.TransactionRepository) *TopupHandler {
	return &TopupHandler{repo: r}
}

func (h *TopupHandler) GetMenu(ctx context.Context, req *emptypb.Empty) (*pb.MenuResponse, error) {
	categories, err := h.repo.GetCategories()
	if err != nil {
		return nil, err
	}

	var pbCategories []*pb.Category
	for _, c := range categories {
		pbCategories = append(pbCategories, domain.ToProtoCategory(c))
	}

	return &pb.MenuResponse{Categories: pbCategories}, nil
}

func (h *TopupHandler) Checkout(ctx context.Context, req *pb.CheckoutRequest) (*pb.TransactionResponse, error) {
	product, err := h.repo.GetProductByID(uint(req.ProductId))
	if err != nil {
		return nil, err
	}

	trx := &domain.Transaction{
		UserIDGame:    req.UserIdGame,
		ZoneIDGame:    req.ZoneIdGame,
		ProductID:     product.ID,
		TotalPrice:    product.PriceSell,
		PaymentMethod: req.PaymentMethod,
		Status:        domain.StatusPending,
	}

	if err := h.repo.Create(trx); err != nil {
		return nil, err
	}

	return &pb.TransactionResponse{
		Id:         trx.ID,
		Status:     trx.Status,
		TotalPrice: trx.TotalPrice,
	}, nil
}
