package teams

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

func (r *Repository) GetTeamByName(ctx context.Context, tx database.Tx, name string) (*domain.Team, error) {
	const op = "repository.teams.GetTeamByName"

	query := `SELECT id, team_name, created_at FROM teams WHERE team_name = $1`

	var team repomodels.Team
	var row pgx.Row

	if tx != nil {
		row = tx.QueryRow(ctx, query, name)
	} else {
		row = r.db.Pool.QueryRow(ctx, query, name)
	}

	if err := row.Scan(&team.Id, &team.Name, &team.Created_at); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	resp := repoconverter.TeamRepoTomModels(team)

	return resp, nil

}
