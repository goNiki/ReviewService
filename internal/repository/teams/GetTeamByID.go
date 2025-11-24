package teams

import (
	"context"
	"fmt"

	"github.com/goNiki/ReviewService/internal/domain"
	"github.com/goNiki/ReviewService/internal/infrastructure/database"
	repoconverter "github.com/goNiki/ReviewService/internal/repository/converter"
	repomodels "github.com/goNiki/ReviewService/internal/repository/models"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetTeamByID(ctx context.Context, tx database.Tx, id int64) (*domain.Team, error) {
	const op = "repository.teams.GetTeamByID"

	query := `SELECT id, team_name, created_at FROM teams WHERE id = $1`

	var team repomodels.Team
	var row pgx.Row

	if tx != nil {
		row = tx.QueryRow(ctx, query, id)
	} else {
		row = r.db.Pool.QueryRow(ctx, query, id)
	}

	if err := row.Scan(&team.Id, &team.Name, &team.Created_at); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	resp := repoconverter.TeamRepoTomModels(team)

	return resp, nil
}
