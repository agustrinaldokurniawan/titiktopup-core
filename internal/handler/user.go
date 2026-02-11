package handler

import (
	"context"

	"titiktopup-core/pb"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.UserResponse, error) {
	// For now, return a simple static/dummy profile based on the requested user ID.
	// This can later be wired to a real user repository.
	return &pb.UserResponse{
		Id:    req.GetUserId(),
		Name:  "Demo User",
		Email: "demo@example.com",
		Phone: "0000000000",
	}, nil
}
