package services

import (
	"context"

	"github.com/goNiki/ReviewService/internal/domain"
)

type UsersService interface {
	GetPRByReviewer(ctx context.Context, userID string) (*domain.ReviewerWithPR, error)
	SetIsActive(ctx context.Context, userID string, isActive bool) (*domain.User, error)
}

type TeamsService interface {
	CreateTeamWithMembers(ctx context.Context, teamWithMembers *domain.TeamWithMembers) (*domain.TeamWithMembers, error)
	GetTeamWithMembers(ctx context.Context, teamName string) (*domain.TeamWithMembers, error)
}

type PullRequestsService interface {
	CreatePullRequest(ctx context.Context, prID, prName, authorID string) (*domain.PullRequestWithReviewers, error)
	ReassignReviewer(ctx context.Context, prID, oldReviewerID string) (*domain.PullRequestWithReviewers, string, error)
	SetMergeInPR(ctx context.Context, prId string) (*domain.PullRequestWithReviewers, error)
}
