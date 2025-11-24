package pullrequests

import (
	"context"
	"fmt"

	"github.com/goNiki/ReviewService/internal/infrastructure/database"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetReviewers(ctx context.Context, tx database.Tx, prId string) ([]string, error) {
	const op = "repository.PullRequests.GetReviewers"

	query := `
		SELECT reviewer_id
		FROM pr_reviwers
		WHERE pr_id = $1
		ORDER BY assigned_at
	`
	var rows pgx.Rows
	var err error

	if tx != nil {
		rows, err = tx.Query(ctx, query, prId)
	} else {
		rows, err = r.db.Pool.Query(ctx, query, prId)
	}

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	reviewer_id := make([]string, 0, 5)

	for rows.Next() {
		var revId string

		err := rows.Scan(&revId)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		reviewer_id = append(reviewer_id, revId)
	}

	return reviewer_id, nil

}
