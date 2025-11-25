package pullrequests

import (
	"context"
	"fmt"
	"time"

	"github.com/goNiki/ReviewService/internal/domain"
	"github.com/goNiki/ReviewService/models/errorapp"
)

func (s *Service) SetMergeInPR(ctx context.Context, prId string) (*domain.PullRequestWithReviewers, error) {

	pr, err := s.PullRequestRepo.GetPRByID(ctx, nil, prId)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to get PR: %v", errorapp.ErrDatabaseQuery, err)
	}

	if pr == nil {
		return nil, errorapp.ErrPullRequestNotFound
	}

	if pr.Status == domain.Merged {
		reviewers, err := s.PullRequestRepo.GetReviewers(ctx, nil, prId)
		if err != nil {
			return nil, fmt.Errorf("%w: failed to get reviwers: %v", errorapp.ErrDatabaseQuery, err)
		}
		return &domain.PullRequestWithReviewers{
			PR:        pr,
			Reviewers: reviewers,
		}, nil
	}

	now := time.Now()
	pr.Status = domain.Merged
	pr.MergedAt = &now

	if err := s.PullRequestRepo.UpdatePullRequest(ctx, nil, pr); err != nil {
		return nil, fmt.Errorf("%w: failed to updated Pull Request : %v", errorapp.ErrDatabaseQuery, err)
	}

	reviewers, err := s.PullRequestRepo.GetReviewers(ctx, nil, prId)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to get reviwers: %v", errorapp.ErrDatabaseQuery, err)
	}

	return &domain.PullRequestWithReviewers{
		PR:        pr,
		Reviewers: reviewers,
	}, nil

}
