package repository

import (
	"context"

	"github.com/goNiki/ReviewService/internal/domain"
	"github.com/goNiki/ReviewService/internal/infrastructure/database"
)

type UsersRepository interface {
	CreateUser(ctx context.Context, tx database.Tx, user *domain.User) (*domain.User, error)
	GetUserByID(ctx context.Context, tx database.Tx, userId string) (*domain.User, error)
	GetUsersByTeamID(ctx context.Context, tx database.Tx, teamID int64) ([]*domain.User, error)
	UpdateUser(ctx context.Context, tx database.Tx, user *domain.User) (*domain.User, error)
	SetIsActive(ctx context.Context, tx database.Tx, userId string, isActive bool) (*domain.User, error)
	GetActiveByTeamIDExluding(ctx context.Context, tx database.Tx, teamID int64, excludeUserID string) ([]*domain.User, error)
}

type TeamsRepository interface {
	CreateTeam(ctx context.Context, tx database.Tx, name string) (*domain.Team, error)
	GetTeamByID(ctx context.Context, tx database.Tx, id int64) (*domain.Team, error)
	GetTeamByName(ctx context.Context, tx database.Tx, name string) (*domain.Team, error)
}

type PullRequestsRepository interface {
	CreatePullRequest(ctx context.Context, tx database.Tx, pr *domain.PullRequest) error
	GetPRByID(ctx context.Context, tx database.Tx, prID string) (*domain.PullRequest, error)
	AddReviewer(ctx context.Context, tx database.Tx, prId, reviewerID string) error
	GetPRsByReviewer(ctx context.Context, tx database.Tx, reviewerID string) ([]*domain.PullRequest, error)
	GetReviewers(ctx context.Context, tx database.Tx, prId string) ([]string, error)
	RemoveReviewer(ctx context.Context, tx database.Tx, prId, reviewerId string) error
	UpdatePullRequest(ctx context.Context, tx database.Tx, pr *domain.PullRequest) error
}
