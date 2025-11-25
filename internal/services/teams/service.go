package teams

import (
	txmanager "github.com/goNiki/ReviewService/internal/infrastructure/database/transactionManager"
	"github.com/goNiki/ReviewService/internal/repository"
)

type Service struct {
	txManager txmanager.TransactionManager
	teamRepo  repository.TeamsRepository
	userRepo  repository.UsersRepository
}

func NewTeamsService(txManager txmanager.TransactionManager, teamRepo repository.TeamsRepository, userRepo repository.UsersRepository) *Service {
	return &Service{
		txManager: txManager,
		teamRepo:  teamRepo,
		userRepo:  userRepo,
	}
}
