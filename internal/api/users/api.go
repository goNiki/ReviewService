package users

import (
	"github.com/goNiki/ReviewService/internal/infrastructure/logger"
	"github.com/goNiki/ReviewService/internal/services/users"
)

type Api struct {
	log         *logger.Log
	userService *users.Service
}

func NewUsersApi(log *logger.Log, userService *users.Service) *Api {
	return &Api{
		log:         log,
		userService: userService,
	}
}
