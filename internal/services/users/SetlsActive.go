package users

import (
	"context"
	"fmt"

	"github.com/goNiki/ReviewService/internal/domain"
	"github.com/goNiki/ReviewService/models/errorapp"
)

func (s *Service) SetIsActive(ctx context.Context, userID string, isActive bool) (*domain.User, error) {

	user, err := s.userRepo.SetIsActive(ctx, nil, userID, isActive)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to set isActive: %v", errorapp.ErrDatabaseQuery, err)
	}

	if user == nil {
		return nil, errorapp.ErrUserNotFound
	}

	return user, nil
}
