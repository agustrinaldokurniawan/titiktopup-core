package handler

import (
	"context"
	"log/slog"
	"titiktopup-core/internal/domain"
	"titiktopup-core/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)

type TopupHandler struct {
	pb.UnimplementedTopupServiceServer
	repo domain.TransactionRepository
	log  *Log
}

type TopupHandlerDeps struct {
	Repo   domain.TransactionRepository
	Logger *slog.Logger
}

func NewTopupHandler(deps TopupHandlerDeps) *TopupHandler {
	return &TopupHandler{
		repo: deps.Repo,
		log:  NewLog(deps.Logger, context.Background()),
	}
}

func (h *TopupHandler) GetMenu(ctx context.Context, req *emptypb.Empty) (*pb.MenuResponse, error) {
	log := h.log.WithContext(ctx)
	logKey := "GetMenu"
	log.Info(logKey, "GetMenu called", NewLogTags(nil).
		WithHandler("TopupHandler"))

	categories, err := h.repo.GetCategories()
	if err != nil {
		log.Error(logKey, "GetMenu failed to get categories", NewLogTags(nil).
			WithHandler("TopupHandler").
			WithError(err))
		return nil, err
	}

	var pbCategories []*pb.Category
	for _, c := range categories {
		pbCategories = append(pbCategories, domain.ToProtoCategory(c))
	}

	log.Info(logKey, "GetMenu success", NewLogTags(nil).
		WithHandler("TopupHandler").
		WithCategoriesCount(len(pbCategories)))
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
