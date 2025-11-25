package users

import "github.com/goNiki/ReviewService/internal/repository"

type Service struct {
	userRepo        repository.UsersRepository
	PullRequestRepo repository.PullRequestsRepository
}

func NewUserService(userRepo repository.UsersRepository, PR repository.PullRequestsRepository) *Service {
	return &Service{
		userRepo:        userRepo,
		PullRequestRepo: PR,
	}
}
