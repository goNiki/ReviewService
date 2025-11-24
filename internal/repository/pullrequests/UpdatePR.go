package pullrequests

import (
	"context"
	"fmt"

	"github.com/goNiki/ReviewService/internal/domain"
	"github.com/goNiki/ReviewService/internal/infrastructure/database"
)

func (r *Repository) UpdatePullRequest(ctx context.Context, tx database.Tx, pr *domain.PullRequest) error {
	const op = "repository.PullRequests.UpdatePullRequest"

	query := `
		UPDATE pull_requests
		SET status = $2, merged_at = $3
		WHERE pr_id = $1
	`
	var err error

	if tx != nil {
		_, err = tx.Exec(ctx, query, pr.Id, pr.Status, pr.MergedAt)
	} else {
		_, err = r.db.Pool.Exec(ctx, query, pr.Id, pr.Status, pr.MergedAt)
	}

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil

}
