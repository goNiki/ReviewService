package teams

import (
	"context"
	"errors"

	"github.com/goNiki/ReviewService/internal/converter"
	"github.com/goNiki/ReviewService/models/errorapp"
	revV1 "github.com/goNiki/ReviewService/shared/pkg/openapi/reviewerservice/v1"
)

func (a *Api) TeamGetGet(ctx context.Context, params revV1.TeamGetGetParams) (r revV1.TeamGetGetRes, _ error) {
	const op = "TeamGetGet"

	team, err := a.teamService.GetTeamWithMembers(ctx, params.TeamName)
	if err != nil {
		a.log.Error(ctx, op, err)
		switch {
		case errors.Is(err, errorapp.ErrTeamNotFound):
			return &revV1.NotFoundError{
				Error: revV1.NotFoundErrorError{
					Code:    revV1.NotFoundErrorErrorCodeNOTFOUND,
					Message: "team not found",
				},
			}, nil
		default:
			return &revV1.InternalServerError{
				Error: revV1.InternalServerErrorError{
					Code:    revV1.InternalServerErrorErrorCodeINTERNALERROR,
					Message: "internal service error",
				},
			}, nil

		}
	}

	resp := converter.TeamModelToDto(team)

	return &resp, nil
}
