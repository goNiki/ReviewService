package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/goNiki/ReviewService/internal/domain"
	"github.com/goNiki/ReviewService/internal/infrastructure/database"
	repoconverter "github.com/goNiki/ReviewService/internal/repository/converter"
	repomodels "github.com/goNiki/ReviewService/internal/repository/models"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetUserByID(ctx context.Context, tx database.Tx, userId string) (*domain.User, error) {
	const op = "repository.users.GetUserByID"

	query := `
		SELECT user_id, username, team_id, is_active, created_at, updated_at
		FROM users
		WHERE user_id = $1 
	`

	var row pgx.Row
	var user repomodels.User

	if tx != nil {
		row = tx.QueryRow(ctx, query, userId)
	} else {
		row = r.db.Pool.QueryRow(ctx, query, userId)
	}

	if err := row.Scan(&user.UserID, &user.UserName, &user.TeamID, &user.IsActive, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	resp := repoconverter.UserRepoToModel(user)

	return resp, nil
}
