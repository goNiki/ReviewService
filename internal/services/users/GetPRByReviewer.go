package users

import (
	"context"
	"fmt"

	"github.com/goNiki/ReviewService/internal/domain"
	"github.com/goNiki/ReviewService/models/errorapp"
)

func (s *Service) GetPRByReviewer(ctx context.Context, userID string) (*domain.ReviewerWithPR, error) {

	user, err := s.userRepo.GetUserByID(ctx, nil, userID)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to get user: %v", errorapp.ErrDatabaseQuery)
	}

	if user == nil {
		return nil, errorapp.ErrUserNotFound
	}

	pr, err := s.PullRequestRepo.GetPRsByReviewer(ctx, nil, userID)
	if err != nil {
		return nil, fmt.Errorf("%w, failed to get PR %v", errorapp.ErrDatabaseQuery, err)
	}

	return &domain.ReviewerWithPR{
		UserID:       userID,
		PullRequests: pr,
	}, nil

}
