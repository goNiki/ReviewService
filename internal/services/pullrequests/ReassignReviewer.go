package pullrequests

import (
	"context"
	"fmt"

	"github.com/goNiki/ReviewService/internal/domain"
	"github.com/goNiki/ReviewService/internal/infrastructure/database"
	"github.com/goNiki/ReviewService/models/errorapp"
)

func (s *Service) ReassignReviewer(ctx context.Context, prID, oldReviewerID string) (*domain.PullRequestWithReviewers, string, error) {

	var result *domain.PullRequestWithReviewers
	var newReviewerID string

	err := s.txManager.WithTransaction(ctx, func(ctx context.Context, tx database.Tx) error {

		pr, err := s.PullRequestRepo.GetPRByID(ctx, tx, prID)
		if err != nil {
			return fmt.Errorf("%w: failed to get PR: %v", errorapp.ErrDatabaseQuery, err)
		}

		if pr == nil {
			return errorapp.ErrPullRequestNotFound
		}

		if pr.Status == domain.Merged {
			return errorapp.ErrPRAlreadyMerged
		}

		currentReviewers, err := s.PullRequestRepo.GetReviewers(ctx, tx, prID)
		if err != nil {
			return fmt.Errorf("%w: failed to get reviewers: %v", errorapp.ErrDatabaseQuery, err)
		}

		isAssigned := false
		for _, reviewerID := range currentReviewers {
			if reviewerID == oldReviewerID {
				isAssigned = true
				break
			}
		}

		if !isAssigned {
			return errorapp.ErrInvalidReviewer
		}

		oldReviewer, err := s.userRepo.GetUserByID(ctx, tx, oldReviewerID)
		if err != nil {
			return fmt.Errorf("%w: failed to get old reviewer: %v", errorapp.ErrDatabaseQuery, err)
		}
		if oldReviewer == nil {
			return errorapp.ErrUserNotFound
		}

		candidates, err := s.userRepo.GetActiveByTeamIDExluding(ctx, tx, oldReviewer.TeamID, oldReviewerID)
		if err != nil {
			return fmt.Errorf("%w: failed to get candidates: %v", errorapp.ErrDatabaseQuery, err)
		}

		filteredCandidates := make([]*domain.User, 0)

		for _, candidate := range candidates {
			if candidate.UserId == pr.AuthorID {
				continue
			}

			isCurrentReviewer := false
			for _, reviewerID := range currentReviewers {
				if candidate.UserId == reviewerID {
					isCurrentReviewer = true
					break
				}
			}

			if !isCurrentReviewer {
				filteredCandidates = append(filteredCandidates, candidate)
			}
		}

		if len(filteredCandidates) == 0 {
			return errorapp.ErrNoCandidate
		}

		selectedReviewers := selectRandomReviewers(filteredCandidates, 1)
		newReviewerID = selectedReviewers[0]

		if err := s.PullRequestRepo.RemoveReviewer(ctx, tx, prID, oldReviewerID); err != nil {
			return fmt.Errorf("%w: failed to remove old reviewer: %v", errorapp.ErrDatabaseQuery, err)
		}

		if err := s.PullRequestRepo.AddReviewer(ctx, tx, prID, newReviewerID); err != nil {
			return fmt.Errorf("%w: failed to add new reviewer: %v", errorapp.ErrDatabaseQuery, err)
		}

		updatedReviewers, err := s.PullRequestRepo.GetReviewers(ctx, tx, prID)
		if err != nil {
			return fmt.Errorf("%w: failed to get updated reviewers: %v", errorapp.ErrDatabaseQuery, err)
		}

		result = &domain.PullRequestWithReviewers{
			PR:        pr,
			Reviewers: updatedReviewers,
		}

		return nil
	})

	if err != nil {
		return nil, "", err
	}

	return result, newReviewerID, nil
}
