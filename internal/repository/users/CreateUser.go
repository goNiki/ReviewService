package users

import (
	"context"
	"fmt"
	"time"

	"github.com/goNiki/ReviewService/internal/domain"
	"github.com/goNiki/ReviewService/internal/infrastructure/database"
)

func (r *Repository) CreateUser(ctx context.Context, tx database.Tx, user *domain.User) (*domain.User, error) {
	const op = "repository.users.CreateUser"

	query := `INSERT INTO users (user_id, username, team_id, is_active, created_at) VALUES ($1, $2, $3, $4, $5)`

	var err error
	created_at := time.Now()

	if tx != nil {
		_, err = tx.Exec(ctx, query, user.UserId, user.Username, user.TeamID, user.IsActive, created_at)
	} else {
		_, err = r.db.Pool.Exec(ctx, query, user.UserId, user.Username, user.TeamID, user.IsActive, created_at)
	}

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
