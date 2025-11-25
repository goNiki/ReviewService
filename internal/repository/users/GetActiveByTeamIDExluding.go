package users

import (
	"context"
	"fmt"

	"github.com/goNiki/ReviewService/internal/domain"
	"github.com/goNiki/ReviewService/internal/infrastructure/database"
	repoconverter "github.com/goNiki/ReviewService/internal/repository/converter"
	repomodels "github.com/goNiki/ReviewService/internal/repository/models"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetActiveByTeamIDExluding(ctx context.Context, tx database.Tx, teamID int64, excludeUserID string) ([]*domain.User, error) {
	const op = "repository.users.GetActiveByTeamIDExluding"

	query := `
		SELECT user_id, username, team_id, is_active, created_at, updated_at
		FROM users
		WHERE team_id = $1 AND is_active = true AND user_id != $2 
		`
	var rows pgx.Rows
	var err error

	if tx != nil {
		rows, err = tx.Query(ctx, query, teamID, excludeUserID)
	} else {
		rows, err = r.db.Pool.Query(ctx, query, teamID, excludeUserID)
	}

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	users := make([]repomodels.User, 0, 5)
	for rows.Next() {
		var user repomodels.User
		if err := rows.Scan(&user.UserID, &user.UserName, &user.TeamID, &user.IsActive, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		users = append(users, user)
	}

	resp := repoconverter.UsersRepoToModel(users)

	return resp, nil
}
