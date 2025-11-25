package users

import (
	"context"
	"fmt"
	"time"

	"github.com/goNiki/ReviewService/internal/domain"
	"github.com/goNiki/ReviewService/internal/infrastructure/database"
)

func (r *Repository) UpdateUser(ctx context.Context, tx database.Tx, user *domain.User) (*domain.User, error) {
	const op = "repository.users.Updateuser"

	query := `
		UPDATE users
		SET username = $2, team_id = $3, is_active = $4, updated_at = $5
		WHERE user_id = $1`

	var err error

	update_at := time.Now()

	if tx != nil {
		_, err = tx.Exec(ctx, query, user.UserId, user.Username, user.TeamID, user.IsActive, update_at)
	} else {
		_, err = r.db.Pool.Exec(ctx, query, user.UserId, user.Username, user.TeamID, user.IsActive, update_at)
	}

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
