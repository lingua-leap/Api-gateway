package handler

import (
	"api-gateway/api/handler/user"
	pb "api-gateway/generated/user"
	"log/slog"
)

type Handler interface {
	NewUserHandler() UserHandler
}

func NewHandler(user pb.UserServiceClient, log *slog.Logger) Handler {
	return &HandlerImpl{user: user, log: log}
}

type HandlerImpl struct {
	user pb.UserServiceClient
	log  *slog.Logger
}

func (h *HandlerImpl) NewUserHandler() UserHandler {
	return user.NewUser(h.user, h.log)
}
