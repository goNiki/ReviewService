package pullrequests

import (
	"context"
	"fmt"

	"github.com/goNiki/ReviewService/internal/infrastructure/database"
)

func (r *Repository) AddReviewer(ctx context.Context, tx database.Tx, prId, reviewerID string) error {
	const op = "repository.PullRequests.AddReviewer"

	query := `
		INSERT INTO pr_reviewers (pr_id, reviewer_id)
		VALUES ($1, $2)
		ON CONFLICT (pr_id, reviewer_id) DO NOTHING
	`

	var err error

	if tx != nil {
		_, err = tx.Exec(ctx, query, prId, reviewerID)
	} else {
		_, err = r.db.Pool.Exec(ctx, query, prId, reviewerID)
	}

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil

}
