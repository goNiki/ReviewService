package pullrequests

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/goNiki/ReviewService/internal/domain"
	"github.com/goNiki/ReviewService/internal/infrastructure/database"
	"github.com/goNiki/ReviewService/models/errorapp"
)

func (s *Service) CreatePullRequest(ctx context.Context, prID, prName, authorID string) (*domain.PullRequestWithReviewers, error) {

	var result *domain.PullRequestWithReviewers

	err := s.txManager.WithTransaction(ctx, func(ctx context.Context, tx database.Tx) error {
		//проверка, есть ли у пулреквест с таким же ID
		existing, err := s.PullRequestRepo.GetPRByID(ctx, tx, prID)
		if err != nil {
			return fmt.Errorf("%w: failed to check PR existence: %v", errorapp.ErrDatabaseQuery, err)
		}
		//если есть, то ошибка
		if existing != nil {
			return errorapp.ErrPRExists
		}
		//проверяем и получаем информацию по автору
		author, err := s.userRepo.GetUserByID(ctx, tx, authorID)
		if err != nil {
			return fmt.Errorf("%w: failed to get author %v", errorapp.ErrDatabaseQuery, err)
		}

		if author == nil {
			return errorapp.ErrUserNotFound
		}

		pr := &domain.PullRequest{
			Id:       prID,
			Name:     prName,
			AuthorID: author.UserId,
			TeamId:   author.TeamID,
			Status:   domain.Open,
		}
		//создаем сам пулреквест
		if err := s.PullRequestRepo.CreatePullRequest(ctx, tx, pr); err != nil {
			return fmt.Errorf("%w: failed to create pull request %v", errorapp.ErrDatabaseQuery, err)
		}
		//получаем список команды для дальнейшего отбора в роли ревьюера (все кандидаты is_active true кроме автора)
		candidates, err := s.userRepo.GetActiveByTeamIDExluding(ctx, tx, pr.TeamId, authorID)
		if err != nil {
			return fmt.Errorf("%w: failed to get candidates : %v", errorapp.ErrDatabaseQuery, err)
		}
		//рандомно назначаем ревьюверов
		reviewers := selectRandomReviewers(candidates, domain.MaxReviewers)

		for _, reviewerId := range reviewers {
			if err := s.PullRequestRepo.AddReviewer(ctx, tx, prID, reviewerId); err != nil {
				return fmt.Errorf("%w: failed to add reviewerID: %s :%v", errorapp.ErrDatabaseQuery, reviewerId, err)
			}
		}

		result = &domain.PullRequestWithReviewers{
			PR:        pr,
			Reviewers: reviewers,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil

}

func selectRandomReviewers(candidates []*domain.User, max int) []string {
	if len(candidates) == 0 {
		return []string{}
	}

	if len(candidates) <= max {
		result := make([]string, 0, len(candidates))
		for _, u := range candidates {
			result = append(result, u.UserId)
		}
		return result
	}

	userMap := make(map[int]*domain.User)
	for i, u := range candidates {
		userMap[i] = u
	}

	result := make([]string, 0, max)
	usedKeys := make(map[int]bool)

	for len(result) < max {
		randomKey := rand.Intn(len(candidates))

		if !usedKeys[randomKey] {
			usedKeys[randomKey] = true
			result = append(result, userMap[randomKey].UserId)
		}
	}

	return result
}
