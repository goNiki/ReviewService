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

func (r *Repository) CreateTeam(ctx context.Context, tx database.Tx, name string) (*domain.Team, error) {
	const op = "repository.teams.createTeam"

	query := `INSERT INTO teams (team_name) VALUES ($1) RETURNING id, team_name, created_at`

	var team repomodels.Team
	var row pgx.Row

	if tx != nil {
		row = tx.QueryRow(ctx, query, name)
	} else {
		row = r.db.Pool.QueryRow(ctx, query, name)
	}

	err := row.Scan(&team.Id, &team.Name, &team.Created_at)
	if err != nil {
		return nil, fmt.Errorf("%s: %W", op, err)
	}

	resp := repoconverter.TeamRepoTomModels(team)

	return resp, nil

}
