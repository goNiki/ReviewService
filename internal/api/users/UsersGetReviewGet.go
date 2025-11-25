package users

import (
	"context"
	"errors"

	"github.com/goNiki/ReviewService/internal/converter"
	"github.com/goNiki/ReviewService/models/errorapp"
	revV1 "github.com/goNiki/ReviewService/shared/pkg/openapi/reviewerservice/v1"
)

func (a *Api) UsersGetReviewGet(ctx context.Context, params revV1.UsersGetReviewGetParams) (r revV1.UsersGetReviewGetRes, _ error) {
	const op = "UsersGetReviewGet"
	UserWithPR, err := a.userService.GetPRByReviewer(ctx, params.UserID)
	if err != nil {
		a.log.Error(ctx, op, err)
		switch {
		case errors.Is(err, errorapp.ErrUserNotFound):
			return &revV1.NotFoundError{
				Error: revV1.NotFoundErrorError{
					Code:    revV1.NotFoundErrorErrorCodeNOTFOUND,
					Message: "user not found",
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

	resp := converter.ReviewerWithPRToDTO(UserWithPR)

	return resp, nil
}
