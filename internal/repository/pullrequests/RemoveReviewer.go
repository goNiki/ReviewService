package pullrequests

import (
	"context"
	"fmt"

	"github.com/goNiki/ReviewService/internal/infrastructure/database"
)

func (r *Repository) RemoveReviewer(ctx context.Context, tx database.Tx, prId, reviewerId string) error {
	const op = "repository.pullRequests.RemoveReviewer"

	query := `DELETE FROM pr_reviewers WHERE pr_id = $1 AND reviewer_id = $2`

	var err error

	if tx != nil {
		_, err = tx.Exec(ctx, query, prId, reviewerId)
	} else {
		_, err = r.db.Pool.Exec(ctx, query, prId, reviewerId)
	}

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
