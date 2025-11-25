package pullrequests

import (
	txmanager "github.com/goNiki/ReviewService/internal/infrastructure/database/transactionManager"
	"github.com/goNiki/ReviewService/internal/repository"
)

type Service struct {
	txManager       txmanager.TransactionManager
	userRepo        repository.UsersRepository
	PullRequestRepo repository.PullRequestsRepository
}

func NewPullRequestService(txmanager txmanager.TransactionManager, userRepo repository.UsersRepository, PRRepo repository.PullRequestsRepository) *Service {
	return &Service{
		txManager:       txmanager,
		userRepo:        userRepo,
		PullRequestRepo: PRRepo,
	}
}
