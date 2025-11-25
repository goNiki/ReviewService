package teams

import (
	"context"
	"errors"

	"github.com/goNiki/ReviewService/internal/converter"
	"github.com/goNiki/ReviewService/models/errorapp"
	revV1 "github.com/goNiki/ReviewService/shared/pkg/openapi/reviewerservice/v1"
)

func (a *Api) TeamAddPost(ctx context.Context, req *revV1.Team) (r revV1.TeamAddPostRes, _ error) {
	const op = "TeamAddPost"
	teamData := converter.TeamDtoToModel(*req)

	result, err := a.teamService.CreateTeamWithMembers(ctx, teamData)
	if err != nil {
		a.log.Error(ctx, op, err)

		switch {
		case errors.Is(err, errorapp.ErrTeamAlreadyExists):
			return &revV1.BadRequestError{
				Error: revV1.BadRequestErrorError{
					Code:    revV1.BadRequestErrorErrorCodeTEAMEXISTS,
					Message: "team already exists",
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

	responseTeam := converter.TeamModelToDto(result)

	return &revV1.TeamAddPostCreated{
		Team: revV1.NewOptTeam(responseTeam),
	}, nil
}
