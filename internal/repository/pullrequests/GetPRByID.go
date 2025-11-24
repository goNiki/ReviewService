package pullrequests

import (
	"context"
	"errors"
	"fmt"

	"github.com/goNiki/ReviewService/internal/domain"
	"github.com/goNiki/ReviewService/internal/infrastructure/database"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetPRByID(ctx context.Context, tx database.Tx, prID string) (*domain.PullRequest, error) {
	const op = "repository.PullRequests.GetPRByID"

	query := `
		SELECT pr_id, pr_name, author_id, team_id, status, created_at, merged_at 
        FROM pull_requests 
        WHERE pr_id = $1
	`

	var pr domain.PullRequest
	var row pgx.Row

	if tx != nil {
		row = tx.QueryRow(ctx, query, prID)
	} else {
		row = r.db.Pool.QueryRow(ctx, query, prID)
	}

	err := row.Scan(&pr.Id, &pr.Name, &pr.AuthorID, &pr.TeamId, &pr.Status, &pr.CreatedAt, &pr.MergedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &pr, nil

}
