package pullrequests

import (
	"context"
	"fmt"
	"time"

	"github.com/goNiki/ReviewService/internal/domain"
	"github.com/goNiki/ReviewService/internal/infrastructure/database"
)

func (r *Repository) CreatePullRequest(ctx context.Context, tx database.Tx, pr *domain.PullRequest) error {
	const op = "repository.pullRequests.CreatePullRequest"

	query := `
		INSERT INTO pull_requests (pr_id, pr_name, author_id, team_id, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	pr.CreatedAt = time.Now()

	var err error

	if tx != nil {
		_, err = tx.Exec(ctx, query, pr.Id, pr.Name, pr.AuthorID, pr.TeamId, pr.Status, pr.CreatedAt)
	} else {
		_, err = r.db.Pool.Exec(ctx, query, pr.Id, pr.Name, pr.AuthorID, pr.TeamId, pr.Status, pr.CreatedAt)
	}

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
