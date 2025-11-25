package pullrequests

import (
	"context"
	"fmt"

	"github.com/goNiki/ReviewService/internal/domain"
	"github.com/goNiki/ReviewService/internal/infrastructure/database"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetPRsByReviewer(ctx context.Context, tx database.Tx, reviewerID string) ([]*domain.PullRequest, error) {
	const op = "repository.pullRequests.GetPRsByReviewer"

	query := `
		SELECT pr.pr_id, pr.pr_name, pr.author_id, pr.team_id, pr.status, pr.created_at, pr.merged_at
		FROM pull_requests pr
		JOIN pr_reviewers prr ON pr.pr_id = prr.pr_id 
		WHERE prr.reviewer_id = $1
		ORDER BY pr.created_at DESC
	`

	var rows pgx.Rows
	var err error

	if tx != nil {
		rows, err = tx.Query(ctx, query, reviewerID)
	} else {
		rows, err = r.db.Pool.Query(ctx, query, reviewerID)
	}

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	pullrequests := make([]*domain.PullRequest, 0, 20)

	for rows.Next() {
		var pr domain.PullRequest
		err := rows.Scan(&pr.Id, &pr.Name, &pr.AuthorID, &pr.TeamId, &pr.Status, &pr.CreatedAt, &pr.MergedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		pullrequests = append(pullrequests, &pr)
	}

	return pullrequests, nil

}
