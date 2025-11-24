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

func (r *Repository) SetIsActive(ctx context.Context, tx database.Tx, userId string, isActive bool) (*domain.User, error) {
	const op = "repository.user.SetIsActive"

	query := `
		UPDATE users
		SET is_active = $2, updated_at = NOW()
		WHERE user_id = $1
		RETURNING
			user_id,
			username,
			is_active,
			(SELECT team_name FROM teams WHERE id = users.team_id) as team_name
		`
	var row pgx.Row
	var user repomodels.UserWithTeamName

	if tx != nil {
		row = tx.QueryRow(ctx, query, userId, isActive)
	} else {
		row = r.db.Pool.QueryRow(ctx, query, userId, isActive)
	}

	if err := row.Scan(&user.UserID, &user.UserName, &user.IsActive, &user.TeamName); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	resp := repoconverter.UserWithTeamNameToModel(user)

	return resp, nil
}
